import {PropsWithChildren} from 'react';
import * as s from './AspectRatio.css';

type AspectRatioProps = PropsWithChildren<{
  ratio: keyof typeof s.ratio;
}>;

export default function AspectRatio({children, ratio}: AspectRatioProps) {
  return (
    <div className={s.ratio[ratio]}>
      <div className={s.media}>{children}</div>
    </div>
  );
}
