import {Card} from '@/volto/Card';
import {CSSProperties, PropsWithChildren, ReactNode} from 'react';
import * as s from './Section.css';
import {Skeleton} from './Skeleton';

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

type SectionSkeletonProps = {
  height?: CSSProperties['height'];
};

export function SectionSkeleton({height}: SectionSkeletonProps) {
  return (
    <Card as="section">
      <div className={s.Container} style={{minHeight: height}}>
        <header className={s.Header}>
          <h2 className={s.Title}>
            <Skeleton variant="title" width="35%" />
          </h2>
        </header>
        <div className={s.SectionSkeletonBody}>
          <Skeleton width="75%" />
          <Skeleton width="50%" />
        </div>
      </div>
    </Card>
  );
}
