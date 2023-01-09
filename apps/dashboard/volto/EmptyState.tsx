import {ReactElement, ReactNode} from 'react';
import * as s from './EmptyState.css';

type EmptyStateProps = {
  icon?: ReactElement;
  title: string;
  description?: ReactNode;
  actions?: ReactNode;
};

export function EmptyState({
  icon,
  title,
  description,
  actions,
}: EmptyStateProps) {
  return (
    <div className={s.EmptyState}>
      {icon && <div className={s.Icon}>{icon}</div>}
      <div className={s.Body}>
        <h1 className={s.Title}>{title}</h1>
        {description && <p className={s.Description}>{description}</p>}
      </div>
      {actions && <div className={s.Actions}>{actions}</div>}
    </div>
  );
}
