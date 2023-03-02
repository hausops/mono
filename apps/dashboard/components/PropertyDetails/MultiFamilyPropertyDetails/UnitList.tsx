import {LeaseModel, useLeaseService} from '@/services/lease';
import {MultiFamilyProperty, RentalUnit} from '@/services/property';
import {Badge, LivenessBadge} from '@/volto/Badge';
import {IconButton} from '@/volto/Button';
import {TextSkeleton} from '@/volto/Skeleton';
import {MoreH} from '@/volto/icons';
import Link from 'next/link';
import {useMemo} from 'react';
import useSWR from 'swr';
import * as s from './UnitList.css';

type UnitListProps = {
  property: MultiFamilyProperty;
};

export function UnitList(props: UnitListProps) {
  const {property} = props;

  const leaseSvc = useLeaseService();
  const {isLoading, data: leaseByUnitId} = useSWR(
    `/api/leases?property_id=${property.id}`,
    async () => {
      const unitIds = property.units.map((u) => u.id).sort();
      const leases = await leaseSvc.getManyByUnitIds(unitIds);
      return new Map(leases.map((lease) => [lease.unitId, lease]));
    }
  );

  if (isLoading) {
    return (
      <div>
        <UnitSkeleton />
        <UnitSkeleton />
        <UnitSkeleton />
      </div>
    );
  }

  return (
    <ul>
      {property.units.map((u) => (
        <Unit key={u.id} unit={u} lease={leaseByUnitId?.get(u.id)} />
      ))}
    </ul>
  );
}

type UnitProps = {
  lease?: LeaseModel;
  unit: RentalUnit;
};

function Unit(props: UnitProps) {
  const {lease, unit} = props;
  return (
    <li className={s.Unit}>
      <div className={s.UnitColumn}>
        <h4 className={s.UnitTitle}>{unit.number}</h4>
        <ul className={s.UnitInfo}>
          <li className={s.UnitInfoItem}>
            {unit.bedrooms === 0 ? 'Studio' : `${unit.bedrooms} bedrooms`}
          </li>
          <li className={s.UnitInfoItem}>{unit.bathrooms} bathrooms</li>
          <li className={s.UnitInfoItem}>
            {unit.size == null
              ? null
              : `${formatters.default.format(unit.size)} sq.ft.`}
          </li>
        </ul>
      </div>
      <div className={s.UnitColumn}>
        {lease ? (
          <Rented lease={lease} />
        ) : (
          <Vacant activeListingId={unit.activeListing?.id} />
        )}
      </div>
      <div>
        <IconButton icon={<MoreH />} />
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
    [pastPayments]
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
      <div className={s.UnitColumn}>
        <TextSkeleton width="6rem" />
        <TextSkeleton width="15rem" />
      </div>
      <div className={s.UnitColumn}>
        <LivenessBadge>
          <TextSkeleton width="4rem" />
        </LivenessBadge>
        <TextSkeleton width="15rem" />
      </div>
      <div>
        <IconButton icon={<MoreH />} />
      </div>
    </div>
  );
}
