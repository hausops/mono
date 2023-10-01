import type {PropsWithChildren} from 'react';
import {PageLayout} from '@/layouts/Page';

export default function PropertiesLayout({children}: PropsWithChildren) {
  return <PageLayout>{children}</PageLayout>;
}

// TODO: move PageLayout here
