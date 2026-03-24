import { useEffect, useState } from 'react';
import Header from '@/components/Header';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import {
  Form,
} from "@/components/ui/form"
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import * as z from 'zod';
import {
  useReadContract,
  useWriteContract,
  useWaitForTransactionReceipt,
  useChainId
} from "wagmi";
import { FACTORY_CONTRACT_ABI, getFactoryContractAddress } from '@/src/contract';
import { parseEther } from 'ethers';
import { getChainSymbol, isHideHomeAndLaunchPage } from "@/src/utils";
import TransactionDialog from '@/components/shared/TransactionDialog';
import { useTransactionDialog } from '@/src/hooks/useTransactionDialog';
import ComingSoon from '@/components/CommingSoon';
import PriceChart from '@/components/Collection/PriceChart';
import Link from 'next/link';

const formSchema = z.object({
  name: z.string().min(1, 'Enter the name'),
  symbol: z.string().min(1, 'Enter the symbol'),
  maxSupply: z.string().min(1, 'Enter the max supply'),
  initialPrice: z.string().min(1, 'Enter the initial price'),
  maxPrice: z.string().default('0'),
  creatorFee: z.number()
    .min(0, 'Creator fee cannot be negative')
    .max(100, 'Creator fee cannot exceed 100')
    .multipleOf(0.1, 'Creator fee must be a multiple of 0.1'),
  uri: z.string().min(1, 'Enter the URI'),
  supportCrossChain: z.boolean().default(false),
});

// Set input style
const inputClassName =
  "h-12 w-full border border-black/10 bg-[color:var(--bg-surface)] px-3 font-primary text-[11px] tracking-[0.02em] text-[color:var(--text-secondary)] outline-none transition-colors placeholder:normal-case placeholder:tracking-normal placeholder:text-[color:var(--text-muted)]/70 focus:border-black/20";

// Set label style
const labelClassName =
  "mb-2 block font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)]";

// Cross-chain default values (hardcoded as requested)
const DEFAULT_GATEWAY_ADDRESS = "0x0c487a766110c85d301d96e33579c5b317fa4995";
const DEFAULT_GAS_LIMIT = BigInt(12000000);
const DEFAULT_SUPPORT_MINT = true;

