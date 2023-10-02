import {
  AddressForm,
  useAddressFormState,
  type AddressFormState,
} from '@/components/AddressForm';
import {Address} from '@/services/address';
import type {MultiFamily} from '@/services/property';
import {Button, MiniTextButton} from '@/volto/Button';
import {Section} from '@/volto/Section';
import {CloseIcon, EditFilledIcon, LocationOnIcon} from '@/volto/icons';
import {useState} from 'react';
import useSWR from 'swr';
import useSWRMutation from 'swr/mutation';
import * as s from './PropertyInfo.css';

type PropertyInfoProps = {
  property: MultiFamily.Property;
};

export function PropertyInfo(props: PropertyInfoProps) {
  const {data: property, mutate: mutateProperty} = useSWR(
    `/api/properties/${props.property.id}`,
    async (endpoint): Promise<MultiFamily.Property> => {
      const res = await fetch(endpoint);
      return res.json();
    },
    {
      fallbackData: props.property,
    },
  );

  const [editing, setEditing] = useState(false);
  const exitEditing = () => setEditing(false);

  return (
    <Section
      title="Property info"
      actions={
        <MiniTextButton
          icon={editing ? <CloseIcon /> : <EditFilledIcon />}
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

function Viewing({property}: {property: MultiFamily.Property}) {
  const addr = Address.from(property.address);
  return (
    <p className={s.Address}>
      <span className={s.AddressIcon}>
        <LocationOnIcon />
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
  property: MultiFamily.Property;
  onCancel: () => void;
  onUpdateSettled: () => void;
  onUpdateSuccess: (updatedProperty: MultiFamily.Property) => void;
}) {
  const address = useAddressFormState(property.address);
  const updateProperty = useSWRMutation(
    `/api/properties/${property.id}`,
    async (endpoint): Promise<MultiFamily.Property> => {
      const d = toPropertyModel(address.fields);
      const res = await fetch(endpoint, {
        body: JSON.stringify(d),
        method: 'PATCH',
      });
      return res.json();
    },
    {
      onError(err) {
        console.error('Cannot update property', err);
        onUpdateSettled();
      },
      onSuccess(updatedProperty) {
        onUpdateSuccess(updatedProperty);
        onUpdateSettled();
      },
      revalidate: false,
    },
  );

  return (
    <>
      <AddressForm namePrefix="PropertyInfo" state={address} />
      <div className={s.EditActions}>
        <Button variant="text" onClick={onCancel}>
          Cancel
        </Button>
        <Button
          // TODO: disable button and show loading state
          disabled={updateProperty.isMutating}
          variant="contained"
          onClick={() => updateProperty.trigger()}
        >
          Save
        </Button>
      </div>
    </>
  );
}

function toPropertyModel(
  address: AddressFormState['fields'],
): Partial<MultiFamily.Property> {
  return {address};
}
