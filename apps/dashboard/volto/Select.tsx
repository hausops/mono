import clsx from 'clsx';
import {SelectHTMLAttributes} from 'react';
import {ExpandMore} from './icons/ExpandMore';
import * as s from './Select.css';

type Option = {
  label: string;
  value: string | number;
};

type OwnSelectProps = {
  className?: string;
  // disabled?: boolean;
  id?: string;
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
  id = name,
  options,
  ...passthroughProps
}: SelectProps) {
  return (
    <div className={clsx(s.Select, className)}>
      <label className={s.Label} htmlFor={id}>
        {label}
      </label>
      <div className={s.InputWrapper}>
        <select {...passthroughProps} className={s.Input} id={id} name={name}>
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
