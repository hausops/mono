import type {Metadata} from 'next';
import type {PropsWithChildren} from 'react';

import {PageHeader} from '@/layouts/PageHeader';

export const metadata = {
  title: 'Add property - HausOps',
} satisfies Metadata;

export default function NewPropertyLayout({children}: PropsWithChildren) {
  return (
    <>
      {/* breadcrumb */}
      <PageHeader title="Add property" />
      {children}
    </>
  );
}
