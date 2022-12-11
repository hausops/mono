const borderWidth = {
  1: '0.0625rem',
  2: '0.125rem',
  3: '0.1875rem',
  4: '0.25rem',
  5: '0.3125rem',
};

export type BorderWidthKey = keyof typeof borderWidth;

export const border = {
  solid,
  width: borderWidth,
};

function solid(width: BorderWidthKey, color: string) {
  const w = borderWidth[width];
  return `${w} solid ${color}`;
}
