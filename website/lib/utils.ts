import { Collection, CollectionMetadata } from "@/src/api";
import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export const fetchCollectionMetadata = async (collection: Collection) => {
  if (collection?.base_uri) {
    try {
      const response = await fetch(collection.base_uri + '/collection.json');
      if (!response.ok) {
        throw new Error(`Failed to fetch collection metadata: ${response.status}`);
      }
      const metadata: CollectionMetadata = await response.json();
      return metadata;
    } catch (error) {
      console.error('Error fetching collection metadata:', error);
      return null;
    }
  }
  return null;
}

export const fetchCollectionImage = async (collection: Collection) => {
  if (collection?.base_uri) {
    try {
      const response = await fetch(collection.base_uri + '/collection.json');
      if (!response.ok) {
        throw new Error(`Failed to fetch collection metadata: ${response.status}`);
      }
      const metadata: CollectionMetadata = await response.json();
      return metadata.image;
    } catch (error) {
      console.error('Error fetching collection image:', error);
      return null;
    }
  }
  return null;
}

export const fetchCollectionBanner = async (collection: Collection) => {
  if (collection?.base_uri) {
    try {
      const response = await fetch(collection.base_uri + '/collection.json');
      if (!response.ok) {
        throw new Error(`Failed to fetch collection metadata: ${response.status}`);
      }
      const metadata: CollectionMetadata = await response.json();
      return metadata.banner_image;
    } catch (error) {
      console.error('Error fetching collection banner:', error);
      return null;
    }
  }
  return null;
}