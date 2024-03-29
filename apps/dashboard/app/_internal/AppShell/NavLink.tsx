'use client';

import type {SvgIcon} from '@/volto/icons/SvgIcon';
import clsx from 'clsx';
import Link from 'next/link';
import {usePathname} from 'next/navigation';
import type {ReactElement} from 'react';
import * as s from './NavLink.css';

export type NavLinkProps = {
  // exact matches active state only if the pathname matches the href exactly
  // otherwise (default: false) it matches pathname that starts with href.
  exact?: boolean;
  href: string;
  icon: ReactElement<typeof SvgIcon>;
  text: string;
};

export function NavLink({exact = false, href, icon, text}: NavLinkProps) {
  const pathname = usePathname();
  const active = pathname
    ? exact
      ? pathname === href
      : pathname.startsWith(href)
    : false;

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
