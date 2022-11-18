import type {AppProps} from 'next/app';
import AppShell from '@/components/AppShell';
import './_app.css';

export default function App({Component, pageProps}: AppProps) {
  return (
    <AppShell>
      <Component {...pageProps} />
    </AppShell>
  );
}
