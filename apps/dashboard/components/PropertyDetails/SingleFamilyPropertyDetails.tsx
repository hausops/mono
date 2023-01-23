import {SingleFamilyProperty} from '@/services/property';
import {PropertyImages} from './PropertyImages';
import {PropertyInfo} from './PropertyInfo';
import {RentInfo} from './RentInfo';
import * as s from './SingleFamilyPropertyDetails.css';

type Props = {
  property: SingleFamilyProperty;
};

export function SingleFamilyPropertyDetails({property}: Props) {
  return (
    <section className={s.SingleFamilyPropertyDetails}>
      <article className={s.Column}>
        <PropertyInfo property={property} />
        <PropertyImages property={property} />
      </article>

      <aside className={s.Column}>
        <RentInfo property={property} />
      </aside>
    </section>
  );
}
