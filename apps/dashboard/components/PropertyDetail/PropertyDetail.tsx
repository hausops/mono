import {Address} from '@/services/address';
import {PropertyModel, SingleFamilyProperty} from '@/services/property';
import {Section} from '@/volto/Section';
import Link from 'next/link';
import * as s from './PropertyDetail.css';
import {RecentPayments} from './RecentPayments';
import {TenantProfile} from './TenantProfile';
import {UpcomingRent} from './UpcomingRent';

type PropertyDetailProps = {
  property: PropertyModel;
};

export function PropertyDetail({property}: PropertyDetailProps) {
  if (property.type === 'single-family') {
    return <SingleFamily property={property} />;
  }
  return <p>TODO: MultiFamily</p>;
}

function SingleFamily({property}: {property: SingleFamilyProperty}) {
  const {coverImageUrl, address, bedrooms, bathrooms, size} = property;
  const addr = Address.from(address);
  return (
    <section className={s.SingleFamily}>
      <article className={s.Column}>
        <Section title="Tenant">
          <TenantProfile
            name="Jane Doe"
            imageUrl="/images/michael-dam-mEZ3PoFGs_k-unsplash-avatar.jpg"
            email="jane.doe@example.com"
            phone="(555) 123-4567"
          />
        </Section>
        <Section title="Upcoming rent">
          <UpcomingRent />
        </Section>
        <Section
          title="Recent payments"
          actions={
            <Link className={s.TextLink} href="#">
              View all
            </Link>
          }
        >
          <RecentPayments />
        </Section>
      </article>

      <aside className={s.Column}>
        <Section title="Property info"></Section>
      </aside>
    </section>
  );
}
