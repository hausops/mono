import {ComponentPropsWithoutRef} from 'react';
import {TextField} from './TextField';

type NumericFieldProps = {
  value?: number;
  onChange: (newValue: number | undefined) => void;
  // min?: number;
} & ComponentPropsWithoutRef<typeof TextField>;

// state -> formatted number (string)
// onValueChange(value: number | undefined): void;
export function NumericField({
  label,
  name,
  value,
  onChange,
}: NumericFieldProps) {
  return (
    <TextField
      type="number"
      label={label}
      name={name}
      value={toString(value)}
      onChange={(event) => {
        onChange(toNumber(event.target.value));
      }}
    />
  );
}

function toString(num: number | undefined): string {
  return typeof num !== 'undefined' ? `${num}` : '';
}

function toNumber(str: string): number | undefined {
  if (str === '' || Number.isNaN(+str)) {
    return undefined;
  }
  return +str;
}
