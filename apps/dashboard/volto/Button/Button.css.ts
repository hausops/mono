import {border} from '@/volto/border.css';
import {color} from '@/volto/color.css';
import {vars} from '@/volto/root.css';
import {unit} from '@/volto/spacing.css';
import {font} from '@/volto/typography.css';
import {style, styleVariants} from '@vanilla-extract/css';

const base = style({
  alignItems: 'center',
  background: 'none',
  border: 0,
  borderRadius: vars.border.radius,
  columnGap: unit(1),
  cursor: 'pointer',
  display: 'inline-flex',
  justifyContent: 'center',
  lineHeight: 1,
  textDecoration: 'none',
  transition: 'background-color 240ms',
  userSelect: 'none',

  // TBD
  // appearance: 'none',
  // outline: 'none',
});

export const Button = style([
  base,
  {
    fontWeight: font.weight.medium,
    minHeight: vars.size[36],
    minWidth: vars.size[64],
    padding: `${unit(2)} ${unit(4)}`,
  },
]);

export const ButtonVariants = styleVariants({
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
    padding: unit(2),
    minHeight: vars.size[36],
    minWidth: vars.size[36],
    borderRadius: '50%',
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

export const MiniTextButton = style([
  base,
  {
    color: color.text.muted,
    fontSize: font.size[12],
    padding: unit(1),
    transition: 'color 240ms',

    ':hover': {
      backgroundColor: color.primary[95],
      color: color.text.$,
    },
  },
]);

export const MiniTextButtonIcon = style({
  fill: 'currentcolor',
  flexShrink: 0,
  height: '1em',
  width: '1em',
});
