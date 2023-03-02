import {border} from '@/volto/border.css';
import {color} from '@/volto/color.css';
import {unit} from '@/volto/spacing.css';
import {font} from '@/volto/typography.css';
import {style} from '@vanilla-extract/css';

export const Unit = style({
  columnGap: unit(4),
  display: 'grid',
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

export const UnitColumn = style({
  display: 'grid',
  gap: unit(1),
});

export const UnitTitle = style({
  fontSize: font.size[16],
  fontWeight: font.weight.semibold,
});

export const UnitInfo = style({
  alignItems: 'center',
  display: 'flex',
});

export const UnitInfoItem = style({
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
