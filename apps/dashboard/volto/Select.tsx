import clsx from 'clsx';
import {SelectHTMLAttributes, useId} from 'react';
import {ExpandMore} from './icons/ExpandMore';
import * as s from './Select.css';

type Option = {
  label: string;
  value: string | number;
};

type OwnSelectProps = {
  className?: string;
  // disabled?: boolean;
  label: string;
  name: string;
  // required?: boolean;
  options: Option[];
};

type SelectProps = OwnSelectProps &
  Omit<SelectHTMLAttributes<HTMLSelectElement>, keyof OwnSelectProps>;

export default function Select({
  className,
  label,
  name,
  options,
  ...passthroughProps
}: SelectProps) {
  const fieldId = useId();
  return (
    <div className={clsx(s.Select, className)}>
      <label className={s.Label} htmlFor={fieldId}>
        {label}
      </label>
      <div className={s.InputWrapper}>
        <select
          {...passthroughProps}
          className={s.Input}
          id={fieldId}
          name={name}
        >
          <option></option>
          {options.map((o) => (
            <option key={o.value} value={o.value}>
              {o.label}
            </option>
          ))}
        </select>
        <div className={s.ExpandIcon}>
          <ExpandMore />
        </div>
      </div>
    </div>
  );
}
