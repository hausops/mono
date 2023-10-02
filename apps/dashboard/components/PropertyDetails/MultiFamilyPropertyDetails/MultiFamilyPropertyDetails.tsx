'use client';

import {PageHeader} from '@/layouts/PageHeader';
import {Address} from '@/services/address';
import type {SingleFamily} from '@/services/property';
import {TooltipsManagerProvider} from '@/volto/Tooltip';
import {useEffect} from 'react';
import useSWR from 'swr';

import {PropertyImages} from '@/components/PropertyDetails/PropertyImages';
import type {MultiFamily} from '@/services/property';
import {Section} from '@/volto/Section';
import * as s from './MultiFamilyPropertyDetails.css';
import {PropertyInfo} from './PropertyInfo';
import {UnitList} from './UnitList';

type Props = {
  property: MultiFamily.Property;
};

export function MultiFamilyPropertyDetails(props: Props) {
  const {data: property, mutate: mutateProperty} = useSWR(
    `/api/properties/${props.property.id}`,
    async (endpoint): Promise<MultiFamily.Property> => {
      const res = await fetch(endpoint);
      return res.json();
    },
    {
      fallbackData: props.property,
    },
  );
  const [streetAddr] = Address.from(property.address).format();

  useEffect(() => {
    document.title = streetAddr;
  }, [streetAddr]);

  return (
    <TooltipsManagerProvider>
      <PageHeader title={streetAddr} />

      <section className={s.TwoColumnsLayout}>
        <article className={s.Column}>
          <PropertyInfo
            property={property}
            onUpdateSuccess={(updatedProperty) => {
              mutateProperty(updatedProperty, {revalidate: false});
            }}
          />
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
