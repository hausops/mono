import {Address} from '@/services/address';
import {SingleFamilyProperty} from '@/services/property';
import {AspectRatio} from '@/volto/AspectRatio';
import {Section} from '@/volto/Section';
import Image from 'next/image';
import * as s from './PropertyImages.css';

export function PropertyImages({property}: {property: SingleFamilyProperty}) {
  const {coverImageUrl, address} = property;
  const addr = Address.from(address);
  return (
    <Section title="Images">
      {coverImageUrl ? (
        <div className={s.Images}>
          <PropertyImage imageUrl={coverImageUrl} alt={addr.toString()} cover />
        </div>
      ) : (
        <NoImage />
      )}
    </Section>
  );
}

function PropertyImage({
  imageUrl,
  alt,
  cover,
}: {
  imageUrl: string;
  alt: string;
  cover?: boolean;
}) {
  return (
    <AspectRatio
      as="figure"
      ratio="1:1"
      className={cover ? s.CoverImageContainer : s.ImageContainer}
    >
      <Image
        src={imageUrl}
        alt={alt}
        className={s.Image}
        fill
        sizes={cover ? '24rem' : '12rem'}
      />
    </AspectRatio>
  );
}

function NoImage() {
  return (
    <div className={s.NoImage}>
      <h3 className={s.NoImageTitle}>No images</h3>
    </div>
  );
}
