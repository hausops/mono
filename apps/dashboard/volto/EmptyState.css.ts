import {color} from '@/volto/color.css';
import {unit} from '@/volto/spacing.css';
import {font} from '@/volto/typography.css';
import {style} from '@vanilla-extract/css';

export const EmptyState = style({
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center',
  rowGap: unit(4),
});

export const Icon = style({
  fill: color.neutral[90],
  height: '3rem',
  width: '3rem',
});

export const Body = style({
  textAlign: 'center',
});

export const Title = style({
  fontSize: font.size[16],
  fontWeight: font.weight.semibold,
});

export const Description = style({
  color: color.text.muted,
});

export const Actions = style({
  display: 'flex',
  columnGap: unit(2),
});
