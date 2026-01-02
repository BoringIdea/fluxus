// @ts-ignore
import { ZetaChainClient } from '@zetachain/toolkit/client';

export interface CCTXRecord {
  transactionHash: string;
  status: string;
  status_message?: string;
  sender_chain_id: string;
  receiver_chainId: string;
  outbound_tx_hash?: string;
  outbound_tx_tss_nonce?: number;
  confirmed_on_destination: boolean;
  blockTimestamp?: Date; // Optional since ZetaChain API doesn't provide this
  event?: string;
}

export interface CCTXTrackingCallbacks {
  onCCTXUpdate: (record: CCTXRecord) => void;
  onTrackingStart: () => void;
  onTrackingComplete: () => void;
  onError: (error: any) => void;
}

export class CCTXTracker {
  private client: ZetaChainClient;
  private network: string;
  private recordsCache: Map<string, CCTXRecord[]> = new Map();
  private inflightRequests: Map<string, Promise<CCTXRecord[]>> = new Map();

  constructor(network: 'testnet' | 'mainnet' = 'testnet') {
    this.network = network;
    this.client = new ZetaChainClient({
      network: this.network
    });
  }

  /**
   * Track a CCTX transaction by hash
   */
  async trackCCTX(
    hash: string,
    callbacks: CCTXTrackingCallbacks
  ): Promise<CCTXRecord[]> {
    // Serve from in-memory cache immediately if available
    const cached = this.recordsCache.get(hash);
    if (cached && cached.length > 0) {
      // Emit cached updates synchronously for UI hydration
      cached.forEach(record => callbacks.onCCTXUpdate(record));
      return cached;
    }

    // De-duplicate concurrent requests for the same hash
    const inflight = this.inflightRequests.get(hash);
    if (inflight) {
      const records = await inflight;
      records.forEach(record => callbacks.onCCTXUpdate(record));
      return records;
    }

    console.log(`%c🔗 Starting CCTX tracking for: ${hash}`, 'color: #4CAF50; font-weight: bold;');
    
    callbacks.onTrackingStart();

    try {
      const promise = this.client.trackCCTX({
        hash,
        json: true,
        emitter: {
          emit: (event: string, data: any) => {
            console.log(`%c🔗 CCTX Event [${event}]`, 'color: #4CAF50; font-weight: bold;');
            
            if (event === 'mined-fail' && data.cctxs) {
              Object.keys(data.cctxs).forEach(transactionHash => {
                const cctx = data.cctxs[transactionHash][0];
                
                const cctxRecord: CCTXRecord = {
                  transactionHash,
                  ...cctx,
                  // blockTimestamp will be set from crossChainStatus API data
                  event
                };
                
                callbacks.onCCTXUpdate(cctxRecord);
                this.logFailedTransaction(transactionHash, cctx);
              });
            } else {
              console.log(data);
            }
          }
        }
      });

      this.inflightRequests.set(hash, promise);
      const result = await promise;

      console.group('%c🎉 CCTX Tracking Complete', 'color: #4CAF50; font-size: 16px; font-weight: bold;');
      
      const records: CCTXRecord[] = [];
      
      console.log('🔍 Raw result from trackCCTX:', result);
      
      if (result && result.cctxs) {
        console.log('📋 Processing CCTX data:', result.cctxs);
        Object.keys(result.cctxs).forEach(transactionHash => {
          const cctxArray = result.cctxs[transactionHash];
          if (cctxArray && cctxArray.length > 0) {
            const cctx = cctxArray[0];
            
            const cctxRecord: CCTXRecord = {
              transactionHash,
              ...cctx
              // blockTimestamp will be set from crossChainStatus API data
            };
            
            console.log(`📄 Created CCTX record for ${transactionHash}:`, cctxRecord);
            records.push(cctxRecord);
            callbacks.onCCTXUpdate(cctxRecord);
            this.logTransaction(transactionHash, cctx);
          } else {
            console.log(`⚠️ Empty CCTX array for ${transactionHash}`);
          }
        });
      } else if (result) {
        // Handle direct result format (if result is the CCTX data itself)
        console.log('🔄 Trying to handle result as direct CCTX data...');
        if (typeof result === 'object') {
          Object.keys(result).forEach(transactionHash => {
            const cctxArray = result[transactionHash];
            if (Array.isArray(cctxArray) && cctxArray.length > 0) {
              const cctx = cctxArray[0];
              
              const cctxRecord: CCTXRecord = {
                transactionHash,
                ...cctx
                // blockTimestamp will be set from crossChainStatus API data
              };
              
              console.log(`📄 Created CCTX record from direct result for ${transactionHash}:`, cctxRecord);
              records.push(cctxRecord);
              callbacks.onCCTXUpdate(cctxRecord);
              this.logTransaction(transactionHash, cctx);
            }
          });
        }
      } else {
        console.log('❌ No CCTX data found in result');
      }
      
      console.groupEnd();
      callbacks.onTrackingComplete();
      // Cache successful result for future navigations
      if (records.length > 0) {
        this.recordsCache.set(hash, records);
      }
      this.inflightRequests.delete(hash);
      
      return records;
    } catch (error: any) {
      console.group('%c❌ CCTX Tracking Error', 'color: #f44336; font-size: 16px; font-weight: bold;');
      console.error('Error details:', error.message || error);
      console.error('Full error:', error);
      console.groupEnd();
      
      callbacks.onError(error);
      callbacks.onTrackingComplete();
      this.inflightRequests.delete(hash);
      
      throw error;
    }
  }


