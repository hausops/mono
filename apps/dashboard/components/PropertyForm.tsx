import {Select, toOption} from '@/volto/Select';
import {ComponentPropsWithoutRef, useMemo} from 'react';

type Props = {
  name: string;
} & Pick<ComponentPropsWithoutRef<typeof Select>, 'value' | 'onChange'>;

export function NumBedroomsSelector({name, value, onChange}: Props) {
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
      onChange={onChange}
    />
  );
}

export function NumBathroomsSelector({name, value, onChange}: Props) {
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
      onChange={onChange}
    />
  );
}
