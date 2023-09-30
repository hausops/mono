import {Attribute, AttributeList} from '@/components/AttributeList';
import type {RentPayment} from '@/services/lease';
import {Badge} from '@/volto/Badge';
import {useMemo} from 'react';
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
    [],
  );

  return (
    <AttributeList>
      <Attribute
        label="Due date"
        value={
          <p className={s.DueDate}>
            <span>{new Date(dueDate).toLocaleDateString()}</span>
            {status === 'overdue' && <Badge status="attention">Overdue</Badge>}
          </p>
        }
      />
      <Attribute label="Rent amount" value={currencyFormatter.format(amount)} />
      {paid ? (
        <Attribute label="Paid" value={currencyFormatter.format(paid)} />
      ) : null}
      {paymentPending ? (
        <Attribute
          label="Payment pending"
          value={currencyFormatter.format(paymentPending)}
        />
      ) : null}
    </AttributeList>
  );
}
