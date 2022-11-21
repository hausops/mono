import AspectRatio from '@/volto/AspectRatio';
import {IconButton} from '@/volto/Button';
import Card from '@/volto/Card';
import {MoreH} from '@/volto/icons';
import Image from 'next/image';
import Link from 'next/link';
import * as s from './PropertySummary.css';

type PropertySummaryProps = {
  property: {
    image?: string;
    address: {
      street: string;
      city: string;
      state: string;
      zipcode: string;
    };
  };
};

export default function PropertySummary({
  property: {image, address},
}: PropertySummaryProps) {
  return (
    <Card as="article">
      <Cover image={image} caption={address.street} />
      <div className={s.Body}>
        <div>
          <p className={s.Title}>
            <Link href="#">{address.street}</Link>
          </p>
          <p>
            {address.city}, {address.state} {address.zipcode}
          </p>
        </div>
        <div>
          <IconButton icon={<MoreH />} />
        </div>
      </div>
    </Card>
  );
}

function Cover({image, caption}: {image?: string; caption: string}) {
  return (
    <AspectRatio as="figure" ratio="2:1" className={s.Cover}>
      {image ? (
        <Image src={image} alt={caption} className={s.CoverImage} fill />
      ) : (
        <NoImage />
      )}
    </AspectRatio>
  );
}

function NoImage() {
  return <div className={s.NoImage}>No image</div>;
}
