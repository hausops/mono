import clsx from 'clsx';
import Image from 'next/image';
import * as s from './Avatar.css';
import {PersonFilled} from './icons';

type AvatarProps = {
  imageUrl?: string;
  name: string;
  size?: keyof typeof s.size;
};

export function Avatar({size = 'medium', imageUrl, name}: AvatarProps) {
  return (
    <figure className={clsx(s.Avatar, s.size[size])}>
      {imageUrl ? (
        <Image
          src={imageUrl}
          alt={name}
          fill
          sizes={s.sizes[size]}
          className={s.Image}
        />
      ) : (
        <Fallback />
      )}
    </figure>
  );
}

function Fallback() {
  return (
    <div className={s.Fallback}>
      <PersonFilled />
    </div>
  );
}
