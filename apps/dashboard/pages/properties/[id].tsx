import {PropertyDetail} from '@/components/PropertyDetail';
import {PageLayout} from '@/layouts/Page';
import {PageHeader} from '@/layouts/PageHeader';
import {Address} from '@/services/address';
import {LocalPropertyService, PropertyModel} from '@/services/property';
import {Button} from '@/volto/Button';
import {EmptyState} from '@/volto/EmptyState';

import {Home as HomeIcon} from '@/volto/icons';
import {GetServerSideProps} from 'next';
import Head from 'next/head';
import Link from 'next/link';
import * as s from './id.css';

type PageProps = {notFound: true} | {notFound?: false; property: PropertyModel};

export default function Page(props: PageProps) {
  if (props.notFound) {
    return <NotFound />;
  }

  const {property} = props;
  const [streetAddr] = Address.from(property.address).format();
  return (
    <>
      <Head>
        <title>{streetAddr} - HausOps</title>
        <meta name="description" content="HausOps" />
        <link rel="icon" href="data:;base64,iVBORw0KGgo=" />
      </Head>

      <PageLayout>
        <PageHeader title={streetAddr} />
        <pre>{JSON.stringify(property, null, 2)}</pre>
      </PageLayout>
    </>
  );
}

export const getServerSideProps: GetServerSideProps<
  PageProps,
  {id: string}
> = async ({params}) => {
  const propertySvc = new LocalPropertyService();

  if (params?.id) {
    const property = await propertySvc.get(params.id);
    return {
      props: property ? {property} : {notFound: true},
    };
  }

  return {props: {notFound: true}};
};

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
