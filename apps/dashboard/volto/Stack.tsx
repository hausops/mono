import clsx from 'clsx';
import type {ElementType, PropsWithChildren} from 'react';
import * as s from './Stack.css';

type StackProps = PropsWithChildren<{
  as: ElementType;
  className?: string;
  direction: 'column' | 'row';
  gap?: number;
}>;

export function Stack(props: StackProps) {
  const {as: Root = 'div', direction = 'column', gap, children} = props;
  return (
    <Root className={clsx(s.Stack, s.direction[direction], props.className)}>
      {children}
    </Root>
  );
}
