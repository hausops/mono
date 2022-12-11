import * as border from './border.css';

export const boxShadow = {
  asBorder,
};

function asBorder(width: keyof typeof border.width, color: string): string {
  const w = border.width[width];
  return `0 0 0 ${w} ${color}`;
}