  private logFailedTransaction(txHash: string, cctx: any) {
    console.group(`%c❌ Transaction Failed: ${txHash.substring(0, 10)}...`, 'color: #f44336; font-weight: bold;');
    console.log(`%c📍 Status:`, 'color: #ff9800; font-weight: bold;', cctx.status);
    console.log(`%c💬 Message:`, 'color: #ff9800; font-weight: bold;', cctx.status_message);
    console.log(`%c🔗 Sender Chain:`, 'color: #2196F3;', cctx.sender_chain_id);
    console.log(`%c🎯 Receiver Chain:`, 'color: #2196F3;', cctx.receiver_chainId);
    console.log(`%c📤 Outbound TX:`, 'color: #9C27B0;', cctx.outbound_tx_hash);
    console.log(`%c🔢 TSS Nonce:`, 'color: #607D8B;', cctx.outbound_tx_tss_nonce);
    console.log(`%c✅ Confirmed on Destination:`, 'color: #4CAF50;', cctx.confirmed_on_destination ? 'Yes' : 'No');
    console.groupEnd();
  }

  private logTransaction(txHash: string, cctx: any) {
    console.group(`%c📋 Transaction: ${txHash.substring(0, 10)}...`, 'color: #2196F3; font-weight: bold;');
    
    const statusColor = cctx.status === 'Aborted' ? '#f44336' : 
                       cctx.status === 'OutboundMined' ? '#4CAF50' : '#ff9800';
    
    console.log(`%c📊 Status:`, 'font-weight: bold;', `%c${cctx.status}`, `color: ${statusColor}; font-weight: bold;`);
    
    if (cctx.status_message) {
      console.log(`%c💬 Status Message:`, 'font-weight: bold;', cctx.status_message);
    }
    
    console.log(`%c🔗 Chain Info:`, 'color: #9C27B0; font-weight: bold;');
    console.log(`  • Sender Chain: ${cctx.sender_chain_id}`);
    console.log(`  • Receiver Chain: ${cctx.receiver_chainId}`);
    
    if (cctx.outbound_tx_hash) {
      console.log(`%c📤 Outbound Transaction:`, 'color: #FF5722; font-weight: bold;');
      console.log(`  • Hash: ${cctx.outbound_tx_hash}`);
      console.log(`  • TSS Nonce: ${cctx.outbound_tx_tss_nonce}`);
    }
    
    console.log(`%c✅ Confirmed on Destination:`, 'color: #4CAF50; font-weight: bold;', 
               cctx.confirmed_on_destination ? '✅ Yes' : '❌ No');
    
    console.groupEnd();
  }
}

export const defaultCCTXTracker = new CCTXTracker('testnet');
