import {Address} from '@/services/address';
import {LeaseModel, useLeaseService} from '@/services/lease';
import {SingleFamilyProperty} from '@/services/property';
import {Section, SectionSkeleton} from '@/volto/Section';
import Link from 'next/link';
import useSWR from 'swr';
import {RecentPayments} from './RecentPayments';
import * as s from './SingleFamilyPropertyDetails.css';
import {TenantProfile} from './TenantProfile';
import {UpcomingRent} from './UpcomingRent';

type Props = {
  property: SingleFamilyProperty;
};

export function SingleFamilyPropertyDetails({property}: Props) {
  const {coverImageUrl, address, unit} = property;
  const addr = Address.from(address);

  const leaseSvc = useLeaseService();
  const {data: lease} = useSWR(`lease.unit-${unit.id}`, () =>
    leaseSvc.getByUnitId(unit.id)
  );

  return (
    <section className={s.SingleFamilyPropertyDetails}>
      <article className={s.Column}>
        <Section title="Property info"></Section>
      </article>

      <aside className={s.Column}>
        {lease ? (
          <RentInfo lease={lease} />
        ) : (
          <>
            <SectionSkeleton />
            <SectionSkeleton height="15rem" />
          </>
        )}
      </aside>
    </section>
  );
}

function RentInfo({lease}: {lease: LeaseModel}) {
  const {tenant, upcomingRent, pastPayments} = lease;
  return (
    <>
      <Section title="Tenant">
        <TenantProfile
          name={`${tenant.firstName} ${tenant.lastName}`}
          imageUrl={tenant.imageUrl}
          email={tenant.email}
          phone={tenant.phone}
        />
      </Section>

      {upcomingRent && (
        <Section title="Upcoming rent">
          <UpcomingRent {...upcomingRent} />
        </Section>
      )}

      <Section
        title="Recent payments"
        // TODO: build payments (per unit) page
        actions={
          <Link className={s.TextLink} href="#">
            View all
          </Link>
        }
      >
        <RecentPayments payments={take(3, pastPayments)} />
      </Section>
    </>
  );
}

function take<T>(n: number, items: T[]): T[] {
  return items.slice(0, n);
}
