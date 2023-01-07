import clsx from 'clsx';
import {
  ComponentPropsWithoutRef,
  ElementType,
  PropsWithChildren,
  ReactElement,
} from 'react';
import * as s from './Button.css';
import type {ButtonVariant} from './types';

type ButtonProps<As extends ElementType> = {
  // https://itnext.io/react-polymorphic-components-with-typescript-f7ce72ea7af2
  as?: As;
  variant: ButtonVariant;
  // color
};

export function Button<As extends ElementType = 'button'>({
  as,
  className,
  variant,
  ...props
}: ButtonProps<As> &
  Omit<ComponentPropsWithoutRef<As>, keyof ButtonProps<As>>) {
  const Root = as ?? 'button';
  return (
    <Root
      {...props}
      className={clsx(s.Button, s.ButtonVariants[variant], className)}
    />
  );
}

type IconButtonProps = PropsWithChildren<{
  icon: ReactElement;
}>;

export function IconButton({icon}: IconButtonProps) {
  return <button className={s.IconButton}>{icon}</button>;
}

type MiniTextButtonProps = PropsWithChildren<{
  icon?: ReactElement;
}> &
  ComponentPropsWithoutRef<'button'>;

export function MiniTextButton({
  children,
  icon,
  ...buttonProps
}: MiniTextButtonProps) {
  return (
    <button {...buttonProps} className={s.MiniTextButton}>
      {icon && <div className={s.MiniTextButtonIcon}>{icon}</div>}
      {children}
    </button>
  );
}
