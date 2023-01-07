import clsx from 'clsx';
import {PropsWithChildren, ReactNode} from 'react';
import * as s from './AttributeList.css';

export function AttributeList({
  children,
  className,
}: PropsWithChildren<{className?: string}>) {
  return <div className={clsx(s.AttributeList, className)}>{children}</div>;
}

export function Attribute({label, value}: {label: string; value: ReactNode}) {
  return (
    <dl className={s.Attribute}>
      <dt>{label}</dt>
      <dd>{value}</dd>
    </dl>
  );
}
