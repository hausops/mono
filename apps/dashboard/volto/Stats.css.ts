import {color} from '@/volto/color.css';
import {font} from '@/volto/typography.css';
import {style} from '@vanilla-extract/css';

export const Stats = style({
  display: 'flex',
  flexDirection: 'column',
});

export const DisplayValue = style({
  fontWeight: font.weight.semibold,
  fontSize: '1.375rem',
  lineHeight: '2rem',
});

export const Unit = style({
  color: color.text.muted,
});
