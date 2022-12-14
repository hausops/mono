import {border} from '@/volto/border.css';
import {boxShadow} from '@/volto/boxShadow.css';
import {color} from '@/volto/color.css';
import {vars} from '@/volto/root.css';
import {unit} from '@/volto/spacing.css';
import {style} from '@vanilla-extract/css';

export const TextField = style({
  display: 'flex',
  flexDirection: 'column',
});

export const Label = style({
  lineHeight: 1,
  paddingBottom: unit(2), // extend touch target
});

const inputFocus = style({
  outline: 'none',

  ':focus-visible': {
    borderColor: color.primary[50],
    boxShadow: boxShadow.asBorder(1, color.primary[50]),
  },
});

export const Input = style([
  {
    border: border.solid(1, color.neutral[90]),
    borderRadius: vars.border.radius,
    lineHeight: 1,
    minHeight: vars.size[36],
    padding: unit(2),
    width: '100%',

    '::placeholder': {
      color: color.text.muted,
    },
  },
  inputFocus,
]);
