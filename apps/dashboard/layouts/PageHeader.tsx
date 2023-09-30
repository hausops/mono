import type {ReactNode} from 'react';
import * as s from './PageHeader.css';

type PageHeaderProps = {
  title: string;
  actions?: ReactNode;
};

export function PageHeader({title, actions}: PageHeaderProps) {
  return (
    <header className={s.Header}>
      <h1 className={s.Title}>{title}</h1>
      {actions && <div className={s.Actions}>{actions}</div>}
    </header>
  );
}
