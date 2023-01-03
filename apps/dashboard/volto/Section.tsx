import {Card} from '@/volto/Card';
import {PropsWithChildren, ReactNode} from 'react';
import * as s from './Section.css';

type SectionProps = PropsWithChildren<{
  title: string;
  actions?: ReactNode;
}>;

// Section is used to group related elements on the page together.
//
// It essentially is a Card with standard padding and title styling
// with optional actions.
export function Section({children, title, actions}: SectionProps) {
  return (
    <Card as="section">
      <div className={s.Container}>
        <header className={s.Header}>
          <h2 className={s.Title}>{title}</h2>
          {actions && <div className={s.Actions}>{actions}</div>}
        </header>
        {children}
      </div>
    </Card>
  );
}
