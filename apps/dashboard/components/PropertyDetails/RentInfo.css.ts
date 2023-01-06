import {color} from '@/volto/color.css';
import {style} from '@vanilla-extract/css';

export const TextLink = style({
  color: color.primary[50],
  ':hover': {
    textDecoration: 'underline',
  },
});
