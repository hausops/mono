import {
  AddressForm,
  AddressFormState,
  useAddressFormState,
} from '@/components/AddressForm';
import {
  DetailsForm,
  DetailsFormState,
  useDetailsFormState,
} from '@/components/NewProperty';
import {PageLayout} from '@/layouts/Page';
import {PageHeader} from '@/layouts/PageHeader';
import {NewPropertyData, usePropertyService} from '@/services/property';
import {Button} from '@/volto/Button';
import {Section} from '@/volto/Section';
import Head from 'next/head';
import Link from 'next/link';
import {useRouter} from 'next/router';
import useSWRMutation from 'swr/mutation';
import * as s from './new.css';

export default function Page() {
  const router = useRouter();
  const propertySvc = usePropertyService();

  const address = useAddressFormState();
  const details = useDetailsFormState();

  const addPropertyMutation = useSWRMutation('/api/properties', () => {
    const d = toNewPropertyData(address.fields, details);
    return propertySvc.add(d);
  });

  return (
    <>
      <Head>
        <title>Add property - HausOps</title>
        <meta name="description" content="HausOps" />
        <link rel="icon" href="data:;base64,iVBORw0KGgo=" />
      </Head>

      <PageLayout>
        <PageHeader title="Add property" />
        <Section title="Address">
          <AddressForm
            layout="full-width"
            namePrefix="Property"
            state={address}
          />
        </Section>
        <DetailsForm state={details} />

        <div className={s.Actions}>
          <Button variant="text" as={Link} href="/properties">
            {/* Discard */}
            Cancel
          </Button>
          <Button
            variant="contained"
            disabled={addPropertyMutation.isMutating}
            onClick={async () => {
              try {
                const created = await addPropertyMutation.trigger();
                console.log('property created', created);
                router.push('/properties');
              } catch (err) {
                console.error('Cannot add property', err);
              }
            }}
          >
            Save
          </Button>
        </div>
      </PageLayout>
    </>
  );
}

function toNewPropertyData(
  address: AddressFormState['fields'],
  details: DetailsFormState
): NewPropertyData {
  // TODO: validate required
  const propertyType = details.propertyType.selectedValue ?? 'single-family';

  if (propertyType === 'single-family') {
    const unit = details.singleFamily.fields;
    return {
      type: propertyType,
      address,
      unit: {
        ...unit,
        size: stringInputToNumber(unit.size),
        rentAmount: stringInputToNumber(unit.rentAmount),
      },
    };
  }

  const {units} = details.multiFamily;
  return {
    type: propertyType,
    address,
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
