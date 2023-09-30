import {border, type BorderWidthKey} from './border.css';

export const boxShadow = {
  asBorder,
};

function asBorder(width: BorderWidthKey, color: string): string {
  const w = border.width[width];
  return `0 0 0 ${w} ${color}`;
}
