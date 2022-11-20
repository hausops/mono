import {PropsWithChildren} from 'react';
import * as s from './Card.css';

export default function Card({children}: PropsWithChildren<{}>) {
  return <div className={s.Card}>{children}</div>;
}
