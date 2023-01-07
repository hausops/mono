import {useFieldsState} from '@/components/useFieldsState';
import {Address, useAddressService} from '@/services/address';
import {SingleFamilyProperty} from '@/services/property';
import {Button} from '@/volto/Button';
import {Close, EditFilled} from '@/volto/icons';
import {Section} from '@/volto/Section';
import {Select, toOption} from '@/volto/Select';
import {TextField} from '@/volto/TextField';
import {useMemo, useState} from 'react';
import {BathroomsSelect, BedroomsSelect} from '@/components/PropertyForm';
import {Attribute, AttributeList} from './AttributeList';
import * as s from './PropertyInfo.css';

type PropertyInfoProps = {
  property: SingleFamilyProperty;
};

export function PropertyInfo({property}: PropertyInfoProps) {
  const [editing, setEditing] = useState(false);
  return (
    <Section
      title="Property info"
      actions={
        <EditToggle editing={editing} onClick={() => setEditing(!editing)} />
      }
    >
      {editing ? (
        <Editing property={property} onCancelClick={() => setEditing(false)} />
      ) : (
        <Viewing property={property} />
      )}
    </Section>
  );
}

function EditToggle({
  editing,
  onClick,
}: {
  editing: boolean;
  onClick: () => void;
}) {
  return (
    <button className={s.EditToggle} onClick={onClick}>
      <span className={s.EditToggleIcon}>
        {editing ? <Close /> : <EditFilled />}
      </span>
      {editing ? 'Cancel' : 'Edit'}
    </button>
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
      {/* TODO: always render Attribute but show "-" if no data */}
      {unit.bedrooms ? <Attribute label="Beds" value={unit.bedrooms} /> : null}
      {unit.bathrooms ? (
        <Attribute label="Baths" value={unit.bathrooms} />
      ) : null}
      {unit.size ? (
        <Attribute
          label="Size"
          value={`${Intl.NumberFormat('en-US').format(unit.size)} sq.ft.`}
        />
      ) : null}
    </AttributeList>
  );
}
function Editing({
  property,
  onCancelClick,
}: {
  property: SingleFamilyProperty;
  onCancelClick: () => void;
}) {
  const namePrefix = 'PropertyInfo';
  const addressSvc = useAddressService();

  const address = useFieldsState(property.address);
  const unit = useFieldsState({
    ...property.unit,
    size: property.unit.size && `${property.unit.size}`,
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
            <BedroomsSelect
              name={`${namePrefix}Beds`}
              value={unit.fields.bedrooms}
              onChange={(e) => unit.updateField('bedrooms', +e.target.value)}
            />
          }
        />
        <Attribute
          label="Baths"
          value={
            <BathroomsSelect
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
        <Button variant="contained">Save</Button>
      </div>
    </>
  );
}