export default function CreateCollection() {
  const chainId = useChainId();
  const factoryContractAddress = getFactoryContractAddress(chainId as any);

  const [contractAddress, setContractAddress] = useState<string | null>(null);
  const [showContractAddress, setShowContractAddress] = useState<boolean>(false);
  const { data: hash, writeContract, isError: isWriteContractError, error: writeContractError, isPending } = useWriteContract();
  const [name, setName] = useState<string>('');
  const [symbol, setSymbol] = useState<string>('');
  const [maxSupply, setMaxSupply] = useState<string>('0');
  const [initialPrice, setInitialPrice] = useState<string>('0');
  const [creatorFee, setCreatorFee] = useState<string>('0');
  const [uri, setUri] = useState<string>('');
  const [isCompleted, setIsCompleted] = useState<boolean>(false);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    mode: 'onChange',
    defaultValues: {
      name: '',
      symbol: '',
      maxSupply: '',
      initialPrice: '',
      maxPrice: '0',
      creatorFee: 0,
      uri: '',
      supportCrossChain: false,
    },
  });
  const { watch } = form;
  const watchedValues = watch();
  
  // Calculate expected address based on cross-chain support
  const { data: expectAddress, refetch: refetchExpectAddress } = useReadContract({
    // @ts-ignore
    address: factoryContractAddress,
    abi: FACTORY_CONTRACT_ABI,
    functionName: watchedValues.supportCrossChain ? 'calculateFluxusCrossChainAddress' : 'calculateFluxusAddress',
    args: isCompleted ? (
      watchedValues.supportCrossChain ? [
        name,
        symbol,
        initialPrice,
        maxSupply,
        parseEther(watchedValues.maxPrice || "0"), // maxPrice from form
        creatorFee,
        uri,
        DEFAULT_GATEWAY_ADDRESS,
        DEFAULT_GAS_LIMIT.toString(),
        DEFAULT_SUPPORT_MINT
      ] : [
        name,
        symbol,
        initialPrice,
        maxSupply,
        parseEther(watchedValues.maxPrice || "0"), // maxPrice from form
        creatorFee,
        uri
      ]
    ) : undefined,
  });

  useEffect(() => {
    if (expectAddress && isCompleted) {
      setContractAddress(expectAddress as string);
      setShowContractAddress(true);
    }
  }, [expectAddress, isCompleted]);

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    if (!factoryContractAddress) {
      console.error('Factory contract address is not set');
      return;
    }

    // Convert all parameters to the same format
    const params = {
      name: values.name,
      symbol: values.symbol,
      initialPrice: parseEther(values.initialPrice.toString()),
      maxSupply: BigInt(values.maxSupply),
      maxPrice: parseEther(values.maxPrice || "0"), // Using user input or 0 as default
      creatorFee: parseEther((Number(values.creatorFee) / 100).toString()),
      uri: values.uri,
      supportCrossChain: values.supportCrossChain,
    };

    console.log('Submit params:', {
      raw: values,
      processed: {
        name: params.name,
        symbol: params.symbol,
        initialPrice: params.initialPrice.toString(),
        maxSupply: params.maxSupply.toString(),
        maxPrice: params.maxPrice.toString(),
        creatorFee: params.creatorFee.toString(),
        uri: params.uri,
        supportCrossChain: params.supportCrossChain,
      }
    });

    setName(params.name);
    setSymbol(params.symbol);
    setMaxSupply(params.maxSupply.toString());
    setInitialPrice(params.initialPrice.toString());
    setCreatorFee(params.creatorFee.toString());
    setUri(params.uri);

    // Call different functions based on cross-chain support
    if (params.supportCrossChain) {
      writeContract({
        // @ts-ignore
        address: factoryContractAddress,
        abi: FACTORY_CONTRACT_ABI,
        functionName: 'createFluxusCrossChain',
        args: [
          params.name,
          params.symbol,
          params.initialPrice.toString(),
          params.maxSupply.toString(),
          params.maxPrice.toString(),
          params.creatorFee.toString(),
          params.uri,
          DEFAULT_GATEWAY_ADDRESS,
          DEFAULT_GAS_LIMIT.toString(),
          DEFAULT_SUPPORT_MINT
        ],
      });
    } else {
      writeContract({
        // @ts-ignore
        address: factoryContractAddress,
        abi: FACTORY_CONTRACT_ABI,
        functionName: 'createFluxus',
        args: [
          params.name,
          params.symbol,
          params.initialPrice.toString(),
          params.maxSupply.toString(),
          params.maxPrice.toString(),
          params.creatorFee.toString(),
          params.uri
        ],
      });
    }

    // Update state immediately and refetch address
    setIsCompleted(true);
    refetchExpectAddress();
  };

  // Dialog state
  const { isLoading: isConfirming, isSuccess: isConfirmed } = useWaitForTransactionReceipt({ hash });
  const { dialogState, onOpenChange } = useTransactionDialog({
    hash,
    isPending,
    isConfirming,
    isConfirmed,
    isWriteContractError,
    writeContractError,
    onConfirmed: () => setShowContractAddress(true),
  });
  const s = watch('maxSupply');
  const i = watch('initialPrice');
  console.log('maxSupply---', s, i);
  return (
    isHideHomeAndLaunchPage ? (
      <ComingSoon />
    ) : (
      <div className="min-h-screen w-full bg-transparent text-primary">
        <div className="mx-auto max-w-6xl space-y-8 px-4 pb-20 pt-16 sm:px-10">
          <div className="flex flex-col gap-4 border-b border-black/10 pb-5 sm:flex-row sm:items-center sm:justify-between">
            <div>
              <p className="flux-kicker mb-3">Launchpad</p>
              <h1 className="flux-title text-[clamp(2.25rem,4.2vw,3.75rem)]">Create Collection</h1>
            </div>
            <Link href={'/docs/creation-guide'}>
              <button className="inline-flex h-10 items-center gap-2 border border-black/10 bg-[color:var(--bg-surface)] px-4 font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-primary)] transition-colors hover:bg-[color:var(--bg-muted)]">
                <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 20 20" fill="none">
                  <path d="M17.1328 15.5918H12.9199L10.6289 18.0488C10.2793 18.3984 9.71484 18.4004 9.36328 18.0508L9.36133 18.0488L7.06055 15.5918H2.85156C2.11328 15.5918 1.51367 15.1582 1.51367 14.4199V3.09375C1.51367 2.35547 2.11133 1.75781 2.84961 1.75586H17.1348C17.873 1.75586 18.4727 2.35547 18.4727 3.09375V14.4199C18.4707 15.1602 17.873 15.5918 17.1328 15.5918Z" fill="currentColor" />
                </svg>
                Creation Guide
              </button>
            </Link>
          </div>

          <div className="border border-black/10 bg-[color:var(--bg-surface)]">
            <div className="flex flex-col md:flex-row">
              <div className="w-full border-b border-black/10 px-4 py-8 md:w-1/2 md:border-b-0 md:border-r sm:px-8">
                <Card className="bg-transparent border-none shadow-none">
                  <CardContent>
                    <Form {...form}>
                      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6">
                        {contractAddress && showContractAddress && (
                          <div className="border border-black/10 bg-[color:var(--bg-muted)] p-4">
                            <div className="flex items-center justify-between font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)]">
                              <span>Contract Address</span>
                              <button
                                onClick={(e) => {
                                  e.preventDefault();
                                  e.stopPropagation();
                                  navigator.clipboard.writeText(contractAddress);
                                }}
                                type="button"
                                className="border border-black/10 bg-[color:var(--bg-surface)] px-3 py-1 text-[10px] tracking-[0.16em] text-[color:var(--color-primary)]"
                              >
                                Copy
                              </button>
                            </div>
                            <p className="mt-2 break-all font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-primary)]">{contractAddress}</p>
                          </div>
                        )}

                          {/* Name and Symbol */}
                          <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 sm:gap-6">
                            <div>
                              <label className={labelClassName}>
                                Name
                              </label>
                              <input {...form.register('name')} className={inputClassName} placeholder="Collection Name" />
                            </div>
                            <div>
                              <label className={labelClassName}>
                                Symbol
                              </label>
                              <input {...form.register('symbol')} className={inputClassName} placeholder="e.g. Fluxus" />
                            </div>
                          </div>

                          {/* Price and Max Supply */}
                          <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 sm:gap-6">
                            <div>
                              <label className={labelClassName}>
                                Initial Price ({getChainSymbol(chainId)})
                              </label>
                              <input
                                type="number"
                                step="0.001"
                                {...form.register('initialPrice')}
                                className={inputClassName}
                                placeholder="0.001"
                              />
                            </div>
                            <div>
                              <label className={labelClassName}>
                                Max Supply
                              </label>
                              <input
                                type="number"
                                {...form.register('maxSupply')}
                                className={inputClassName}
                                placeholder="10000"
                              />
                            </div>
                          </div>

                          {/* Creator Fee and Max Price */}
                          <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 sm:gap-6">
                            <div>
                              <label className={labelClassName}>
                                Creator Fee (%)
                              </label>
                              <input
                                type="number"
                                {...form.register('creatorFee', {
                                  min: 0,
                                  max: 100,
                                  valueAsNumber: true,
                                })}
                                className={inputClassName}
                                placeholder="10"
                                step="0.1"
                                min="0"
                                max="100"
                              />
                            </div>
                            <div>
                              <label className={`${labelClassName} flex items-center justify-between gap-2`}>
                                <span>Max Price ({getChainSymbol(chainId)})</span>
                              </label>
                              <input
                                type="number"
                                step="0.001"
                                {...form.register('maxPrice')}
                                className={inputClassName}
                                placeholder="0"
                              />
                            </div>
                          </div>

                          <div>
                            <label className={labelClassName}>Base URI</label>
                            <input {...form.register('uri')} className={inputClassName} placeholder="https://..." />
                          </div>

                          <div className="space-y-3 border border-black/10 bg-[color:var(--bg-muted)] p-4">
                            <div className="flex items-center justify-between">
                              <div>
                                <p className="font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)]">Cross-Chain Support</p>
                                <p className="mt-2 text-sm text-[color:var(--text-secondary)]">
                                  Enable cross-chain functionality for your NFT collection
                                </p>
                              </div>
                              <label className="relative inline-flex cursor-pointer items-center">
                                <input type="checkbox" {...form.register('supportCrossChain')} className="sr-only peer" />
                                <span className="relative h-7 w-14 rounded-full border border-black/10 bg-[color:var(--bg-surface)] shadow-[inset_0_0_0_1px_rgba(17,24,39,0.03)] transition-all after:absolute after:left-1 after:top-1 after:h-5 after:w-5 after:rounded-full after:border after:border-black/10 after:bg-white after:shadow-sm after:transition-all after:content-[''] peer-checked:border-[#16A34A] peer-checked:bg-[#16A34A] peer-checked:after:translate-x-7 peer-checked:after:border-white/60 peer-checked:after:bg-white" />
                              </label>
                            </div>
                            <div className="space-y-1 font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-muted)]">
                              <p>Gateway: {DEFAULT_GATEWAY_ADDRESS}</p>
                              <p>Gas Limit: {DEFAULT_GAS_LIMIT.toString()}</p>
                              <p>Support Mint: {DEFAULT_SUPPORT_MINT ? 'Enabled' : 'Disabled'}</p>
                            </div>
                          </div>

                          <div className="border border-black/10 bg-[color:var(--bg-muted)] px-4 py-3 font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-secondary)]">
                            Verify your BaseURI is accessible; otherwise the collection will not display.
                          </div>

                          <button
                            type="submit"
                            disabled={isPending || !form.formState.isValid}
                            className={`flex h-12 w-full items-center justify-center gap-2 border font-primary text-[11px] uppercase tracking-[0.18em] ${
                              isPending || !form.formState.isValid
                                ? 'cursor-not-allowed border-black/10 bg-[color:var(--bg-muted)] text-[color:var(--text-muted)]'
                                : 'border-[#16A34A] bg-[#16A34A] text-white hover:border-[#15803D] hover:bg-[#15803D]'
                            }`}
                          >
                            {isPending ? 'Creating...' : 'Create Collection'}
                          </button>
                      </form>
                    </Form>
                    <TransactionDialog
                      isOpen={dialogState.isOpen}
                      onOpenChange={onOpenChange}
                      status={dialogState.status}
                      hash={dialogState.hash}
                      error={dialogState.error}
                      title="NFT Collection Created Successfully"
                      chainId={chainId}
                    />
                  </CardContent>
                </Card>
              </div>

              <div className="w-full px-4 py-8 md:w-1/2 sm:px-8">
                <div className="mb-6 border-b border-black/10 pb-4">
                  <p className="flux-kicker mb-2">Preview</p>
                  <h2 className="flux-h2">Price Curve</h2>
                  <p className="mt-3 text-sm text-[color:var(--text-secondary)]">
                    Based on the set price and supply, the generated price curve
                  </p>
                </div>

                <div className="flex h-[640px] flex-col overflow-hidden border border-black/10 bg-[color:var(--bg-muted)] p-4 sm:mb-40">
                  {Number(s) > 0 && Number(i) > 0 ? (
                    <div className="relative w-full flex-1">
                      <div className="absolute inset-0 overflow-hidden pr-2 pb-4">
                        <PriceChart
                          collection={{
                            max_supply: Number(s),
                            initial_price: parseEther(i).toString(),
                            current_supply: 0,
                            floor_price: parseEther(i).toString(),
                          }}
                          chainId={chainId}
                        />
                      </div>
                    </div>
                  ) : (
                    <div className="flex h-full flex-col items-center justify-center px-4 text-center text-[color:var(--text-muted)]">
                      <p>[Fill the form on the left to preview the curve]</p>
                    </div>
                  )}
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    )
  )
}
