import * as s from './TextSkeleton.css';

type TextSkeletonProps = {
  width?: string;
};

export function TextSkeleton({width}: TextSkeletonProps) {
  const style = {maxWidth: width};
  return (
    <div className={s.line} style={style}>
      <span className={s.text}></span>
    </div>
  );
}
