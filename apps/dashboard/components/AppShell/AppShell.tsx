import {PropsWithChildren} from 'react';
import * as s from './AppShell.css';
import MainNav from './MainNav';

export default function AppFrame({children}: PropsWithChildren<{}>) {
  return (
    <div className={s.AppShell}>
      <MainNav />
      <main>{children}</main>
    </div>
  );
}
