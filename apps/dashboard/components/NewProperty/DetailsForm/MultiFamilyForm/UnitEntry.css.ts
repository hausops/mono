import {unit} from '@/volto/spacing.css';
import {font} from '@/volto/typography.css';
import {style} from '@vanilla-extract/css';

export const Header = style({
  display: 'flex',
  justifyContent: 'space-between',
  alignItems: 'center',
  marginBottom: unit(2),
});

export const Title = style({
  fontWeight: font.weight.semibold,
});

export const Actions = style({
  columnGap: unit(2),
  display: 'flex',
});

export const Form = style({
  display: 'grid',
  gridTemplateColumns: 'repeat(3, 1fr)',
  gap: unit(4),
});
