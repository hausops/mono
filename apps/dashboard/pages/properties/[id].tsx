import {PropertyDetails} from '@/components/PropertyDetails';
import {PageLayout} from '@/layouts/Page';
import {PageHeader} from '@/layouts/PageHeader';
import {Address} from '@/services/address';
import {usePropertyService} from '@/services/property';
import {Button} from '@/volto/Button';
import {EmptyState} from '@/volto/EmptyState';
import {HomeIcon} from '@/volto/icons';
import {Skeleton} from '@/volto/Skeleton';
import Head from 'next/head';
import Link from 'next/link';
import {useRouter} from 'next/router';
import useSWR from 'swr';
import * as s from './id.css';

export default function Page() {
  const router = useRouter();
  const propertyId = router.query.id;

  const propertySvc = usePropertyService();
  const {isLoading, data: property} = useSWR(
    `/api/properties/${propertyId}`,
    () =>
      // this is undefined during SSR
      typeof propertyId === 'string' ? propertySvc.getById(propertyId) : null,
  );

  if (isLoading) {
    return <Loading />;
  }

  if (!property) {
    return <NotFound />;
  }

  const [streetAddr] = Address.from(property.address).format();
  // Cannot do in JSX due to next.js bug:
  // https://github.com/vercel/next.js/discussions/38256
  const pageTitle = `${streetAddr} - HausOps`;
  return (
    <>
      <Head>
        <title>{pageTitle}</title>
        <meta name="description" content="HausOps" />
        <link rel="icon" href="data:;base64,iVBORw0KGgo=" />
      </Head>

      <PageLayout>
        <PageHeader title={streetAddr} />
        <PropertyDetails property={property} />
      </PageLayout>
    </>
  );
}

function Loading() {
  return (
    <>
      <Head>
        <title>Property Details - HausOps</title>
        <meta name="description" content="HausOps" />
        <link rel="icon" href="data:;base64,iVBORw0KGgo=" />
      </Head>

      {/* TODO better loading state */}
      <PageLayout>
        <Skeleton width="10rem" />
      </PageLayout>
    </>
  );
}

function NotFound() {
  return (
    <>
      <Head>
        <title>Property not found - HausOps</title>
        <meta name="description" content="HausOps" />
        <link rel="icon" href="data:;base64,iVBORw0KGgo=" />
      </Head>

      <PageLayout>
        <div className={s.NotFound}>
          <EmptyState
            icon={<HomeIcon />}
            title="Property not found"
            description="Get started by adding a new property."
            actions={
              <>
                <Button as={Link} variant="outlined" href="/properties">
                  All properties
                </Button>
                <Button as={Link} variant="contained" href="/properties/new">
                  Add property
                </Button>
              </>
            }
          />
        </div>
      </PageLayout>
    </>
  );
}
