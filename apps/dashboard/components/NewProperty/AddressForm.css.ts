import {unit} from '@/volto/spacing.css';
import {style, styleVariants} from '@vanilla-extract/css';

export const Address = style({
  display: 'grid',
  gap: unit(4),

  '@media': {
    '(min-width: 64rem)': {
      gridTemplateColumns: 'repeat(4, 1fr)',
    },
  },
});

export const gridColumnSpan = styleVariants({
  1: {
    gridColumn: 'span 1 / span 1',
  },
  2: {
    gridColumn: 'span 2 / span 2',
  },
});
