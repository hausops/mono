import clsx from 'clsx';
import Link from 'next/link';
import {useRouter} from 'next/router';
import type {ComponentType} from 'react';
import * as s from './NavLink.css';

export type NavLinkProps = {
  // exact matches active state only if the pathname matches the href exactly
  // otherwise (default: false) it matches pathname that starts with href.
  exact?: boolean;
  href: string;
  icon: ComponentType;
  text: string;
};

export function NavLink({exact = false, href, icon: Icon, text}: NavLinkProps) {
  const {pathname} = useRouter();
  const active = exact ? pathname === href : pathname.startsWith(href);
  const linkClassName = clsx(s.link, {
    [s.state.active]: active,
  });
  return (
    <div className={s.root}>
      {active && <ActiveMarker />}
      <Link className={linkClassName} href={href}>
        <span className={s.icon}>
          <Icon />
        </span>
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
