import {Address} from '@/services/address';
import type {PropertyModel} from '@/services/property';
import {AspectRatio} from '@/volto/AspectRatio';
import {IconButton} from '@/volto/Button';
import {Card} from '@/volto/Card';
import {MoreHIcon} from '@/volto/icons';
import Image from 'next/image';
import Link from 'next/link';
import * as s from './PropertySummary.css';

type PropertySummaryProps = {
  property: PropertyModel;
};

export function PropertySummary({property}: PropertySummaryProps) {
  const {id, coverImageUrl, address} = property;

  const addr = Address.from(address);
  const [street, region] = addr.format();

  return (
    <Card as="article">
      <Cover image={coverImageUrl} caption={addr.toString()} />
      <div className={s.Body}>
        <div>
          <p className={s.Title}>
            <Link href={`/properties/${id}`}>{street}</Link>
          </p>
          <p>{region}</p>
        </div>
        <div>
          <IconButton icon={<MoreHIcon />} />
        </div>
      </div>
    </Card>
  );
}

function Cover({image, caption = ''}: {image?: string; caption?: string}) {
  return (
    <AspectRatio as="figure" ratio="2:1" className={s.Cover}>
      {image ? (
        <Image
          src={image}
          alt={caption}
          className={s.CoverImage}
          fill
          sizes={coverSizes}
        />
      ) : (
        <NoImage />
      )}
    </AspectRatio>
  );
}

const coverSizes = [
  '(min-width: 75rem) 40vw',
  '(min-width: 60rem) 34vw',
  '(min-width: 50rem) 62vw',
  '54vw',
].join(', ');

function NoImage() {
  return <div className={s.NoImage}>No image</div>;
}
