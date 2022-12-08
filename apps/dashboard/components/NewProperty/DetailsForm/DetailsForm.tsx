import {PropertyType} from '@/services/property';
import {RadioGroupState, useRadioGroupState} from '@/volto/useRadioGroupState';
import {Section} from '../Section';
import {PropertyTypeSelector} from './PropertyTypeSelector';

import {
  SingleFamilyForm,
  SingleFamilyFormState,
  useSingleFamilyFormState,
} from './SingleFamilyForm';

import {
  MultiFamilyForm,
  MultiFamilyFormState,
  useMultiFamilyFormState,
} from './MultiFamilyForm/MultiFamilyForm';

type DetailsFormState = {
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
