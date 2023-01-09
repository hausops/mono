import {Select, toOption} from '@/volto/Select';
import {ComponentPropsWithoutRef, useMemo} from 'react';

type Props = {
  name: string;
  onChange: (selection: number | undefined) => void;
} & Pick<ComponentPropsWithoutRef<typeof Select>, 'value'>;

export function BedroomsSelect({name, value, onChange}: Props) {
  const options = useMemo(
    () => [{label: 'Studio', value: 0}, ...[1, 2, 3, 4, 5].map(toOption)],
    []
  );

  return (
    <Select
      label="Beds"
      name={name}
      options={options}
      value={value}
      onChange={(e) => {
        const v = e.target.value;
        onChange(v ? +v : undefined);
      }}
    />
  );
}

export function BathroomsSelect({name, value, onChange}: Props) {
  const options = useMemo(
    () => [
      {label: 'None', value: 0},
      ...[1, 1.5, 2, 2.5, 3, 3.5, 4].map(toOption),
    ],
    []
  );
  return (
    <Select
      label="Baths"
      name={name}
      options={options}
      value={value}
      onChange={(e) => {
        const v = e.target.value;
        onChange(v ? +v : undefined);
      }}
    />
  );
}
