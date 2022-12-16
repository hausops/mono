import clsx from 'clsx';
import * as s from './Skeleton.css';

type SkeletonProps = {
  variant?: keyof typeof s.variants;
  width?: string;
};

export function Skeleton({variant = 'text', width}: SkeletonProps) {
  const className = clsx(s.Skeleton, s.variants[variant]);
  const style = {maxWidth: width};
  return <div className={className} style={style} />;
}
