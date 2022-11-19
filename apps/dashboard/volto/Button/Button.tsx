import {PropsWithChildren} from 'react';
import type {ButtonVarient} from './types';
import * as s from './Button.css';
import clsx from 'clsx';

type ButtonProps = PropsWithChildren<{
  varient: ButtonVarient;
  // color
}>;

export default function Button({children, varient}: ButtonProps) {
  const className = clsx(s.base, s.varient[varient]);
  return <button className={className}>{children}</button>;
}

export function IconButton() {}
