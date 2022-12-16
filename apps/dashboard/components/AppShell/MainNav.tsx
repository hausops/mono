// import {Stack} from '@/volto/Stack';

import {
  Chair,
  CreditCard,
  Description,
  HomeFilled,
  Settings,
  SpaceDashboard,
} from '@/volto/icons';
import {PropsWithChildren} from 'react';
import * as s from './MainNav.css';
import {NavLink, NavLinkProps} from './NavLink';

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
      <NavLink href="/settings" icon={Settings} text="Settings" />
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
    icon: SpaceDashboard,
    text: 'Dashboard',
  },
  {
    href: '/properties',
    icon: HomeFilled,
    text: 'Properties',
  },
  {
    href: '/units',
    icon: Chair,
    text: 'Units',
  },
  {
    href: '/applications',
    icon: Description,
    text: 'Applications',
  },
  {
    href: '/payments',
    icon: CreditCard,
    text: 'Payments',
  },
];
