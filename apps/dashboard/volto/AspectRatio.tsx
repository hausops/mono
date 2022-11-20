import {ElementType, PropsWithChildren, ReactNode} from 'react';
import * as s from './AspectRatio.css';

type AspectRatioProps<As extends ElementType> = PropsWithChildren<{
  as?: As;
  ratio: keyof typeof s.ratio;
}>;

export default function AspectRatio<As extends ElementType = 'div'>({
  as,
  children,
  ratio,
}: AspectRatioProps<As>) {
  const Root = as ?? 'div';
  return <Root className={s.ratio[ratio]}>{children}</Root>;
}
