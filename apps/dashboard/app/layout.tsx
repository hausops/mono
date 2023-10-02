import type {Metadata} from 'next';
import type {PropsWithChildren} from 'react';
import {AppShell} from './_internal/AppShell';
import './globals.css';

export const metadata = {
  title: 'Dashboard - HausOps',
  description: 'HausOps',
} satisfies Metadata;

export default function RootLayout(props: PropsWithChildren) {
  const {children} = props;
  return (
    <html lang="en">
      <body>
        <AppShell>{children}</AppShell>
      </body>
    </html>
  );
}
