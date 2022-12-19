import {FieldsState, useFieldsState} from '@/components/useFieldsState';
import {Select, toOption} from '@/volto/Select';
import {TextField} from '@/volto/TextField';
import * as s from './SingleFamilyForm.css';

export type SingleFamilyFormState = FieldsState<{
  bedrooms?: number;
  bathrooms?: number;
  size: string;
  rentAmount: string;
}>;

// TODO: refactor
const bedsOptions = [
  {label: 'Studio', value: 0},
  ...[1, 2, 3, 4, 5].map(toOption),
];

// TODO: refactor
const bathsOptions = [
  {label: 'None', value: 0},
  ...[1, 1.5, 2, 2.5, 3, 3.5, 4].map(toOption),
];

export function SingleFamilyForm({state}: {state: SingleFamilyFormState}) {
  const {fields, updateField} = state;
  return (
    <div className={s.SingleFamilyForm}>
      <Select
        label="Beds"
        name="property.single.beds"
        options={bedsOptions}
        value={fields.bedrooms}
        onChange={(e) => updateField('bedrooms', +e.target.value)}
      />
      <Select
        label="Baths"
        name="property.single.baths"
        options={bathsOptions}
        value={fields.bathrooms}
        onChange={(e) => updateField('bathrooms', +e.target.value)}
      />
      <TextField
        type="number"
        label="Size"
        name="property.single.size"
        placeholder="Sq.ft."
        value={fields.size}
        onChange={(e) => updateField('size', e.target.value)}
      />
      <TextField
        type="number"
        label="Rent"
        name="property.single.rent"
        value={fields.rentAmount}
        onChange={(e) => updateField('rentAmount', e.target.value)}
      />
    </div>
  );
}

const initialState = {
  size: '',
  rentAmount: '',
};

export function useSingleFamilyFormState(): SingleFamilyFormState {
  return useFieldsState(initialState);
}
