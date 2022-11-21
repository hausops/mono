import {unit} from '@/volto/spacing.css';
import {style} from '@vanilla-extract/css';

export const PropertyList = style({
  display: 'grid',
  gap: unit(4),

  '@media': {
    '(min-width: 60rem)': {
      gridTemplateColumns: 'repeat(2, 1fr)',
    },
  },
});
