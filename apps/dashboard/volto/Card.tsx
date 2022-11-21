import {ElementType, PropsWithChildren, ReactNode} from 'react';
import * as s from './Card.css';

type CardProps = PropsWithChildren<{
  as?: ElementType<{className: string; children: ReactNode}>;
}>;

export default function Card({as: Root = 'div', children}: CardProps) {
  return <Root className={s.Card}>{children}</Root>;
}
