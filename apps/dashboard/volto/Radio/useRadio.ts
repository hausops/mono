import {ChangeEvent, RefObject} from 'react';
import {RadioGroupState} from './useRadioGroupState';

type RadioProps<T> = {
  name: string;
  value: T;
  isDisabled?: boolean;
};

type Radio<T> = {
  // pass clickableZoneProps to the element that contains the entire Radio option.
  clickableZoneProps: {
    onClick: () => void;
  };
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
  state: RadioGroupState<T>,
  ref: RefObject<HTMLInputElement>,
): Radio<T> {
  const {name, value, isDisabled = state.isDisabled} = props;
  const isSelected = state.selectedValue === value;

  return {
    clickableZoneProps: {
      onClick() {
        state.setSelectedValue(value);
        ref.current?.focus();
      },
    },
    inputProps: {
      checked: isSelected,
      disabled: isDisabled,
      name,
      type: 'radio',
      value,
      onChange(e) {
        e.stopPropagation();
        state.setSelectedValue(value);
      },
    },
    isDisabled,
    isSelected,
  };
}
