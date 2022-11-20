import clsx from 'clsx';
import {PropsWithChildren, ReactElement} from 'react';
import * as s from './Button.css';
import type {ButtonVarient} from './types';

type ButtonProps = PropsWithChildren<{
  varient: ButtonVarient;
  // color
}>;

export default function Button({children, varient}: ButtonProps) {
  const className = clsx(s.base, s.varient[varient]);
  return <button className={className}>{children}</button>;
}

type IconButtonProps = PropsWithChildren<{
  icon: ReactElement;
}>;

export function IconButton({icon}: IconButtonProps) {
  return <button className={s.IconButton}>{icon}</button>;
}
