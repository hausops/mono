import {
  AddressForm,
  AddressFormState,
  DetailsForm,
  DetailsFormState,
  useAddressFormState,
  useDetailsFormState,
} from '@/components/NewProperty';
import {PageLayout} from '@/layouts/Page';
import {PageHeader} from '@/layouts/PageHeader';
import {PropertyData, usePropertyService} from '@/services/property';
import {Button} from '@/volto/Button';
import Head from 'next/head';
import Link from 'next/link';
import {useRouter} from 'next/router';
import * as s from './new.css';

export default function Page() {
  const router = useRouter();
  const propertySvc = usePropertyService();

  const addressForm = useAddressFormState();
  const detailsForm = useDetailsFormState();

  return (
    <>
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
            // TODO: validation
            onClick={async () => {
              const d = toPropertyData(addressForm, detailsForm);
              const created = await propertySvc.add(d);
              console.log('property created', created);
              router.push('/properties');
            }}
          >
            Save
          </Button>
        </div>
      </PageLayout>
    </>
  );
}

function toPropertyData(
  addressForm: AddressFormState,
  detailsForm: DetailsFormState
): PropertyData {
  // TODO: validate required
  const propertyType =
    detailsForm.propertyType.selectedValue ?? 'single-family';
  const address = addressForm.fields;

  if (propertyType === 'single-family') {
    const unit = detailsForm.singleFamily.fields;
    return {
      type: propertyType,
      address: address,
      ...unit,
      size: stringInputToNumber(unit.size),
      rentAmount: stringInputToNumber(unit.rentAmount),
    };
  }

  const {units} = detailsForm.multiFamily;
  return {
    type: propertyType,
    address: address,
    units: units.map((unit) => ({
      ...unit,
      size: stringInputToNumber(unit.size),
      rentAmount: stringInputToNumber(unit.rentAmount),
    })),
  };
}

// convert input state (string) to a number, return undefined for an empty string
function stringInputToNumber(str: string): number | undefined {
  return str ? +str : undefined;
}
