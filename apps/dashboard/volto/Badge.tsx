import clsx from 'clsx';
import * as s from './Badge.css';

type BadgeProps = {
  children: string;
  status?: keyof typeof s.status;
};

export function Badge({children, status = 'default'}: BadgeProps) {
  return <span className={clsx(s.Badge, s.status[status])}>{children}</span>;
}
