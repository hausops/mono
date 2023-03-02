import {LeaseModel, useLeaseService} from '@/services/lease';
import {SingleFamilyProperty} from '@/services/property';
import {Button} from '@/volto/Button';
import {EmptyState} from '@/volto/EmptyState';
import {HomeIcon} from '@/volto/icons';
import {Section, SectionSkeleton} from '@/volto/Section';
import Link from 'next/link';
import useSWR from 'swr';
import {RecentPayments} from './RecentPayments';
import * as s from './RentInfo.css';
import {TenantProfile} from './TenantProfile';
import {UpcomingRent} from './UpcomingRent';

type RentInfoProps = {
  property: SingleFamilyProperty;
};

export function RentInfo({property}: RentInfoProps) {
  const leaseSvc = useLeaseService();
  const unitId = property.unit.id;
  const {isLoading, data} = useSWR(`lease.unit-${unitId}`, () =>
    leaseSvc.getByUnitId(unitId)
  );

  if (isLoading) {
    return (
      <>
        <SectionSkeleton />
        <SectionSkeleton height="15rem" />
      </>
    );
  }

  return data ? <Rented lease={data} /> : <Vacant property={property} />;
}

function Rented({lease}: {lease: LeaseModel}) {
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

function Vacant({property}: {property: SingleFamilyProperty}) {
  const {description, actions} = property.unit.activeListing
    ? {
        description: 'The property has an active listing.',
        actions: (
          <>
            <Button as={Link} variant="outlined" href="#">
              See listing
            </Button>
            <Button as={Link} variant="contained" href="#">
              Mark as rented
            </Button>
          </>
        ),
      }
    : {
        description: 'Some options that you can take…',
        actions: (
          <>
            <Button as={Link} variant="outlined" href="#">
              List property
            </Button>
            <Button as={Link} variant="contained" href="#">
              Mark as rented
            </Button>
          </>
        ),
      };

  return (
    <Section title="Rent info">
      <EmptyState
        icon={<HomeIcon />}
        title="Not currently rented"
        description={description}
        actions={actions}
      />
    </Section>
  );
}
