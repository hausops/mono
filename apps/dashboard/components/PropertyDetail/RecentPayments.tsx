import {Badge} from '@/volto/Badge';
import {ComponentType, PropsWithChildren, ReactElement, ReactNode} from 'react';
import * as s from './RecentPayments.css';

type RecentPaymentsProps = {
  // payments: {
  //   date: string;
  //   amount: number;
  //   status?: 'overdue' | 'paid';
  // }[];
};

export function RecentPayments(props: RecentPaymentsProps) {
  return (
    <table className={s.Table}>
      <thead className={s.TableHeader}>
        <HeaderRow />
      </thead>
      <tbody>
        <Row
          date="12/1/2022"
          amount="$4,000"
          badge={<Badge status="attention">Overdue</Badge>}
        />
        <Row date="11/1/2022" amount="$4,000" badge={<Badge>Paid</Badge>} />
        <Row date="10/1/2022" amount="$4,000" badge={<Badge>Paid</Badge>} />
      </tbody>
    </table>
  );
}

function HeaderRow() {
  return (
    <tr>
      <td scope="col" className={s.TableCell}>
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
  badge,
}: {
  date: string;
  amount: string;
  badge?: ReactElement<typeof Badge>;
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
        {badge}
      </td>
    </tr>
  );
}
