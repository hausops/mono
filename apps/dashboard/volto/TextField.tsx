import clsx from 'clsx';
import {HTMLInputTypeAttribute, InputHTMLAttributes} from 'react';
import * as s from './TextField.css';

type OwnTextFieldProps = {
  className?: string;
  // disabled?: boolean;
  id?: string;
  label: string;
  name: string;
  // required?: boolean;
};

type TextFieldProps = OwnTextFieldProps &
  Omit<InputHTMLAttributes<HTMLInputElement>, keyof OwnTextFieldProps>;

export default function TextField({
  className,
  label,
  // required,
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
