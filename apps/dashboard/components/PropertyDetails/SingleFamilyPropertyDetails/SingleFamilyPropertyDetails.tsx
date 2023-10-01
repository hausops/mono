'use client';

import type {SingleFamily} from '@/services/property';
import {TooltipsManagerProvider} from '@/volto/Tooltip';
import {PropertyImages} from '../PropertyImages';
import {PropertyInfo} from './PropertyInfo';
import {RentInfo} from './RentInfo';
import * as s from './SingleFamilyPropertyDetails.css';

export function SingleFamilyPropertyDetails({
  property,
}: {
  property: SingleFamily.Property;
}) {
  return (
    <TooltipsManagerProvider>
      <section className={s.SingleFamilyPropertyDetails}>
        <article className={s.Column}>
          <PropertyInfo property={property} />
          <PropertyImages property={property} />
        </article>

        <aside className={s.Column}>
          <RentInfo property={property} />
        </aside>
      </section>
    </TooltipsManagerProvider>
  );
}
