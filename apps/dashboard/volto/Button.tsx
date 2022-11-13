import {PropsWithChildren} from 'react';
import * as s from './Button.css';

export default function Button({children}: PropsWithChildren<{}>) {
  return <button className={s.base}>{children}</button>;
}

export function IconButton() {}
