import {Address} from '@/services/address';
import {SingleFamilyProperty} from '@/services/property';
import {AspectRatio} from '@/volto/AspectRatio';
import {Section} from '@/volto/Section';
import Image from 'next/image';
import {Entry, EntryList} from './EntryList';
import * as s from './PropertyInfo.css';

type PropertyInfoProps = {
  property: SingleFamilyProperty;
};

export function PropertyInfo({property}: PropertyInfoProps) {
  const {coverImageUrl, address} = property;
  const addr = Address.from(address);
  return (
    <>
      <Section title="Property info">
        <PropertyInfoSection property={property} />
      </Section>

      <Section title="Images">
        <div className={s.Images}>
          {coverImageUrl && (
            <PropertyImage
              imageUrl={coverImageUrl}
              alt={addr.toString()}
              cover
            />
          )}
        </div>
      </Section>
    </>
  );
}

function PropertyInfoSection({property}: {property: SingleFamilyProperty}) {
  const {address, unit} = property;
  const [street, region] = Address.from(address).format();
  return (
    <EntryList>
      <Entry
        label="Address"
        value={
          <div>
            <p>{street}</p>
            <p>{region}</p>
          </div>
        }
      />
      {unit.bedrooms ? <Entry label="Beds" value={unit.bedrooms} /> : null}
      {unit.bathrooms ? <Entry label="Baths" value={unit.bathrooms} /> : null}
      {unit.size ? (
        <Entry
          label="Size"
          value={`${Intl.NumberFormat('en-US').format(unit.size)} sq.ft.`}
        />
      ) : null}
    </EntryList>
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
        fill
        className={s.Image}
        sizes={cover ? '24rem' : '12rem'}
      />
    </AspectRatio>
  );
}
