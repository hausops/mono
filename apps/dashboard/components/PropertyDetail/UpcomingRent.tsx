import {Badge} from '@/volto/Badge';
import {ReactNode} from 'react';
import * as s from './UpcomingRent.css';

type ISO8601DateString = string;
type UpcomingRent = {
  dueDate: ISO8601DateString;
  rentAmount: number;
  paid?: number;
  paymentPending?: number;
};

const UPCOMING_RENT_DEMO: UpcomingRent = {
  dueDate: '2023-03-01T00:00:00',
  rentAmount: 4000,
  paid: 3000,
  paymentPending: 1000,
};

type UpcomingRentProps = {};

export function UpcomingRent(props: UpcomingRentProps) {
  const {dueDate, rentAmount, paid, paymentPending} = UPCOMING_RENT_DEMO;
  const currencyFormatter = Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD',
    maximumSignificantDigits: 2,
  });

  return (
    <div className={s.UpcomingRent}>
      <Entry label="Due date" value={<DueDate dueDate={dueDate} />} />
      <Entry label="Rent amount" value={currencyFormatter.format(rentAmount)} />
      {paid && <Entry label="Paid" value={currencyFormatter.format(paid)} />}
      {paymentPending && (
        <Entry
          label="Payment pending"
          value={currencyFormatter.format(paymentPending)}
        />
      )}
    </div>
  );
}

function Entry({label, value}: {label: string; value: ReactNode}) {
  return (
    <dl className={s.Entry}>
      <dt className={s.EntryLabel}>{label}</dt>
      <dd className={s.EntryValue}>{value}</dd>
    </dl>
  );
}

function DueDate({dueDate}: {dueDate: string}) {
  const d = new Date(dueDate);
  // TODO: move isOverdue logic to the backend
  const isOverdue = d.getTime() - Date.now() < 0;
  return (
    <p className={s.DueDate}>
      <span>{d.toLocaleDateString()}</span>
      {isOverdue && <Badge status="attention">Overdue</Badge>}
    </p>
  );
}
