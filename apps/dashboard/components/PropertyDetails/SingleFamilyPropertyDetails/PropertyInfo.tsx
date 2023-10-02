import {
  AddressForm,
  useAddressFormState,
  type AddressFormState,
} from '@/components/AddressForm';
import {Attribute, AttributeList} from '@/components/AttributeList';
import {BathroomsSelect, BedroomsSelect} from '@/components/PropertyForm';
import {useFieldsState} from '@/components/useFieldsState';
import {Address} from '@/services/address';
import type {SingleFamily} from '@/services/property';
import {Button, MiniTextButton} from '@/volto/Button';
import {Section} from '@/volto/Section';
import {TextField} from '@/volto/TextField';
import {CloseIcon, EditFilledIcon} from '@/volto/icons';
import {useState} from 'react';
import useSWRMutation from 'swr/mutation';
import * as s from './PropertyInfo.css';

type PropertyInfoProps = {
  property: SingleFamily.Property;
  onUpdateSuccess: (updatedProperty: SingleFamily.Property) => void;
};

export function PropertyInfo({property, onUpdateSuccess}: PropertyInfoProps) {
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
          onUpdateSuccess={onUpdateSuccess}
        />
      ) : (
        <Viewing property={property} />
      )}
    </Section>
  );
}

function Viewing({property}: {property: SingleFamily.Property}) {
  const {address, unit} = property;
  const [street, region] = Address.from(address).format();
  return (
    <AttributeList>
      <Attribute
        label="Address"
        value={
          <>
            <p>{street}</p>
            <p>{region}</p>
          </>
        }
      />
      <Attribute label="Beds" value={unit.bedrooms} />
      <Attribute label="Baths" value={unit.bathrooms} />
      <Attribute
        label="Size"
        value={
          unit.size == null
            ? null
            : `${Intl.NumberFormat('en-US').format(unit.size)} sq.ft.`
        }
      />
    </AttributeList>
  );
}

type UnitFields = Omit<SingleFamily.Unit, 'size'> & {size: string};

function Editing({
  property,
  onCancel,
  onUpdateSettled,
  onUpdateSuccess,
}: {
  property: SingleFamily.Property;
  onCancel: () => void;
  onUpdateSettled: () => void;
  onUpdateSuccess: (updatedProperty: SingleFamily.Property) => void;
}) {
  const address = useAddressFormState(property.address);
  const unit = useFieldsState<UnitFields>({
    ...property.unit,
    size: property.unit.size ? `${property.unit.size}` : '',
  });

  const updateProperty = useSWRMutation(
    `/api/properties/${property.id}`,
    async (endpoint): Promise<SingleFamily.Property> => {
      const d = toPropertyModel(address.fields, unit.fields);
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
      <AttributeList className={s.EditingAttributeList}>
        <Attribute
          label="Address"
          value={<AddressForm namePrefix="PropertyInfo" state={address} />}
        />
        <Attribute
          label="Beds"
          value={
            <BedroomsSelect
              value={unit.fields.bedrooms}
              onChange={(selection) => unit.updateField('bedrooms', selection)}
            />
          }
        />
        <Attribute
          label="Baths"
          value={
            <BathroomsSelect
              value={unit.fields.bathrooms}
              onChange={(selection) => unit.updateField('bathrooms', selection)}
            />
          }
        />
        <Attribute
          label="Size"
          value={
            <TextField
              type="number"
              label="Size"
              placeholder="Sq.ft."
              value={unit.fields.size}
              onChange={(e) => unit.updateField('size', e.target.value)}
            />
          }
        />
      </AttributeList>
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
  unit: UnitFields,
): Partial<SingleFamily.Property> {
  return {
    address,
    unit: {
      ...unit,
      size: unit.size ? +unit.size : undefined,
    },
  };
}
