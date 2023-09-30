import {PropertyImages} from '@/components/PropertyDetails/PropertyImages';
import type {MultiFamily} from '@/services/property';
import {Section} from '@/volto/Section';
import * as s from './MultiFamilyPropertyDetails.css';
import {PropertyInfo} from './PropertyInfo';
import {UnitList} from './UnitList';

type MultiFamilyPropertyDetailsProps = {
  property: MultiFamily.Property;
};

export function MultiFamilyPropertyDetails(
  props: MultiFamilyPropertyDetailsProps,
) {
  const {property} = props;
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

      <Section title="Units">
        <UnitList property={property} />
      </Section>
    </>
  );
}
