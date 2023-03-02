import {
  AddressForm,
  AddressFormState,
  useAddressFormState,
} from '@/components/AddressForm';
import {Address} from '@/services/address';
import {MultiFamilyProperty, usePropertyService} from '@/services/property';
import {Button, MiniTextButton} from '@/volto/Button';
import {Close, EditFilled, LocationOn} from '@/volto/icons';
import {Section} from '@/volto/Section';
import {useState} from 'react';
import useSWR from 'swr';
import * as s from './PropertyInfo.css';

type PropertyInfoProps = {
  property: MultiFamilyProperty;
};

export function PropertyInfo(props: PropertyInfoProps) {
  const propertySvc = usePropertyService();
  const {data, mutate: mutateProperty} = useSWR(
    `/api/property/${props.property.id}`,
    async () => {
      const p = await propertySvc.get(props.property.id);
      return p?.type === 'multi-family' ? p : undefined;
    }
  );
  const property = data ?? props.property;

  const [editing, setEditing] = useState(false);
  const exitEditing = () => setEditing(false);

  return (
    <Section
      title="Property info"
      actions={
        <MiniTextButton
          icon={editing ? <Close /> : <EditFilled />}
          onClick={() => setEditing(!editing)}
        >
          {editing ? 'Cancel' : 'Edit'}
        </MiniTextButton>
      }
    >
      {editing ? (
        <Editing
          property={property}
          onCancel={exitEditing}
          onUpdateSettled={exitEditing}
          onUpdateSuccess={(updatedProperty) => {
            mutateProperty(updatedProperty, {revalidate: false});
          }}
        />
      ) : (
        <Viewing property={property} />
      )}
    </Section>
  );
}

function Viewing({property}: {property: MultiFamilyProperty}) {
  const addr = Address.from(property.address);
  return (
    <p className={s.Address}>
      <span className={s.AddressIcon}>
        <LocationOn />
      </span>
      {addr.toString()}
    </p>
  );
}

function Editing({
  property,
  onCancel,
  onUpdateSettled,
  onUpdateSuccess,
}: {
  property: MultiFamilyProperty;
  onCancel: () => void;
  onUpdateSettled: () => void;
  onUpdateSuccess: (updatedProperty: MultiFamilyProperty) => void;
}) {
  const propertySvc = usePropertyService();
  const address = useAddressFormState(property.address);
  return (
    <>
      <AddressForm namePrefix="PropertyInfo" state={address} />
      <div className={s.EditActions}>
        <Button variant="text" onClick={onCancel}>
          Cancel
        </Button>
        <Button
          variant="contained"
          // TODO: disable button and show loading state
          onClick={async () => {
            try {
              const d = toPropertyModel(address.fields);
              const updated = await propertySvc.update(property.id, d);
              onUpdateSuccess(updated);
            } finally {
              onUpdateSettled();
            }
          }}
        >
          Save
        </Button>
      </div>
    </>
  );
}

function toPropertyModel(
  address: AddressFormState['fields']
): Partial<MultiFamilyProperty> {
  return {address};
}
