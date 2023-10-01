import {AspectRatio} from '@/volto/AspectRatio';
import {Card} from '@/volto/Card';
import {TextSkeleton} from '@/volto/Skeleton';
import * as s from './_internal/PropertySummary.css';

export default function LoadingPropertyList() {
  return Array.from({length: 8}).map((_, i) => (
    <li key={i}>
      <PropertySummarySkeleton />
    </li>
  ));
}

function PropertySummarySkeleton() {
  return (
    <Card as="article">
      <AspectRatio as="figure" ratio="2:1" className={s.Cover} />
      <div className={s.SkeletonBody}>
        <TextSkeleton width="50%" />
        <TextSkeleton width="70%" />
      </div>
    </Card>
  );
}
