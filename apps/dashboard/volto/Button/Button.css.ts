import {border} from '@/volto/border.css';
import {color} from '@/volto/color.css';
import {vars} from '@/volto/root.css';
import {unit} from '@/volto/spacing.css';
import {font} from '@/volto/typography.css';
import {ComplexStyleRule, style, styleVariants} from '@vanilla-extract/css';
import type {ButtonVariant} from './types';

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

export const variant = styleVariants<{
  [K in ButtonVariant]: ComplexStyleRule;
}>({
  contained: {
    backgroundColor: color.primary[35],
    color: 'white',
    ':hover': {
      backgroundColor: color.primary[30],
    },
    ':active': {
      backgroundColor: color.primary[25],
    },
  },
  outlined: {
    border: border.solid(1, color.primary[35]),
    backgroundColor: color.background.transparent,
    color: color.primary[35],
    ':hover': {
      backgroundColor: color.primary[95],
    },
    ':active': {
      backgroundColor: color.primary[90],
    },
  },
  text: {
    backgroundColor: color.background.transparent,
    color: color.primary[35],
    ':hover': {
      backgroundColor: color.primary[95],
    },
    ':active': {
      backgroundColor: color.primary[90],
    },
  },
});

export const IconButton = style([
  base,
  {
    display: 'inline-flex',
    padding: unit(2),
    minHeight: vars.size[36],
    minWidth: vars.size[36],
    borderRadius: '50%',
    // borderRadius: vars.border.radius,
    backgroundColor: color.background.transparent,
    fill: 'currentcolor',
    ':hover': {
      backgroundColor: 'rgb(0 0 0 / 0.04)',
    },
    ':active': {
      backgroundColor: 'rgb(0 0 0 / 0.07)',
    },
  },
]);

// export const icon = style({
//   display: 'inline-block',
//   verticalAlign: 'bottom',
//   height: '1em',
// });
