'use client';

import {LeaseServiceProvider, LocalLeaseService} from '@/services/lease';
import type {PropertyModel} from '@/services/property';
import {TooltipsManagerProvider} from '@/volto/Tooltip';
import {MultiFamilyPropertyDetails} from './MultiFamilyPropertyDetails';
import {SingleFamilyPropertyDetails} from './SingleFamilyPropertyDetails';

type PropertyDetailsProps = {
  property: PropertyModel;
};

const leaseSvc = new LocalLeaseService();

export function PropertyDetails({property}: PropertyDetailsProps) {
  return (
    <LeaseServiceProvider service={leaseSvc}>
      <TooltipsManagerProvider>
        {property.type === 'single-family' ? (
          <SingleFamilyPropertyDetails property={property} />
        ) : (
          <MultiFamilyPropertyDetails property={property} />
        )}
      </TooltipsManagerProvider>
    </LeaseServiceProvider>
  );
}
