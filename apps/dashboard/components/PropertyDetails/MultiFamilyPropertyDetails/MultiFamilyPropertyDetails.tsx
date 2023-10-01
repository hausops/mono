'use client';

import {PropertyImages} from '@/components/PropertyDetails/PropertyImages';
import type {MultiFamily} from '@/services/property';
import {Section} from '@/volto/Section';
import {TooltipsManagerProvider} from '@/volto/Tooltip';
import * as s from './MultiFamilyPropertyDetails.css';
import {PropertyInfo} from './PropertyInfo';
import {UnitList} from './UnitList';

export function MultiFamilyPropertyDetails({
  property,
}: {
  property: MultiFamily.Property;
}) {
  return (
    <TooltipsManagerProvider>
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
    </TooltipsManagerProvider>
  );
}
