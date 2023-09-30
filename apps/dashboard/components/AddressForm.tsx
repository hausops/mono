import {FieldsState, useFieldsState} from '@/components/useFieldsState';
import {AddressModel, useAddressService} from '@/services/address';
import {Select, toOption} from '@/volto/Select';
import {TextField} from '@/volto/TextField';
import {useMemo} from 'react';
import * as s from './AddressForm.css';

type AddressFormProps = {
  // layout dictate css to format the component layout
  layout?: 'full-width' | 'stacked';
  // namePrefix allows defining the prefix for field names
  // in order to make field names and ids unique on the page.
  namePrefix?: string;
  state: AddressFormState;
};

export type AddressFormState = FieldsState<AddressFields>;
type AddressFields = Required<AddressModel>;

export function AddressForm({
  layout = 'stacked',
  namePrefix = '',
  state,
}: AddressFormProps) {
  const addressSvc = useAddressService();
  const {fields, updateField} = state;

  const cs =
    layout === 'full-width'
      ? {
          root: s.layout.fourColumns,
          line1: s.gridColumnSpan[2],
          line2: s.gridColumnSpan[2],
          city: s.gridColumnSpan[2],
          state: s.gridColumnSpan[1],
          zip: s.gridColumnSpan[1],
        }
      : {
          root: s.layout.oneColumn,
        };

  return (
    <div className={cs.root}>
      <TextField
        className={cs.line1}
        label="Street address"
        name={`${namePrefix}AddressLine1`}
        placeholder="200 Main St."
        value={fields.line1}
        onChange={(e) => updateField('line1', e.target.value)}
      />
      <TextField
        className={cs.line2}
        label="Apartment, suite, etc."
        name={`${namePrefix}AddressLine2`}
        value={fields.line2}
        onChange={(e) => updateField('line2', e.target.value)}
      />
      <TextField
        className={cs.city}
        label="City"
        name={`${namePrefix}AddressCity`}
        value={fields.city}
        onChange={(e) => updateField('city', e.target.value)}
      />
      <Select
        className={cs.state}
        label="State"
        name={`${namePrefix}AddressState`}
        options={useMemo(
          () => addressSvc.getAllStates().map((s) => toOption(s.code)),
          [addressSvc],
        )}
        value={fields.state}
        onChange={(e) => updateField('state', e.target.value)}
      />
      <TextField
        className={cs.zip}
        label="ZIP code"
        name={`${namePrefix}AddressZip`}
        value={fields.zip}
        onChange={(e) => updateField('zip', e.target.value)}
      />
    </div>
  );
}

export function useAddressFormState(
  initialState?: AddressModel,
): AddressFormState {
  return useFieldsState({
    line1: '',
    line2: '',
    city: '',
    state: '',
    zip: '',
    ...initialState,
  });
}
