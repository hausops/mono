import {border} from '@/volto/border.css';
import {color} from '@/volto/color.css';
import {unit} from '@/volto/spacing.css';
import {font} from '@/volto/typography.css';
import {style} from '@vanilla-extract/css';

export const Unit = style({
  display: 'grid',
  gap: unit(4),
  gridTemplateColumns: '1fr 1fr auto',
  selectors: {
    // line separator between Unit
    '&:not(:first-of-type)': {
      borderBlockStart: border.solid(1, color.neutral[95]),
      marginBlockStart: unit(4),
      paddingBlockStart: unit(4),
    },
  },
});

export const Column = style({
  display: 'grid',
  alignItems: 'start',
  gap: unit(1),
});

export const Title = style({
  lineHeight: 1,
  fontSize: font.size[16],
  fontWeight: font.weight.semibold,
});

export const Info = style({
  alignItems: 'center',
  display: 'flex',
});

export const InfoItem = style({
  selectors: {
    '&:not(:first-of-type)::before': {
      content: '•',
      display: 'inline-flex',
      marginInline: unit(1),
    },
  },
});

export const TextLink = style({
  color: color.primary[50],
  ':hover': {
    textDecoration: 'underline',
  },
});
