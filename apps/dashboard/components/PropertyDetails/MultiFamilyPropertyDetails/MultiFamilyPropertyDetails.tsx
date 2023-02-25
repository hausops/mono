import {PropertyImages} from '@/components/PropertyDetails/PropertyImages';
import {MultiFamilyProperty} from '@/services/property';
import * as s from './MultiFamilyPropertyDetails.css';
import {PropertyInfo} from './PropertyInfo';

type Props = {
  property: MultiFamilyProperty;
};

export function MultiFamilyPropertyDetails({property}: Props) {
  return (
    <>
      <section className={s.TwoColumnsLayout}>
        <article className={s.Column}>
          <PropertyInfo property={property} />
        </article>
        <aside className={s.Column}>
          <PropertyImages property={property} />
        </aside>
      </section>
    </>
  );
}
