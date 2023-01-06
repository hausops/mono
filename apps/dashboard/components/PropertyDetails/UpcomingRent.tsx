import {RentPayment} from '@/services/lease';
import {Badge} from '@/volto/Badge';
import {useMemo} from 'react';
import {Entry, EntryList} from './EntryList';
import * as s from './UpcomingRent.css';

type UpcomingRentProps = RentPayment;

export function UpcomingRent({
  dueDate,
  status,
  amount,
  paid,
  paymentPending,
}: UpcomingRentProps) {
  const currencyFormatter = useMemo(
    () =>
      Intl.NumberFormat('en-US', {
        style: 'currency',
        currency: 'USD',
        maximumSignificantDigits: 2,
      }),
    []
  );

  return (
    <EntryList>
      <Entry
        label="Due date"
        value={
          <p className={s.DueDate}>
            <span>{new Date(dueDate).toLocaleDateString()}</span>
            {status === 'overdue' && <Badge status="attention">Overdue</Badge>}
          </p>
        }
      />
      <Entry label="Rent amount" value={currencyFormatter.format(amount)} />
      {paid ? (
        <Entry label="Paid" value={currencyFormatter.format(paid)} />
      ) : null}
      {paymentPending ? (
        <Entry
          label="Payment pending"
          value={currencyFormatter.format(paymentPending)}
        />
      ) : null}
    </EntryList>
  );
}
