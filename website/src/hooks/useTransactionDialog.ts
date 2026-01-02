import { useCallback, useEffect, useMemo, useRef, useState } from 'react';

export type TransactionStatus = 'pending' | 'confirming' | 'confirmed' | 'error';

export interface TransactionDialogState {
  isOpen: boolean;
  status: TransactionStatus;
  hash?: string;
  error?: string;
}

interface UseTransactionDialogParams {
  hash?: `0x${string}` | string | null;
  isPending?: boolean;
  isConfirming?: boolean;
  isConfirmed?: boolean;
  isWriteContractError?: boolean;
  writeContractError?: unknown;
  onConfirmed?: () => void;
}

export function useTransactionDialog(params: UseTransactionDialogParams) {
  const {
    hash,
    isPending,
    isConfirming,
    isConfirmed,
    isWriteContractError,
    writeContractError,
    onConfirmed,
  } = params;

  const [dialogState, setDialogState] = useState<TransactionDialogState>({ isOpen: false, status: 'pending' });
  const viewedRef = useRef<Set<string>>(new Set());

  const errorMessage = useMemo(() => {
    if (!isWriteContractError) return undefined;
    if (!writeContractError) return 'Unknown error';
    if (typeof writeContractError === 'string') return writeContractError;
    if (writeContractError instanceof Error) return writeContractError.message;
    try {
      return JSON.stringify(writeContractError);
    } catch {
      return 'Unknown error';
    }
  }, [isWriteContractError, writeContractError]);

  useEffect(() => {
    const viewed = viewedRef.current;

    if (hash && !viewed.has(hash)) {
      if (isPending) {
        setDialogState({ isOpen: true, status: 'pending', hash: hash as string });
      } else if (isConfirming) {
        setDialogState({ isOpen: true, status: 'confirming', hash: hash as string });
      } else if (isConfirmed) {
        setDialogState({ isOpen: true, status: 'confirmed', hash: hash as string });
        onConfirmed?.();
      }
    } else if (isWriteContractError && !viewed.has('error')) {
      setDialogState({ isOpen: true, status: 'error', error: errorMessage });
    }
  }, [hash, isPending, isConfirming, isConfirmed, isWriteContractError, errorMessage, onConfirmed]);

  const onOpenChange = useCallback((open: boolean) => {
    if (!open) {
      setDialogState(prev => ({ ...prev, isOpen: false }));
      if (dialogState.hash) {
        viewedRef.current.add(dialogState.hash);
      } else if (dialogState.status === 'error') {
        viewedRef.current.add('error');
      }
    }
  }, [dialogState.hash, dialogState.status]);

  return {
    dialogState,
    onOpenChange,
    setDialogState,
  } as const;
}


