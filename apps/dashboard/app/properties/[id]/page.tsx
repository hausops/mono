import {PropertyDetails} from '@/components/PropertyDetails';
import {PageHeader} from '@/layouts/PageHeader';
import {Address} from '@/services/address';
import {propertySvc} from '@/services/property';
import type {Metadata} from 'next';
import {notFound} from 'next/navigation';

type Params = {
  id: string;
};

export async function generateMetadata({params}: {params: Params}) {
  const property = await propertySvc.getById(params.id);
  if (!property) {
    notFound();
  }

  const [streetAddr] = Address.from(property.address).format();
  return {
    title: `${streetAddr} - HausOps`,
  } satisfies Metadata;
}

export default async function PropertyDetailsPage({params}: {params: Params}) {
  const property = await propertySvc.getById(params.id);
  if (!property) {
    notFound();
  }

  const [streetAddr] = Address.from(property.address).format();
  return (
    <>
      <PageHeader title={streetAddr} />
      <PropertyDetails property={property} />
    </>
  );
}
