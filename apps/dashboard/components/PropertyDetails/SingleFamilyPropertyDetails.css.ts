import {color} from '@/volto/color.css';
import {vars} from '@/volto/root.css';
import {unit} from '@/volto/spacing.css';
import {font} from '@/volto/typography.css';
import {style} from '@vanilla-extract/css';

export const SingleFamilyPropertyDetails = style({
  display: 'grid',
  gap: unit(8),

  '@media': {
    '(min-width: 75rem)': {
      // gridTemplateColumns: '2fr 1fr',
      gridTemplateColumns: '1fr 1fr',
    },
  },
});

export const Column = style({
  display: 'grid',
  rowGap: unit(4),
});

export const TextLink = style({
  color: color.primary[50],
  ':hover': {
    textDecoration: 'underline',
  },
});
