import {ReactNode} from 'react';
import * as s from './Page.css';

export default function PageLayout({children}: {children: ReactNode}) {
  return <section className={s.Page}>{children}</section>;
}
