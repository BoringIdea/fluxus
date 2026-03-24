"use client";
import { useState, useEffect, useRef } from 'react';
import { useAccount, useSignMessage } from 'wagmi';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { io, Socket } from 'socket.io-client';

interface ChatProps {
  contractAddress: string;
  collection?: any;
  chainId?: number;
}

interface Message {
  id: number;
  userId: number;
  walletAddress: string;
  content: string;
  createdAt: string;
}

interface OnlineUser {
  id: number;
  walletAddress: string;
}

// Format wallet address for display
const formatWalletAddress = (address: string) => {
  return `${address.slice(0, 6)}...${address.slice(-4)}`;
};

// Get default avatar from wallet address
const getDefaultAvatar = (address: string) => {
  return address.charAt(2).toUpperCase();
};

export default function Chat({ contractAddress, collection, chainId }: ChatProps) {
  const { address } = useAccount();
  const { signMessageAsync, isPending, error: signError } = useSignMessage();
  const [socket, setSocket] = useState<Socket | null>(null);
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [messages, setMessages] = useState<Message[]>([]);
  const [onlineUsers, setOnlineUsers] = useState<OnlineUser[]>([]);
  const [newMessage, setNewMessage] = useState('');
  const [error, setError] = useState<string | null>(null);
  const [roomId, setRoomId] = useState<number | null>(null);
  const [isLoadingHistory, setIsLoadingHistory] = useState(false);
  const [currentPage, setCurrentPage] = useState(1);
  const [hasMoreMessages, setHasMoreMessages] = useState(true);
  const messagesEndRef = useRef<HTMLDivElement>(null);
  const backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL || 'http://localhost:8080';

  // Auto scroll to bottom when new messages arrive
  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
  };

  // Get chat room ID from backend
  const getChatRoomId = async () => {
    try {
      const response = await fetch(
        `${backendUrl}/api/v1/chat/rooms/${contractAddress}?chainId=${chainId}`
      );
      const data = await response.json();
      if (data.success && data.room) {
        setRoomId(data.room.id);
        return data.room.id;
      }
      throw new Error('Failed to get chat room ID');
    } catch (error) {
      console.error('Error getting chat room ID:', error);
      setError('Failed to get chat room information');
      return null;
    }
  };

  // Load chat history
  const loadChatHistory = async (page = 1, append = false) => {
    setIsLoadingHistory(true);
    try {
      // Validate parameters
      if (!contractAddress) {
        throw new Error('Contract address is required');
      }

      const url = `${backendUrl}/api/v1/chat/rooms/${encodeURIComponent(contractAddress)}/messages?page=${page}&pageSize=50`;
      
      const response = await fetch(url);
      
      if (!response.ok) {
        const errorText = await response.text();
        console.error('API Error:', response.status, response.statusText, errorText);
        throw new Error(`API Error: ${response.status} ${response.statusText}`);
      }
      
      const data = await response.json();
      
      if (data.success && data.messages) {
        if (append) {
          setMessages(prev => [...data.messages, ...prev]);
        } else {
          setMessages(data.messages);
        }
        setCurrentPage(page);
        setHasMoreMessages(data.messages.length === 50);
      } else {
        console.log('No messages found or API error:', data);
      }
    } catch (error) {
      console.error('Error loading chat history:', error);
      setError(`Failed to load chat history: ${error instanceof Error ? error.message : 'Unknown error'}`);
    } finally {
      setIsLoadingHistory(false);
    }
  };

  // Load more messages
  const loadMoreMessages = async () => {
    if (hasMoreMessages && !isLoadingHistory) {
      await loadChatHistory(currentPage + 1, true);
    }
  };

  useEffect(() => {
    scrollToBottom();
  }, [messages]);

  // Initialize socket connection
  useEffect(() => {
    if (isAuthenticated) {
      const connectToChat = async () => {
        // First get the chat room ID
        const chatRoomId = await getChatRoomId();
        if (!chatRoomId) {
          setError('Failed to get chat room information');
          return;
        }

        const newSocket = io(backendUrl, {
          transports: ['websocket', 'polling'],
          withCredentials: true,
        });
        setSocket(newSocket);

        // Load chat history first
        await loadChatHistory();

        // Join room with collection address (backend will get/create room)
        newSocket.emit('join_room', {
          roomId: contractAddress, // Use collection address for join_room
          token: localStorage.getItem('chat_token'),
          chainId: chainId
        });

      // Listen for messages
      newSocket.on('receive_message', (message: Message) => {
        setMessages(prev => [...prev, message]);
      });

      // Listen for message send errors
      newSocket.on('message_sent', (data: any) => {
        if (!data.success) {
          console.error('Failed to send message:', data.message);
          setError(data.message || 'Failed to send message');
        }
      });

      // Listen for user events
      newSocket.on('user_joined', (user: OnlineUser) => {
        setOnlineUsers(prev => [...prev, user]);
      });

      newSocket.on('user_left', (user: OnlineUser) => {
        setOnlineUsers(prev => prev.filter(u => u.id !== user.id));
      });

      newSocket.on('room_joined', (data: any) => {
        if (data.success) {
          console.log('Joined room successfully');
        } else {
          console.error('Failed to join room:', data.message);
          setError(data.message || 'Failed to join chat room');
        }
      });

        return () => {
          newSocket.disconnect();
        };
      };

      connectToChat();
    }
  }, [isAuthenticated, contractAddress, backendUrl]);

  // Handle authentication
  const handleSignMessage = async () => {
    if (!address) return;

    setIsLoading(true);
    setError(null);
    
    // Reset any previous signature errors
    if (signError) {
      // Clear previous signature error
    }

    try {
      // Get signature message from backend
      const response = await fetch(`${backendUrl}/api/v1/auth/signature-message?walletAddress=${address}`);
      if (!response.ok) {
        throw new Error(`Failed to fetch signature message: ${response.statusText}`);
      }
      const data = await response.json();

      if (!data.success) {
        throw new Error(data.message);
      }

      // Check if there's already a pending signature
      if (isPending) {
        throw new Error('Signature is already pending');
      }

      // Check for previous signature errors
      if (signError) {
        throw new Error(`Previous signature failed: ${signError.message}`);
      }
      
      // Sign message with wallet
      const signature = await signMessageAsync({ 
        message: data.message 
      });

      if (!signature || typeof signature !== 'string') {
        throw new Error('Signature was not generated or is not a string');
      }

      // Verify signature with backend
      const verifyResponse = await fetch(`${backendUrl}/api/v1/auth/verify-signature`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          walletAddress: address,
          signature,
          message: data.message,
          collectionAddress: contractAddress,
          chainId: chainId
        })
      });

      const verifyData = await verifyResponse.json();

      if (verifyData.success && verifyData.token) {
        localStorage.setItem('chat_token', verifyData.token);
        setIsAuthenticated(true);
        // Load chat room data
        await loadChatRoom();
      } else {
        throw new Error(verifyData.message || 'Authentication failed');
      }
    } catch (err: any) {
      setError(err.message || 'Authentication failed');
    } finally {
      setIsLoading(false);
    }
  };

  // Load chat room data
  const loadChatRoom = async () => {
    try {
      const token = localStorage.getItem('chat_token');
      const response = await fetch(`${backendUrl}/api/v1/chat/rooms/${contractAddress}?chainId=${chainId}`, {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      });

      const data = await response.json();
      if (data.success) {
        setOnlineUsers(data.members || []);
        // Load messages
        await loadMessages();
      }
    } catch (err) {
      console.error('Failed to load chat room:', err);
    }
  };

  // Load messages
  const loadMessages = async () => {
    try {
      // Temporarily disable API calls to avoid database issues
      const token = localStorage.getItem('chat_token');
      const response = await fetch(`${backendUrl}/api/v1/chat/rooms/${contractAddress}/messages?chainId=${chainId}`, {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      });

      const data = await response.json();
      if (data.success) {
        setMessages(data.messages || []);
      }
      
      // For now, just set empty messages to avoid backend errors
      // setMessages([]);
    } catch (err) {
      console.error('Failed to load messages:', err);
    }
  };

  // Send message
  const handleSendMessage = async () => {
    if (!newMessage.trim() || !socket) return;

    const message = newMessage.trim();
    setNewMessage('');

    try {
      socket.emit('send_message', {
        roomId: contractAddress, // Use collection address for send_message too
        messageType: 'text',
        content: message,
        chainId: chainId
      });
    } catch (err) {
      console.error('Failed to send message:', err);
    }
  };

  // Handle key press
  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault();
      handleSendMessage();
    }
  };

  // Check if user owns NFT
  const checkNFTOwnership = async () => {
    if (!address) return false;

    try {
      const token = localStorage.getItem('chat_token');
      const response = await fetch(`${backendUrl}/api/v1/chat/rooms/${contractAddress}?chainId=${chainId}`, {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      });

      return response.ok;
    } catch {
      return false;
    }
  };

  // Check authentication on mount
  useEffect(() => {
    const token = localStorage.getItem('chat_token');
    if (token) {
      setIsAuthenticated(true);
      loadChatRoom();
    }
  }, []);

  if (!address) {
    return (
      <div className="flex h-64 flex-col items-center justify-center border border-black/10 bg-[color:var(--bg-surface)] px-4 text-center">
        <p className="mb-2 font-heading text-[28px] leading-none text-[color:var(--text-primary)]">Connect your wallet to join the chat</p>
        <p className="text-sm text-[color:var(--text-secondary)]">You need to connect your wallet to participate in the chat.</p>
      </div>
    );
  }

  if (!isAuthenticated) {
    return (
      <div className="flex h-64 flex-col items-center justify-center gap-4 border border-black/10 bg-[color:var(--bg-surface)] px-4 text-center">
        <div>
          <h3 className="mb-2 font-heading text-[26px] leading-none text-[color:var(--text-primary)]">Join Collection Chat</h3>
          <p className="text-sm text-[color:var(--text-secondary)]">
            Sign a message to verify you own NFTs from this collection.
          </p>
        </div>
        {error && (
          <div className="w-full max-w-md border border-red-200 bg-red-50 p-3 text-sm text-red-600">
            {error}
          </div>
        )}
        <Button onClick={handleSignMessage} disabled={isLoading || isPending} className="px-6 py-2">
          {isLoading || isPending ? 'Signing...' : 'Sign Message'}
        </Button>
      </div>
    );
  }

  return (
    <div className="flex h-[600px] flex-col overflow-hidden border border-black/10 bg-[color:var(--bg-surface)]">
      <div className="flex items-center justify-between border-b border-black/10 bg-[color:var(--bg-muted)] p-4">
        <div>
          <h3 className="font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)]">Collection Chat</h3>
          <p className="text-sm text-[color:var(--text-secondary)]">
            {onlineUsers.length} member{onlineUsers.length !== 1 ? 's' : ''} online
          </p>
        </div>
        <div className="flex items-center gap-2">
          <div className="h-2 w-2 bg-[color:var(--color-primary)]" />
          <span className="font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">Connected</span>
        </div>
      </div>

      {/* Messages */}
      <div className="flex-1 overflow-y-auto p-4 space-y-4">
        {/* Load More Button */}
        {hasMoreMessages && messages.length > 0 && (
          <div className="text-center">
            <Button onClick={loadMoreMessages} disabled={isLoadingHistory} className="px-4 py-2">
              {isLoadingHistory ? 'Loading...' : 'Load More Messages'}
            </Button>
          </div>
        )}

        {isLoadingHistory && messages.length === 0 ? (
          <div className="py-8 text-center font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">
            <p>Loading chat history...</p>
          </div>
        ) : messages.length === 0 ? (
          <div className="py-8 text-center font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">
            <p>No messages yet. Start the conversation!</p>
          </div>
        ) : (
          messages.map((message) => {
            const isOwnMessage = message.walletAddress.toLowerCase() === address?.toLowerCase();
            return (
              <div key={message.id} className={`flex items-start ${isOwnMessage ? 'justify-end' : 'justify-start'}`}>
                <div className={`flex items-start space-x-3 max-w-[80%] ${isOwnMessage ? 'flex-row-reverse space-x-reverse' : ''}`}>
                  <div className="flex h-8 w-8 flex-shrink-0 items-center justify-center bg-[color:var(--bg-muted)] font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--color-primary)]">
                    {getDefaultAvatar(message.walletAddress)}
                  </div>
                  <div className={`flex-1 ${isOwnMessage ? 'text-right' : 'text-left'}`}>
                    <div className={`flex items-center space-x-2 mb-1 ${isOwnMessage ? 'justify-end' : 'justify-start'}`}>
                      <span className="font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-primary)]">
                        {formatWalletAddress(message.walletAddress)}
                      </span>
                      <span className="font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-muted)]">
                        {new Date(message.createdAt).toLocaleTimeString()}
                      </span>
                    </div>
                    <div className={`inline-block border border-black/10 p-3 ${
                      isOwnMessage 
                        ? 'bg-[color:var(--color-primary)] text-white' 
                        : 'bg-[color:var(--bg-muted)] text-[color:var(--text-primary)]'
                    }`}>
                      <p className="text-sm">{message.content}</p>
                    </div>
                  </div>
                </div>
              </div>
            );
          })
        )}
        <div ref={messagesEndRef} />
      </div>

      {/* Message Input */}
      <div className="border-t border-black/10 bg-[color:var(--bg-muted)] p-4 pb-12 pt-6">
        <div className="flex space-x-2">
          <Input
            value={newMessage}
            onChange={(e) => setNewMessage(e.target.value)}
            onKeyPress={handleKeyPress}
            placeholder="Type a message..."
            className="flex-1"
            maxLength={500}
          />
          <Button
            onClick={handleSendMessage}
            disabled={!newMessage.trim()}
            className="px-4"
          >
            Send
          </Button>
        </div>
        <div className="mt-1 font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-muted)]">
          {newMessage.length}/500 characters
        </div>
      </div>
    </div>
  );
}
