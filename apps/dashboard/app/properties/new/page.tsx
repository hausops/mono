'use client';

import {
  AddressForm,
  useAddressFormState,
  type AddressFormState,
} from '@/components/AddressForm';
import {propertySvc, type NewPropertyData} from '@/services/property';
import {Button} from '@/volto/Button';
import {Section} from '@/volto/Section';
import Link from 'next/link';
import {useRouter} from 'next/navigation';
import useSWRMutation from 'swr/mutation';
import {
  DetailsForm,
  useDetailsFormState,
  type DetailsFormState,
} from './_internal/DetailsForm';
import * as s from './page.css';

export default function NewPropertyPage() {
  const router = useRouter();

  const address = useAddressFormState();
  const details = useDetailsFormState();

  const addPropertyMutation = useSWRMutation('/api/properties', () => {
    const d = toNewPropertyData(address.fields, details);
    return propertySvc.add(d);
  });

  return (
    <>
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
    </>
  );
}

function toNewPropertyData(
  address: AddressFormState['fields'],
  details: DetailsFormState,
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
