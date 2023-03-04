import {
  AddressForm,
  AddressFormState,
  useAddressFormState,
} from '@/components/AddressForm';
import {Attribute, AttributeList} from '@/components/AttributeList';
import {BathroomsSelect, BedroomsSelect} from '@/components/PropertyForm';
import {useFieldsState} from '@/components/useFieldsState';
import {Address} from '@/services/address';
import {SingleFamily, usePropertyService} from '@/services/property';
import {Button, MiniTextButton} from '@/volto/Button';
import {CloseIcon, EditFilledIcon} from '@/volto/icons';
import {Section} from '@/volto/Section';
import {TextField} from '@/volto/TextField';
import {useState} from 'react';
import useSWR from 'swr';
import * as s from './PropertyInfo.css';

type PropertyInfoProps = {
  property: SingleFamily.Property;
};

export function PropertyInfo(props: PropertyInfoProps) {
  const propertySvc = usePropertyService();
  const {data, mutate: mutateProperty} = useSWR(
    `/api/property/${props.property.id}`,
    async () => {
      const p = await propertySvc.get(props.property.id);
      return p?.type === 'single-family' ? p : undefined;
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
          onEditSuccess={(updatedProperty) => {
            mutateProperty(updatedProperty, {revalidate: false});
            exitEditing();
          }}
          onCancelClick={exitEditing}
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
  onCancelClick,
  onEditSuccess,
}: {
  property: SingleFamily.Property;
  onCancelClick: () => void;
  onEditSuccess: (updatedProperty: SingleFamily.Property) => void;
}) {
  const namePrefix = 'PropertyInfo';
  const propertySvc = usePropertyService();

  const address = useAddressFormState(property.address);
  const unit = useFieldsState<UnitFields>({
    ...property.unit,
    size: property.unit.size ? `${property.unit.size}` : '',
  });

  return (
    <>
      <AttributeList className={s.EditingAttributeList}>
        <Attribute
          label="Address"
          value={<AddressForm namePrefix={namePrefix} state={address} />}
        />
        <Attribute
          label="Beds"
          value={
            <BedroomsSelect
              name={`${namePrefix}Beds`}
              value={unit.fields.bedrooms}
              onChange={(selection) => unit.updateField('bedrooms', selection)}
            />
          }
        />
        <Attribute
          label="Baths"
          value={
            <BathroomsSelect
              name={`${namePrefix}Baths`}
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
              name={`${namePrefix}Size`}
              placeholder="Sq.ft."
              value={unit.fields.size}
              onChange={(e) => unit.updateField('size', e.target.value)}
            />
          }
        />
      </AttributeList>
      <div className={s.EditActions}>
        <Button variant="text" onClick={onCancelClick}>
          Cancel
        </Button>
        <Button
          variant="contained"
          // TODO: disable button and show loading state
          onClick={async () => {
            const d = toPropertyModel(address.fields, unit.fields);
            try {
              const updated = await propertySvc.update(property.id, d);
              console.log('property updated', updated);
              onEditSuccess(updated);
            } catch (err) {
              console.error('Cannot update property', err);
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
  address: AddressFormState['fields'],
  unit: UnitFields
): Partial<SingleFamily.Property> {
  return {
    address,
    unit: {
      ...unit,
      size: unit.size ? +unit.size : undefined,
    },
  };
}
