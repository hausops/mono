import {
  PropertySummary,
  PropertySummarySkeleton,
} from '@/components/PropertySummary';
import {PageLayout} from '@/layouts/Page';
import {PageHeader} from '@/layouts/PageHeader';
import {usePropertyService} from '@/services/property';
import {Button} from '@/volto/Button';
import Head from 'next/head';
import Link from 'next/link';
import useSWR from 'swr';
import * as s from './index.css';

export default function Page() {
  return (
    <>
      <Head>
        <title>Properties - HausOps</title>
        <meta name="description" content="HausOps" />
        <link rel="icon" href="data:;base64,iVBORw0KGgo=" />
      </Head>

      <PageLayout>
        {/* breadcrumb */}

        <PageHeader
          title="Properties"
          actions={
            <Button as={Link} variant="contained" href="/properties/new">
              Add property
            </Button>
          }
        />
        <PropertyList />
      </PageLayout>
    </>
  );
}

function PropertyList() {
  const propertySvc = usePropertyService();
  const {data} = useSWR('/api/properties', () => propertySvc.getAll());

  const properties = data
    ? data.map((p) => (
        <li key={p.id}>
          <PropertySummary property={p} />
        </li>
      ))
    : Array.from({length: 8}).map((_, i) => (
        <li key={i}>
          <PropertySummarySkeleton />
        </li>
      ));

  return <ul className={s.PropertyList}>{properties}</ul>;
}
