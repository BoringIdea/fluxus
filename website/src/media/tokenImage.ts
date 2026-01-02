export function buildTokenMetadataUri(baseUri: string, tokenId: number | string) {
  return `${baseUri}/${tokenId}.json`;
}

export async function fetchTokenMetadataByUri<T = any>(tokenUri: string): Promise<T | null> {
  try {
    const res = await fetch(tokenUri);
    if (!res.ok) return null;
    return (await res.json()) as T;
  } catch {
    return null;
  }
}

export async function fetchTokenImageUrlByUri(tokenUri: string, fallback: string = '/fluxus-logo.png'): Promise<string> {
  const md = await fetchTokenMetadataByUri<any>(tokenUri);
  if (md && typeof md.image === 'string' && md.image.length > 0) {
    return md.image as string;
  }
  return fallback;
}

export async function fetchTokenImageUrl(baseUri: string, tokenId: number, fallback: string = '/fluxus-logo.png'): Promise<string> {
  const uri = buildTokenMetadataUri(baseUri, tokenId);
  return fetchTokenImageUrlByUri(uri, fallback);
}

export async function batchFetchTokenImageUrlsByUri(
  tokenUris: string[],
  concurrency: number = 8,
  fallback: string = '/fluxus-logo.png',
): Promise<Record<string, string>> {
  const uriQueue = [...tokenUris];
  const result: Record<string, string> = {};

  async function worker() {
    const next = uriQueue.shift();
    if (!next) return;
    try {
      result[next] = await fetchTokenImageUrlByUri(next, fallback);
    } catch {
      result[next] = fallback;
    }
    await worker();
  }

  await Promise.all(Array.from({ length: Math.min(concurrency, uriQueue.length) }, worker));
  return result;
}


