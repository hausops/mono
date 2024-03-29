import {leaseSvc} from '@/services/lease';
import type {MultiFamily} from '@/services/property';
import useSWR from 'swr';
import {Unit, UnitSkeleton} from './Unit';

type UnitListProps = {
  property: MultiFamily.Property;
};

export function UnitList(props: UnitListProps) {
  const {property} = props;

  const {isLoading, data: leaseByUnitId} = useSWR(
    `/api/leases?property_id=${property.id}`,
    async () => {
      const unitIds = property.units.map((u) => u.id).sort();
      const leases = await leaseSvc.getManyByUnitIds(unitIds);
      return new Map(leases.map((lease) => [lease.unitId, lease]));
    },
  );

  if (isLoading) {
    return (
      <div>
        <UnitSkeleton />
        <UnitSkeleton />
        <UnitSkeleton />
      </div>
    );
  }

  return (
    <ul>
      {property.units.map((u) => (
        <Unit key={u.id} unit={u} lease={leaseByUnitId?.get(u.id)} />
      ))}
    </ul>
  );
}
