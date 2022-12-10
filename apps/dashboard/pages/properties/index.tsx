import PropertySummary from '@/components/PropertySummary';
import Button from '@/volto/Button';
import PageLayout from '@/layouts/Page';
import PageHeader from '@/layouts/PageHeader';
import Head from 'next/head';
import Link from 'next/link';
import data from './data.json';
import * as s from './index.css';

export default function Page() {
  const properties = Object.values(data);
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

        <ul className={s.PropertyList}>
          {properties.map((p) => (
            <li key={p.id}>
              <PropertySummary property={p} />
            </li>
          ))}
        </ul>
      </PageLayout>
    </>
  );
}
