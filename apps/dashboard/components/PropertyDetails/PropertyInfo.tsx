import {
  NumBathroomsSelector,
  NumBedroomsSelector,
} from '@/components/PropertyForm';
import {useFieldsState} from '@/components/useFieldsState';
import {Address, useAddressService} from '@/services/address';
import {
  RentalUnit,
  SingleFamilyProperty,
  usePropertyService,
} from '@/services/property';
import {Button, MiniTextButton} from '@/volto/Button';
import {Close, EditFilled} from '@/volto/icons';
import {Section} from '@/volto/Section';
import {Select, toOption} from '@/volto/Select';
import {TextField} from '@/volto/TextField';
import {useMemo, useState} from 'react';
import useSWR from 'swr';
import {Attribute, AttributeList} from './AttributeList';
import * as s from './PropertyInfo.css';

type PropertyInfoProps = {
  property: SingleFamilyProperty;
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
          onEditSuccess={(updatedProperty: SingleFamilyProperty) => {
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

function Viewing({property}: {property: SingleFamilyProperty}) {
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

type EditingUnitState = Omit<RentalUnit, 'size'> & {size?: string};

function Editing({
  property,
  onCancelClick,
  onEditSuccess,
}: {
  property: SingleFamilyProperty;
  onCancelClick: () => void;
  onEditSuccess: (updatedProperty: SingleFamilyProperty) => void;
}) {
  const namePrefix = 'PropertyInfo';
  const addressSvc = useAddressService();
  const propertySvc = usePropertyService();

  const address = useFieldsState(property.address);
  const unit = useFieldsState<EditingUnitState>({
    ...property.unit,
    size: property.unit.size ? `${property.unit.size}` : undefined,
  });

  return (
    <>
      <AttributeList className={s.EditingAttributeList}>
        <Attribute
          label="Address"
          value={
            <div className={s.AddressForm}>
              <TextField
                label="Street address"
                name={`${namePrefix}AddressLine1`}
                placeholder="200 Main St."
                value={address.fields.line1}
                onChange={(e) => address.updateField('line1', e.target.value)}
              />
              <TextField
                label="Apartment, suite, etc."
                name={`${namePrefix}AddressLine2`}
                value={address.fields.line2}
                onChange={(e) => address.updateField('line2', e.target.value)}
              />
              <TextField
                label="City"
                name={`${namePrefix}AddressCity`}
                value={address.fields.city}
                onChange={(e) => address.updateField('city', e.target.value)}
              />
              <Select
                label="State"
                name={`${namePrefix}AddressState`}
                options={useMemo(
                  () => addressSvc.getAllStates().map((s) => toOption(s.code)),
                  [addressSvc]
                )}
                value={address.fields.state}
                onChange={(e) => address.updateField('state', e.target.value)}
              />
              <TextField
                label="ZIP code"
                name={`${namePrefix}AddressZip`}
                value={address.fields.zip}
                onChange={(e) => address.updateField('zip', e.target.value)}
              />
            </div>
          }
        />
        <Attribute
          label="Beds"
          value={
            <NumBedroomsSelector
              name={`${namePrefix}Beds`}
              value={unit.fields.bedrooms}
              onChange={(e) => unit.updateField('bedrooms', +e.target.value)}
            />
          }
        />
        <Attribute
          label="Baths"
          value={
            <NumBathroomsSelector
              name={`${namePrefix}Baths`}
              value={unit.fields.bathrooms}
              onChange={(e) => unit.updateField('bathrooms', +e.target.value)}
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
  address: SingleFamilyProperty['address'],
  unit: EditingUnitState
): Partial<SingleFamilyProperty> {
  return {
    address,
    unit: {
      ...unit,
      size: unit.size ? +unit.size : undefined,
    },
  };
}
