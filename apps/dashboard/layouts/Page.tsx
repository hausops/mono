import type {ReactNode} from 'react';
import * as s from './Page.css';

export function PageLayout({children}: {children: ReactNode}) {
  return <section className={s.Page}>{children}</section>;
}
