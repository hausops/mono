import type {SvgIconProps} from '@mui/material';
import clsx from 'clsx';
import Link from 'next/link';
import {useRouter} from 'next/router';
import {ReactElement} from 'react';
import * as s from './NavLink.css';

export type NavLinkProps = {
  href: string;
  icon: ReactElement<SvgIconProps>;
  text: string;
};

export default function NavLink({href, icon, text}: NavLinkProps) {
  const {pathname} = useRouter();
  const className = clsx(s.base, {
    [s.state.active]: href === pathname,
  });
  return (
    <Link className={className} href={href}>
      <span className={s.icon}>{icon}</span>
      {text}
    </Link>
  );
}
