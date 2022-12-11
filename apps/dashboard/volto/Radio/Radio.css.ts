import {border} from '@/volto/border.css';
import * as color from '@/volto/color.css';
import {style} from '@vanilla-extract/css';

export const Radio = style({
  appearance: 'none',
  cursor: 'pointer',
  position: 'relative',
  width: '1em',
  height: '1em',
  padding: '0.2em',
  borderRadius: '50%',
  border: border.solid(1, color.neutral[70]),

  ':focus': {
    outline: 'none',
  },

  '::before': {
    content: '',
    position: 'absolute',
    inset: '0.15em',
    backgroundColor: 'transparent',
    borderRadius: '50%',
  },

  ':checked': {
    borderColor: color.primaryPallete[35],
  },

  selectors: {
    '&:checked::before': {
      backgroundColor: color.primaryPallete[35],
    },
  },
});
