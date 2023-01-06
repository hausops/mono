import {PropertyModel} from '@/services/property';
import {SingleFamilyPropertyDetails} from './SingleFamilyPropertyDetails';
import {LocalLeaseService, LeaseServiceProvider} from '@/services/lease';

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
  return <p>TODO: MultiFamily</p>;
}
