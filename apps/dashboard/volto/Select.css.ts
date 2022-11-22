import * as border from '@/volto/border.css';
import * as color from '@/volto/color.css';
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
  outline: `${border.width[2]} solid transparent`,
  outlineOffset: `-${border.width[1]}`,

  ':focus': {
    borderRadius: vars.border.radius,
    borderColor: 'transparent',
    outlineColor: color.primaryPallete[50],
  },
});

export const InputWrapper = style({
  position: 'relative',
});

export const Input = style([
  {
    appearance: 'none', // remove user agent caret
    border: `${border.width[1]} solid ${color.neutral[90]}`,
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
