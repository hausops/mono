import {PropsWithChildren} from 'react';
import * as s from './AppShell.css';
import {MainNav} from './MainNav';

export function AppShell({children}: PropsWithChildren<{}>) {
  return (
    <div className={s.AppShell}>
      <MainNav />
      <main className={s.Main}>{children}</main>
    </div>
  );
}
