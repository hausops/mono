import {PropertyModel} from '@/services/property';
import {AspectRatio} from '@/volto/AspectRatio';
import {IconButton} from '@/volto/Button';
import {Card} from '@/volto/Card';
import {MoreH} from '@/volto/icons';
import {TextSkeleton} from '@/volto/Skeleton';
import Image from 'next/image';
import Link from 'next/link';
import * as s from './PropertySummary.css';

type PropertySummaryProps = {
  property: PropertyModel;
};

export function PropertySummary({
  property: {coverImageUrl, name, address},
}: PropertySummaryProps) {
  return (
    <Card as="article">
      <Cover image={coverImageUrl} caption={name} />
      <div className={s.Body}>
        <div>
          <p className={s.Title}>
            <Link href="#">{address.line1}</Link>
          </p>
          <p>
            {address.city}, {address.state} {address.zip}
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

export function PropertySummarySkeleton() {
  return (
    <Card as="article">
      <AspectRatio as="figure" ratio="2:1" className={s.Cover} />
      <div className={s.SkeletonBody}>
        <TextSkeleton width="50%" />
        <TextSkeleton width="70%" />
      </div>
    </Card>
  );
}
