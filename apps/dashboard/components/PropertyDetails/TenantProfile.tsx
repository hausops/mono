import {Avatar} from '@/volto/Avatar';
import Link from 'next/link';
import * as s from './TenantProfile.css';

type TenantProfileProps = {
  name: string;
  imageUrl: string;
  email: string;
  phone?: string;
};

export function TenantProfile({
  name,
  imageUrl,
  email,
  phone,
}: TenantProfileProps) {
  return (
    <article className={s.TenantProfile}>
      <Link className={s.Avatar} href="#">
        <Avatar name={name} imageUrl={imageUrl} />
      </Link>
      <div className={s.Details}>
        <Link className={s.Name} href="#">
          {name}
        </Link>
        <p className={s.ContactInfo}>
          {/* TODO: copy to clipboard on click, needs toast/notify */}
          <span>{email}</span>
          {phone && <span className={s.Phone}>{phone}</span>}
        </p>
      </div>
    </article>
  );
}
