'use client';

import {PageHeader} from '@/layouts/PageHeader';
import {Address} from '@/services/address';
import type {SingleFamily} from '@/services/property';
import {TooltipsManagerProvider} from '@/volto/Tooltip';
import {useEffect} from 'react';
import useSWR from 'swr';

import {PropertyImages} from '../PropertyImages';
import {PropertyInfo} from './PropertyInfo';
import {RentInfo} from './RentInfo';
import * as s from './SingleFamilyPropertyDetails.css';

type Props = {
  property: SingleFamily.Property;
};

export function SingleFamilyPropertyDetails(props: Props) {
  const {data: property, mutate: mutateProperty} = useSWR(
    `/api/properties/${props.property.id}`,
    async (endpoint): Promise<SingleFamily.Property> => {
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

      <section className={s.SingleFamilyPropertyDetails}>
        <article className={s.Column}>
          <PropertyInfo
            property={property}
            onUpdateSuccess={(updatedProperty) => {
              mutateProperty(updatedProperty, {revalidate: false});
            }}
          />
          <PropertyImages property={property} />
        </article>

        <aside className={s.Column}>
          <RentInfo property={property} />
        </aside>
      </section>
    </TooltipsManagerProvider>
  );
}
