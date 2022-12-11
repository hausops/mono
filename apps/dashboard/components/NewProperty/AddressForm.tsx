import {FieldsState, useFieldsState} from '@/components/useFieldsState';
import {AddressModel, useAddressService} from '@/services/address';
import Select, {toOption} from '@/volto/Select';
import TextField from '@/volto/TextField';
import {useMemo} from 'react';
import * as s from './AddressForm.css';
import {Section} from './Section';

type AddressFormProps = {
  // namePrefix allows defining the prefix for field names
  // in order to make field names and ids unique on the page.
  namePrefix?: string;
  state: AddressFormState;
};

export type AddressFormState = FieldsState<AddressFields>;
type AddressFields = Required<AddressModel>;

export function AddressForm({namePrefix = '', state}: AddressFormProps) {
  const addressSvc = useAddressService();
  const fields = state.toJSON();
  return (
    <Section title="Address">
      <div className={s.Address}>
        <TextField
          className={s.gridColumnSpan[2]}
          label="Street address"
          name={`${namePrefix}AddressLine1`}
          placeholder="200 Main St."
          value={fields.line1}
          onChange={(e) => state.updateField('line1', e.target.value)}
        />
        <TextField
          className={s.gridColumnSpan[2]}
          label="Apartment, suite, etc."
          name={`${namePrefix}AddressLine2`}
          value={fields.line2}
          onChange={(e) => state.updateField('line2', e.target.value)}
        />
        <TextField
          className={s.gridColumnSpan[2]}
          label="City"
          name={`${namePrefix}AddressCity`}
          value={fields.city}
          onChange={(e) => state.updateField('city', e.target.value)}
        />
        <Select
          className={s.gridColumnSpan[1]}
          label="State"
          name={`${namePrefix}AddressState`}
          options={useMemo(
            () => addressSvc.getAllStates().map((s) => toOption(s.code)),
            [addressSvc]
          )}
          value={fields.state}
          onChange={(e) => state.updateField('state', e.target.value)}
        />
        <TextField
          className={s.gridColumnSpan[1]}
          label="ZIP code"
          name={`${namePrefix}AddressZip`}
          value={fields.zip}
          onChange={(e) => state.updateField('zip', e.target.value)}
        />
      </div>
    </Section>
  );
}

export function useAddressFormState(): AddressFormState {
  return useFieldsState(initialState);
}

const initialState: AddressFields = {
  line1: '',
  line2: '',
  city: '',
  state: '',
  zip: '',
};
