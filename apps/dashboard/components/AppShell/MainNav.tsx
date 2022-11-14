// import {Stack} from '@/volto/Stack';

import Chair from '@mui/icons-material/Chair';
import CreditCard from '@mui/icons-material/CreditCard';
import Dashboard from '@mui/icons-material/Dashboard';
import Description from '@mui/icons-material/Description';
import Home from '@mui/icons-material/Home';
import Settings from '@mui/icons-material/Settings';
import {PropsWithChildren} from 'react';
import * as s from './MainNav.css';
import NavLink, {NavLinkProps} from './NavLink';

export default function MainNav() {
  return (
    <aside className={s.MainNav}>
      <header className={s.Header}>Logo</header>
      <NavSection>
        {navLinks.map((props, i) => (
          <NavLink key={i} {...props} />
        ))}
      </NavSection>
      <footer className={s.Footer}>
        <NavLink href="/settings" icon={Settings} text="Settings" />
      </footer>
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
    href: '/',
    icon: Dashboard,
    text: 'Dashboard',
  },
  {
    href: '/properties',
    icon: Home,
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
