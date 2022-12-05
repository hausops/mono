import {ChangeEvent} from 'react';
import {RadioGroupState} from './useRadioGroupState';

type RadioProps<T> = {
  name: string;
  value: T;
  isDisabled?: boolean;
};

type Radio<T> = {
  inputProps: {
    checked: boolean;
    disabled: boolean;
    name: string;
    onChange: (event: ChangeEvent<HTMLInputElement>) => void;
    type: 'radio';
    value: T;
  };
  isDisabled: boolean;
  isSelected: boolean;
};

export function useRadio<T extends string>(
  props: RadioProps<T>,
  state: RadioGroupState<T>
): Radio<T> {
  const {name, value, isDisabled = state.isDisabled} = props;
  const isSelected = state.selectedValue === value;

  function onChange(e: ChangeEvent<HTMLInputElement>): void {
    e.stopPropagation();
    state.setSelectedValue(value);
  }

  return {
    inputProps: {
      checked: isSelected,
      disabled: isDisabled,
      name,
      onChange,
      type: 'radio',
      value,
    },
    isDisabled,
    isSelected,
  };
}
