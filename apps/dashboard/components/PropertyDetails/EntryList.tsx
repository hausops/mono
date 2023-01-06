import {ReactNode} from 'react';
import * as s from './EntryList.css';

export function EntryList({children}: {children: ReactNode}) {
  return <div className={s.EntryList}>{children}</div>;
}

export function Entry({label, value}: {label: string; value: ReactNode}) {
  return (
    <dl className={s.Entry}>
      <dt>{label}</dt>
      <dd>{value}</dd>
    </dl>
  );
}
