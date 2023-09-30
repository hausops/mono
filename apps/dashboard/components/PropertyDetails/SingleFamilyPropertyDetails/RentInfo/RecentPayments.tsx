import type {RentPayment} from '@/services/lease';
import {Badge} from '@/volto/Badge';
import {useMemo} from 'react';
import * as s from './RecentPayments.css';

type RecentPaymentsProps = {
  payments: RentPayment[];
};

export function RecentPayments({payments}: RecentPaymentsProps) {
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
    <table className={s.Table}>
      <thead className={s.TableHeader}>
        <HeaderRow />
      </thead>
      <tbody>
        {payments.map((p) => {
          const date = new Date(p.dueDate).toLocaleDateString();
          return (
            <Row
              key={date}
              date={date}
              amount={currencyFormatter.format(p.amount)}
              status={p.status}
            />
          );
        })}
      </tbody>
    </table>
  );
}

function HeaderRow() {
  return (
    <tr>
      <td scope="col" className={s.TableCell} width="50%">
        <span className={s.HeaderLabel}>Date</span>
      </td>
      <td scope="col" className={s.TableCell}>
        <span className={s.HeaderLabel}>Amount</span>
      </td>
      <td scope="col" className={s.BadgeCell}></td>
    </tr>
  );
}

function Row({
  date,
  amount,
  status,
}: {
  date: string;
  amount: string;
  status: RentPayment['status'];
}) {
  return (
    <tr>
      <td scope="col" className={s.TableCell}>
        {date}
      </td>
      <td scope="col" className={s.TableCell}>
        {amount}
      </td>
      <td scope="col" className={s.BadgeCell}>
        <StatusBadge status={status} />
      </td>
    </tr>
  );
}

function StatusBadge({status}: {status: RentPayment['status']}) {
  switch (status) {
    case 'fully-paid':
      return <Badge>Paid</Badge>;
    case 'overdue':
      return <Badge status="attention">Overdue</Badge>;
    default:
      return null;
  }
}
