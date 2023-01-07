import {BathroomsSelect, BedroomsSelect} from '@/components/PropertyForm';
import {FieldsState, useFieldsState} from '@/components/useFieldsState';
import {TextField} from '@/volto/TextField';
import * as s from './SingleFamilyForm.css';

export type SingleFamilyFormState = FieldsState<{
  bedrooms?: number;
  bathrooms?: number;
  size: string;
  rentAmount: string;
}>;

export function SingleFamilyForm({state}: {state: SingleFamilyFormState}) {
  const {fields, updateField} = state;
  return (
    <div className={s.SingleFamilyForm}>
      <BedroomsSelect
        name="property.single.beds"
        value={fields.bedrooms}
        onChange={(e) => updateField('bedrooms', +e.target.value)}
      />
      <BathroomsSelect
        name="property.single.baths"
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
