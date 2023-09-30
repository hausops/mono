import type {LeaseModel} from '@/services/lease';
import type {MultiFamily} from '@/services/property';
import {Badge, LivenessBadge} from '@/volto/Badge';
import {IconButton} from '@/volto/Button';
import {TextSkeleton} from '@/volto/Skeleton';
import {MoreHIcon} from '@/volto/icons';
import Link from 'next/link';
import {useMemo} from 'react';
import * as s from './Unit.css';

type UnitProps = {
  lease?: LeaseModel;
  unit: MultiFamily.Unit;
};

export function Unit(props: UnitProps) {
  const {lease, unit} = props;
  return (
    <li className={s.Unit}>
      <div className={s.Column}>
        <h4 className={s.Title}>{unit.number}</h4>
        <ul className={s.Info}>
          <li className={s.InfoItem}>
            {unit.bedrooms === 0 ? 'Studio' : `${unit.bedrooms} bedrooms`}
          </li>
          <li className={s.InfoItem}>{unit.bathrooms} bathrooms</li>
          <li className={s.InfoItem}>
            {unit.size == null
              ? null
              : `${formatters.default.format(unit.size)} sq.ft.`}
          </li>
        </ul>
      </div>
      <div className={s.Column}>
        {lease ? (
          <Rented lease={lease} />
        ) : (
          <Vacant activeListingId={unit.activeListing?.id} />
        )}
      </div>
      <div>
        <IconButton icon={<MoreHIcon />} />
      </div>
    </li>
  );
}

type VacantProps = {
  activeListingId?: string;
};

function Vacant(props: VacantProps) {
  const {activeListingId} = props;
  return (
    <>
      <LivenessBadge>Vacant</LivenessBadge>
      {activeListingId && (
        <Link className={s.TextLink} href="#">
          See listing
        </Link>
      )}
    </>
  );
}

type RentedProps = {
  lease: LeaseModel;
};

function Rented(props: RentedProps) {
  const {lease} = props;
  const {pastPayments, upcomingRent} = lease;
  const overdue = useMemo(
    () => pastPayments.find((p) => p.status === 'overdue'),
    [pastPayments],
  );
  return (
    <>
      <LivenessBadge status="live">Rented</LivenessBadge>
      {upcomingRent && (
        <p>
          Upcoming rent {formatters.currency.format(upcomingRent.amount)} on{' '}
          {new Date(upcomingRent.dueDate).toLocaleDateString()}.
        </p>
      )}
      {overdue && (
        <p>
          Outstanding rent{' '}
          {formatters.currency.format(overdue.amount - overdue.paid)} due{' '}
          {new Date(overdue.dueDate).toLocaleDateString()}.{' '}
          <Badge status="attention">Overdue</Badge>
        </p>
      )}
    </>
  );
}

const formatters = {
  currency: Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD',
    maximumSignificantDigits: 2,
  }),

  default: Intl.NumberFormat('en-US'),
};

export function UnitSkeleton() {
  return (
    <div className={s.Unit}>
      <div className={s.Column}>
        <TextSkeleton width="6rem" />
        <TextSkeleton width="15rem" />
      </div>
      <div className={s.Column}>
        <LivenessBadge>
          <TextSkeleton width="4rem" />
        </LivenessBadge>
        <TextSkeleton width="15rem" />
      </div>
      <div>
        <IconButton icon={<MoreHIcon />} />
      </div>
    </div>
  );
}
