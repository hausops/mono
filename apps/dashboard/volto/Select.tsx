import {ExpandMoreIcon} from '@/volto/icons';
import clsx from 'clsx';
import {SelectHTMLAttributes, useId} from 'react';
import * as s from './Select.css';

type Option = {
  label: string;
  value: string | number;
};

type OwnSelectProps = {
  className?: string;
  // disabled?: boolean;
  label: string;
  // required?: boolean;
  options: Option[];
};

type SelectProps = OwnSelectProps &
  Omit<SelectHTMLAttributes<HTMLSelectElement>, keyof OwnSelectProps>;

export function Select({
  className,
  label,
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
        <select {...passthroughProps} className={s.Input} id={fieldId}>
          <option></option>
          {options.map((o) => (
            <option key={o.value} value={o.value}>
              {o.label}
            </option>
          ))}
        </select>
        <div className={s.ExpandIcon}>
          <ExpandMoreIcon />
        </div>
      </div>
    </div>
  );
}

export function toOption<T extends number | string>(
  value: T
): {label: string; value: T} {
  return {label: `${value}`, value};
}
