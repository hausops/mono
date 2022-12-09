import {
  AddressForm,
  DetailsForm,
  useAddressFormState,
  useDetailsFormState,
} from '@/components/NewProperty';
import PageLayout from '@/layouts/Page';
import PageHeader from '@/layouts/PageHeader';
import {AddressServiceProvider, LocalAddressService} from '@/services/address';
import Button from '@/volto/Button';
import Head from 'next/head';
import Link from 'next/link';
import {useMemo} from 'react';
import * as s from './new.css';

export default function Page() {
  const addressSvc = useMemo(() => new LocalAddressService(), []);
  const addressForm = useAddressFormState();
  const detailsForm = useDetailsFormState();

  return (
    <AddressServiceProvider service={addressSvc}>
      <Head>
        <title>Add property - HausOps</title>
        <meta name="description" content="HausOps" />
        <link rel="icon" href="data:;base64,iVBORw0KGgo=" />
      </Head>

      <PageLayout>
        <PageHeader title="Add property" />
        <AddressForm namePrefix="Property" state={addressForm} />
        <DetailsForm state={detailsForm} />

        <div className={s.Actions}>
          <Button variant="text" as={Link} href="/properties">
            {/* Discard */}
            Cancel
          </Button>
          <Button
            variant="contained"
            onClick={() =>
              console.log({
                type: detailsForm.propertyType.selectedValue,
                address: addressForm,
                ...(detailsForm.propertyType.selectedValue === 'single-family'
                  ? detailsForm.singleFamily
                  : detailsForm.multiFamily),
              })
            }
          >
            Save
          </Button>
        </div>
      </PageLayout>
    </AddressServiceProvider>
  );
}
