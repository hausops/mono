import type {AppProps} from 'next/app';
import AppShell from '@/components/AppShell';
import {AddressServiceProvider, LocalAddressService} from '@/services/address';
import {LocalPropertyService} from '@/services/property';
import './_app.css';
import {PropertyServiceProvider} from '@/services/property/PropertyServiceContext';

const addressSvc = new LocalAddressService();
const propertySvc = new LocalPropertyService();

export default function App({Component, pageProps}: AppProps) {
  return (
    <AddressServiceProvider service={addressSvc}>
      <PropertyServiceProvider service={propertySvc}>
        <AppShell>
          <Component {...pageProps} />
        </AppShell>
      </PropertyServiceProvider>
    </AddressServiceProvider>
  );
}
