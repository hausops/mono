import {FieldsState, useFieldsState} from '@/components/useFieldsState';
import Select, {toOption} from '@/volto/Select';
import TextField from '@/volto/TextField';
import * as s from './SingleFamilyForm.css';

export type SingleFamilyFormState = FieldsState<{
  bedrooms?: number;
  bathrooms?: number;
  size?: number;
  rentAmount?: number;
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
  return (
    <div className={s.SingleFamilyForm}>
      <Select
        label="Beds"
        name="property.single.beds"
        options={bedsOptions}
        value={state.bedrooms}
        onChange={(e) => state.updateField('bedrooms', +e.target.value)}
      />
      <Select
        label="Baths"
        name="property.single.baths"
        options={bathsOptions}
        value={state.bathrooms}
        onChange={(e) => state.updateField('bathrooms', +e.target.value)}
      />
      <TextField
        type="number"
        label="Size"
        name="property.single.size"
        placeholder="Sq.ft."
        value={state.size}
        onChange={(e) => state.updateField('size', +e.target.value)}
      />
      <TextField
        type="number"
        label="Rent"
        name="property.single.rent"
        value={state.rentAmount}
        onChange={(e) => state.updateField('rentAmount', +e.target.value)}
      />
    </div>
  );
}

export function useSingleFamilyFormState(): SingleFamilyFormState {
  return useFieldsState({});
}
