import {unit} from '@/volto/spacing.css';
import {font} from '@/volto/typography.css';
import {style} from '@vanilla-extract/css';

export const Header = style({
  display: 'flex',
  justifyContent: 'space-between',
  alignItems: 'center',
});

export const Title = style({
  fontSize: font.size[24],
  fontWeight: font.weight.semibold,
});

export const Actions = style({
  display: 'flex',
  alignItems: 'center',
  gap: unit(2),
  flexShrink: 0,
});
