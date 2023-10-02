import type {Metadata} from 'next';

// Ideally, we want to keep the component here but due to:
// https://github.com/vanilla-extract-css/vanilla-extract/issues/1069,
//
// Vanilla Extract does not process files under dynamic routes.
//
// Note: nesting it in a subfolder e.g. (_internal) also doesn't work.
export {PropertyNotFound as default} from '@/components/PropertyDetails';
export const metadata = {
  title: 'Property not found - HausOps',
} satisfies Metadata;
