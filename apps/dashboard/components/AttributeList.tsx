import clsx from 'clsx';
import type {PropsWithChildren, ReactNode} from 'react';
import * as s from './AttributeList.css';

export function AttributeList({
  children,
  className,
}: PropsWithChildren<{className?: string}>) {
  return <dl className={clsx(s.AttributeList, className)}>{children}</dl>;
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
    <div className={s.Attribute}>
      <dt>{label}</dt>
      <dd>{value ?? fallbackValue}</dd>
    </div>
  );
}
