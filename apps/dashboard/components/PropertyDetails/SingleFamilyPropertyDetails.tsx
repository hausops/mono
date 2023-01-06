import {Address} from '@/services/address';
import {SingleFamilyProperty} from '@/services/property';
import {Section} from '@/volto/Section';
import {RentInfo} from './RentInfo';
import * as s from './SingleFamilyPropertyDetails.css';

type Props = {
  property: SingleFamilyProperty;
};

export function SingleFamilyPropertyDetails({property}: Props) {
  const {coverImageUrl, address, unit} = property;
  const addr = Address.from(address);

  return (
    <section className={s.SingleFamilyPropertyDetails}>
      <article className={s.Column}>
        <Section title="Property info"></Section>
      </article>

      <aside className={s.Column}>
        <RentInfo property={property} />
      </aside>
    </section>
  );
}
