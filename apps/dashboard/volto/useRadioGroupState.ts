// A simplified copy of @react-stately/radio - useRadioGroupState to avoid
// its issue with SSR due to randomly generated name.

import {useState} from 'react';

export type RadioGroupProps<T> = {
  initialValue?: T;
  isDisabled?: boolean;
  isReadOnly?: boolean;
};

export type RadioGroupState<T> = {
  readonly selectedValue: T | null;
  readonly lastFocusedValue: T | null;
  readonly isReadOnly: boolean;
  readonly isDisabled: boolean;
  setSelectedValue(value: T | null): void;
  setLastFocusedValue(value: T | null): void;
};

export function useRadioGroupState<T extends string>(
  props: RadioGroupProps<T> = {}
): RadioGroupState<T> {
  const {initialValue, isReadOnly = false, isDisabled = false} = props;

  const [selectedValue, setSelected] = useState<T | null>(initialValue ?? null);
  const [lastFocusedValue, setLastFocusedValue] = useState<T | null>(null);

  function setSelectedValue(value: T): void {
    if (!isReadOnly && !isDisabled) {
      setSelected(value);
    }
  }

  return {
    selectedValue,
    lastFocusedValue,
    isReadOnly,
    isDisabled,
    setSelectedValue,
    setLastFocusedValue,
  };
}
