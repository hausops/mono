import * as color from '@/volto/color.css';
import {vars} from '@/volto/root.css';
import {unit} from '@/volto/spacing.css';
import {font} from '@/volto/typography.css';
import {ComplexStyleRule, style, styleVariants} from '@vanilla-extract/css';
import type {ButtonVarient} from './types';

export const base = style({
  alignItems: 'center',
  border: 0,
  borderRadius: vars.border.radius,
  cursor: 'pointer',
  display: 'inline-flex',
  fontWeight: font.weight.medium,
  justifyContent: 'center',
  lineHeight: 1,
  minHeight: vars.size[36],
  minWidth: vars.size[64],
  // outline: 'none',
  padding: `${unit(2)} ${unit(4)}`,
  rowGap: unit(2),
  textDecoration: 'none',
  transition: 'background-color 240ms',
  userSelect: 'none',

  // TBD
  // appearance: 'none',
  // verticalAlign: 'middle',
});

export const varient = styleVariants<{
  [K in ButtonVarient]: ComplexStyleRule;
}>({
  contained: {
    backgroundColor: color.primaryPallete[35],
    color: 'white',
    ':hover': {
      backgroundColor: color.primaryPallete[30],
    },
    ':active': {
      backgroundColor: color.primaryPallete[25],
    },
  },
  outlined: {
    border: `1px solid ${color.primaryPallete[35]}`,
    backgroundColor: 'transparent',
    color: color.primaryPallete[35],
    ':hover': {
      backgroundColor: color.primaryPallete[95],
    },
    ':active': {
      backgroundColor: color.primaryPallete[90],
    },
  },
  text: {
    backgroundColor: 'transparent',
    color: color.primaryPallete[35],
    ':hover': {
      backgroundColor: color.primaryPallete[95],
    },
    ':active': {
      backgroundColor: color.primaryPallete[90],
    },
  },
});

export const icon = style({
  display: 'inline-block',
  verticalAlign: 'bottom',
  height: '1em',
});
