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
      <AlertDialogContent className="bg-black text-white">
        <AlertDialogHeader>
          <AlertDialogTitle>Transaction Status</AlertDialogTitle>
          <AlertDialogDescription className="whitespace-pre-wrap break-all">
            {status === 'pending' && 'Transaction in progress...'}
            {status === 'confirming' && (
              <>
                Transaction sent, waiting for confirmation...
                <br />
                <a
                  href={`${getExplorerUrl(chainId)}/tx/${hash}`}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-blue-500 hover:underline"
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
                  className="text-blue-500 hover:underline"
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
            className="bg-green-400 hover:bg-green-300 text-black font-bold"
          >
            OK
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
}

export default TransactionDialog;


