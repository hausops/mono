import type {PropertyType} from '@/services/property';
import {useRadioGroupState, type RadioGroupState} from '@/volto/Radio';
import {Section} from '@/volto/Section';
import {PropertyTypeSelector} from './PropertyTypeSelector';

import {
  SingleFamilyForm,
  useSingleFamilyFormState,
  type SingleFamilyFormState,
} from './SingleFamilyForm';

import {
  MultiFamilyForm,
  useMultiFamilyFormState,
  type MultiFamilyFormState,
} from './MultiFamilyForm/MultiFamilyForm';

export type DetailsFormState = {
  propertyType: RadioGroupState<PropertyType>;
  singleFamily: SingleFamilyFormState;
  multiFamily: MultiFamilyFormState;
};

export function DetailsForm({state}: {state: DetailsFormState}) {
  return (
    <Section title="Details">
      <PropertyTypeSelector state={state.propertyType} />
      {state.propertyType.selectedValue === 'single-family' ? (
        <SingleFamilyForm state={state.singleFamily} />
      ) : (
        <MultiFamilyForm state={state.multiFamily} />
      )}
    </Section>
  );
}

export function useDetailsFormState(): DetailsFormState {
  const propertyType = useRadioGroupState({
    initialValue: 'single-family',
  });
  const singleFamily = useSingleFamilyFormState();
  const multiFamily = useMultiFamilyFormState();
  return {propertyType, singleFamily, multiFamily};
}

/*
Property Info:
- year built
- heating (gas, wood, central, fireplace)
- a/c ?
- lot size
- garage
*/
