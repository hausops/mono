import {LeaseServiceProvider, LocalLeaseService} from '@/services/lease';
import {PropertyModel} from '@/services/property';
import {MultiFamilyPropertyDetails} from './MultiFamilyPropertyDetails';
import {SingleFamilyPropertyDetails} from './SingleFamilyPropertyDetails';

type PropertyDetailsProps = {
  property: PropertyModel;
};

const leaseSvc = new LocalLeaseService();

export function PropertyDetails({property}: PropertyDetailsProps) {
  if (property.type === 'single-family') {
    return (
      <LeaseServiceProvider service={leaseSvc}>
        <SingleFamilyPropertyDetails property={property} />
      </LeaseServiceProvider>
    );
  }
  return <MultiFamilyPropertyDetails property={property} />;
}
