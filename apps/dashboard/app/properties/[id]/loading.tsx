import {Skeleton} from '@/volto/Skeleton';
import type {Metadata} from 'next';

export const metadata = {
  title: 'Property Details - HausOps',
} satisfies Metadata;

// TODO: better loading state
export default function LoadingPropertyDetails() {
  return <Skeleton width="10rem" />;
}
