import PropertySummary from '@/components/PropertySummary';
import Button from '@/volto/Button';
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

      <section className={s.Page}>
        {/* breadcrumb */}

        <header className={s.Header}>
          <h1 className={s.Title}>Properties</h1>
          <Button as={Link} variant="contained" href="/properties/new">
            Add property
          </Button>
        </header>

        <ul className={s.PropertyList}>
          {properties.map((p) => (
            <li key={p.id}>
              <PropertySummary property={p} />
            </li>
          ))}
        </ul>
      </section>
    </>
  );
}
