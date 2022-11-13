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
  const active = href === pathname;
  const linkClassName = clsx(s.link, {
    [s.state.active]: active,
  });
  return (
    <div className={s.root}>
      {active && <ActiveMarker />}
      <Link className={linkClassName} href={href}>
        <span className={s.icon}>{icon}</span>
        {text}
      </Link>
    </div>
  );
}

function ActiveMarker() {
  return (
    <div className={s.activeMarkerContainer}>
      <span className={s.activeMarker} />
    </div>
  );
}
