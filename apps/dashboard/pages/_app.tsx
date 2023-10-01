import {AppShell} from '@/components/AppShell';
import {TooltipsManagerProvider} from '@/volto/Tooltip';
import type {AppProps} from 'next/app';
import './_app.css';

export default function App({Component, pageProps}: AppProps) {
  return (
    <TooltipsManagerProvider>
      <AppShell>
        <Component {...pageProps} />
      </AppShell>
    </TooltipsManagerProvider>
  );
}
