import {propertySvc} from '@/services/property';
import 'server-only';

import {PropertySummary} from './_internal/PropertySummary';

export default async function PropertiesPage() {
  const properties = await propertySvc.getAll();
  return properties.map((p) => (
    <li key={p.id}>
      <PropertySummary property={p} />
    </li>
  ));
}
