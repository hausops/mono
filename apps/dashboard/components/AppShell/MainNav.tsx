// import {Stack} from '@/volto/Stack';

import {
  ChairIcon,
  CreditCardIcon,
  DescriptionIcon,
  HomeFilledIcon,
  SettingsIcon,
  SpaceDashboardIcon,
} from '@/volto/icons';
import type {PropsWithChildren} from 'react';
import * as s from './MainNav.css';
import {NavLink, type NavLinkProps} from './NavLink';

export function MainNav() {
  return (
    <aside className={s.MainNav}>
      <header className={s.Header}>Logo</header>
      <NavSection>
        {navLinks.map((props, i) => (
          <NavLink key={i} {...props} />
        ))}
      </NavSection>
      <footer className={s.Footer}></footer>
      <NavLink href="/settings" icon={SettingsIcon} text="Settings" />
    </aside>
  );
}

export function NavSection({
  title,
  children,
}: PropsWithChildren<{title?: string}>) {
  // NavSectionTitle
  return <nav className={s.NavSection}>{children}</nav>;
}

const navLinks: NavLinkProps[] = [
  {
    exact: true,
    href: '/',
    icon: SpaceDashboardIcon,
    text: 'Dashboard',
  },
  {
    href: '/properties',
    icon: HomeFilledIcon,
    text: 'Properties',
  },
  {
    href: '/units',
    icon: ChairIcon,
    text: 'Units',
  },
  {
    href: '/applications',
    icon: DescriptionIcon,
    text: 'Applications',
  },
  {
    href: '/payments',
    icon: CreditCardIcon,
    text: 'Payments',
  },
];
