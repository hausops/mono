import clsx from 'clsx';
import {useId, type InputHTMLAttributes} from 'react';
import * as s from './TextField.css';

type OwnTextFieldProps = {
  className?: string;
  // disabled?: boolean;
  label: string;
  // required?: boolean;
};

type TextFieldProps = OwnTextFieldProps &
  Omit<InputHTMLAttributes<HTMLInputElement>, keyof OwnTextFieldProps>;

export function TextField({className, label, ...inputProps}: TextFieldProps) {
  const fieldId = useId();
  return (
    <div className={clsx(s.TextField, className)}>
      <label className={s.Label} htmlFor={fieldId}>
        {label}
      </label>
      <input {...inputProps} className={s.Input} id={fieldId} />
    </div>
  );
}
