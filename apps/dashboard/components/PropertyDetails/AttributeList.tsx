import clsx from 'clsx';
import {PropsWithChildren, ReactNode} from 'react';
import * as s from './AttributeList.css';

export function AttributeList({
  children,
  className,
}: PropsWithChildren<{className?: string}>) {
  return <div className={clsx(s.AttributeList, className)}>{children}</div>;
}

export function Attribute({
  label,
  value,
  fallbackValue = '-',
}: {
  label: string;
  value: ReactNode;
  fallbackValue?: ReactNode;
}) {
  return (
    <dl className={s.Attribute}>
      <dt>{label}</dt>
      <dd>{value ?? fallbackValue}</dd>
    </dl>
  );
}
