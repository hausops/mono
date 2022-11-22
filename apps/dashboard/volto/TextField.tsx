import clsx from 'clsx';
import {HTMLInputTypeAttribute} from 'react';
import * as s from './TextField.css';

type TextFieldProps = {
  className?: string;
  // disabled?: boolean;
  id?: string;
  label: string;
  name: string;
  optional?: boolean;
  placeholder?: string;
  type?: HTMLInputTypeAttribute;
};

export default function TextField({
  className,
  label,
  optional,
  name,
  id = name,
  ...inputProps
}: TextFieldProps) {
  return (
    <div className={clsx(s.TextField, className)}>
      <label className={s.Label} htmlFor={id}>
        {label}
      </label>
      <input {...inputProps} className={s.Input} id={id} name={name} />
    </div>
  );
}
