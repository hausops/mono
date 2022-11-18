import {style} from '@vanilla-extract/css';

export const base = style({
  appearance: 'none',
  border: 0,
  lineHeight: 1,
  // fontWeight: 500,
  textDecoration: 'none',
  display: 'inline-flex',
  alignItems: 'center',
  userSelect: 'none',
  // outline: 'none',
  rowGap: '0.5rem',
});

export const icon = style({
  display: 'inline-block',
  verticalAlign: 'bottom',
  height: '1em',
});
