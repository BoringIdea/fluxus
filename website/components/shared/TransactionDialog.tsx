import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from '@/components/ui/alert-dialog';
import { getExplorerUrl } from '@/src/utils';

interface TransactionDialogProps {
  isOpen: boolean;
  onOpenChange: (open: boolean) => void;
  status: 'pending' | 'confirming' | 'confirmed' | 'error' | string;
  hash?: string;
  error?: string;
  title?: string;
  chainId: number;
}

export function TransactionDialog({
  isOpen,
  onOpenChange,
  status,
  hash,
  error,
  title,
  chainId,
}: TransactionDialogProps) {
  return (
    <AlertDialog open={isOpen} onOpenChange={onOpenChange}>
      <AlertDialogContent className="border-black/10 !bg-white text-[color:var(--text-primary)]">
        <AlertDialogHeader>
          <AlertDialogTitle className="font-primary text-[11px] uppercase tracking-[0.18em]">Transaction Status</AlertDialogTitle>
          <AlertDialogDescription className="whitespace-pre-wrap break-all text-[color:var(--text-secondary)]">
            {status === 'pending' && 'Transaction in progress...'}
            {status === 'confirming' && (
              <>
                Transaction sent, waiting for confirmation...
                <br />
                <a
                  href={`${getExplorerUrl(chainId)}/tx/${hash}`}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-[color:var(--color-primary)] hover:underline"
                >
                  View on Explorer: {hash}
                </a>
              </>
            )}
            {status === 'confirmed' && (
              <>
                {title || 'Transaction success'}
                <br />
                <a
                  href={`${getExplorerUrl(chainId)}/tx/${hash}`}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-[color:var(--color-primary)] hover:underline"
                >
                  View details on Explorer: {hash}
                </a>
              </>
            )}
            {status === 'error' && `Error: ${error}`}
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogAction 
            onClick={() => onOpenChange(false)}
            className="border border-black/10 bg-[color:var(--bg-muted)] font-primary text-[11px] uppercase tracking-[0.16em] text-[color:var(--text-primary)] hover:bg-[color:var(--bg-card-hover)]"
          >
            OK
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
}

export default TransactionDialog;
