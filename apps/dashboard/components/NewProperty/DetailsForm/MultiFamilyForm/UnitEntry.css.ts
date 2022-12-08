import {unit} from '@/volto/spacing.css';
import {font} from '@/volto/typography.css';
import {style} from '@vanilla-extract/css';

export const Header = style({
  display: 'flex',
  justifyContent: 'space-between',
  alignItems: 'center',
});

export const Title = style({
  fontWeight: font.weight.semibold,
  marginBottom: unit(2),
});

export const Actions = style({
  display: 'flex',
  fontSize: '0.75rem',
  gap: unit(2),
});

export const Form = style({
  display: 'grid',
  gridTemplateColumns: 'repeat(3, 1fr)',
  gap: unit(4),
});
