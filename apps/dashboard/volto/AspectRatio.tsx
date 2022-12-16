import clsx from 'clsx';
import {ElementType, PropsWithChildren, ReactNode} from 'react';
import * as s from './AspectRatio.css';

type AspectRatioProps = PropsWithChildren<{
  as?: ElementType<{className: string; children: ReactNode}>;
  className?: string;
  ratio: keyof typeof s.ratio;
}>;

export function AspectRatio({
  as: Root = 'div',
  children,
  className,
  ratio,
}: AspectRatioProps) {
  return <Root className={clsx(s.ratio[ratio], className)}>{children}</Root>;
}
