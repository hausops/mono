import {unit} from '@/volto/spacing.css';
import {style} from '@vanilla-extract/css';

export const TwoColumnsLayout = style({
  display: 'grid',
  columnGap: unit(8),
  rowGap: unit(4),

  '@media': {
    '(min-width: 75rem)': {
      gridTemplateColumns: '1fr 1fr',
    },
  },
});

export const Column = style({
  display: 'flex',
  flexDirection: 'column',
  rowGap: unit(4),
});
