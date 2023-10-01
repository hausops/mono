import type {Metadata} from 'next';
import Link from 'next/link';
import type {PropsWithChildren} from 'react';

import {PageHeader} from '@/layouts/PageHeader';
import {Button} from '@/volto/Button';
import * as s from './_internal/PropertyList.css';

export const metadata = {
  title: 'Properties - HausOps',
} satisfies Metadata;

export default function PropertiesLayout({children}: PropsWithChildren) {
  return (
    <>
      {/* breadcrumb */}

      <PageHeader
        title="Properties"
        actions={
          <Button as={Link} variant="contained" href="/properties/new">
            Add property
          </Button>
        }
      />

      <ul className={s.PropertyList}>{children}</ul>
    </>
  );
}
