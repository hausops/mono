import {border} from '@/volto/border.css';
import {boxShadow} from '@/volto/boxShadow.css';
import {color} from '@/volto/color.css';
import {vars} from '@/volto/root.css';
import {unit} from '@/volto/spacing.css';
import {style} from '@vanilla-extract/css';

export const Select = style({
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

export const InputWrapper = style({
  position: 'relative',
});

export const Input = style([
  {
    appearance: 'none', // remove user agent caret
    backgroundColor: 'transparent', // reset for Firefox
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

export const ExpandIcon = style({
  position: 'absolute',
  insetBlock: 0,
  right: unit(2),
  display: 'flex',
  alignItems: 'center',
  width: '1rem',
  fill: color.text.muted,
});
