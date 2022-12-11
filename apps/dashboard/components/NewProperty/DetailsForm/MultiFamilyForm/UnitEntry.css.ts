import {color} from '@/volto/color.css';
import {vars} from '@/volto/root.css';
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
  fontSize: '0.75rem',
  lineHeight: 1,
});

export const ActionButton = style({
  alignItems: 'center',
  background: 'none',
  border: 0,
  borderRadius: vars.border.radius,
  color: color.text.muted,
  cursor: 'pointer',
  display: 'inline-flex',
  justifyContent: 'center',
  padding: unit(1),
  transition: 'color 240ms',
  userSelect: 'none',

  ':hover': {
    backgroundColor: color.primary[95],
    color: color.text.$,
  },
});

export const ActionButtonIcon = style({
  fill: 'currentcolor',
  height: '1em',
  width: '1em',
});

export const Form = style({
  display: 'grid',
  gridTemplateColumns: 'repeat(3, 1fr)',
  gap: unit(4),
});
