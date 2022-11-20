import {unit} from '@/volto/spacing.css';
import {font} from '@/volto/typography.css';
import {style} from '@vanilla-extract/css';

// TODO: extract to Page layout
export const Page = style({
  display: 'flex',
  flexDirection: 'column',
  padding: `${unit(8)} ${unit(16)}`,
  // TBD
  rowGap: unit(4),
});

// TODO: extract to AppHeader layout
// while at it, we should wrap actions using wrapper with flex-shrink: 0
export const Header = style({
  display: 'flex',
  justifyContent: 'space-between',
  alignItems: 'center',
});

export const Title = style({
  fontSize: font.size[24],
  fontWeight: font.weight.semibold,
});

export const PropertyList = style({
  display: 'grid',
  gap: unit(4),

  '@media': {
    '(min-width: 60rem)': {
      gridTemplateColumns: 'repeat(2, 1fr)',
    },
  },
});
