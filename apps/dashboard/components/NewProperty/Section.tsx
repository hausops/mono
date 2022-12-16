import {Card} from '@/volto/Card';
import {PropsWithChildren} from 'react';
import * as s from './Section.css';

type SectionProps = PropsWithChildren<{
  title: string;
}>;

export function Section({children, title}: SectionProps) {
  return (
    <Card>
      <div className={s.Container}>
        <h2 className={s.Title}>{title}</h2>
        {children}
      </div>
    </Card>
  );
}
