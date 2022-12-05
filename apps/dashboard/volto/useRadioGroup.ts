// A simplified copy of @react-aria/radio - useRadioGroup to avoid
// its issue with SSR due to randomly generated name.

import {useId} from 'react';

type RadioGroupProps = {
  name: string;
  isReadOnly?: boolean;
  isRequired?: boolean;
  isDisabled?: boolean;
};

type RadioGroup = {
  labelProps: {id: string};
  radioGroupProps: {
    role: 'radiogroup';
    'aria-labelledby': string;
    'aria-readonly'?: boolean;
    'aria-required'?: boolean;
    'aria-disabled'?: boolean;
  };
  radioProps: {name: string};
};

export function useRadioGroup({
  name,
  isReadOnly,
  isRequired,
  isDisabled,
}: RadioGroupProps): RadioGroup {
  const labelId = useId();
  return {
    labelProps: {
      // we dont' use htmlFor fieldId here because RadioGroup
      // is not an HTML input and should not be used with HTML label element.
      id: labelId,
    },
    radioGroupProps: {
      role: 'radiogroup',
      'aria-labelledby': labelId,
      'aria-readonly': isReadOnly || undefined,
      'aria-required': isRequired || undefined,
      'aria-disabled': isDisabled || undefined,
    },
    // pass to useRadio
    radioProps: {
      name,
    },
  };
}
