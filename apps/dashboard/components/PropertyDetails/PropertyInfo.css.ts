import {color} from '@/volto/color.css';
import {vars} from '@/volto/root.css';
import {unit} from '@/volto/spacing.css';
import {style} from '@vanilla-extract/css';

export const EditToggle = style({
  alignItems: 'center',
  background: 'none',
  border: 0,
  borderRadius: vars.border.radius,
  color: color.text.muted,
  columnGap: unit(1),
  cursor: 'pointer',
  display: 'inline-flex',
  fontSize: '0.75rem',
  justifyContent: 'center',
  lineHeight: 1,
  padding: unit(1),
  transition: 'color 240ms',
  userSelect: 'none',

  ':hover': {
    backgroundColor: color.primary[95],
    color: color.text.$,
  },
});

export const EditToggleIcon = style({
  fill: 'currentcolor',
  height: '1em',
  width: '1em',
});

export const EditingAttributeList = style({
  gap: unit(8),
});

export const AddressForm = style({
  display: 'grid',
  rowGap: unit(4),
});

export const EditActions = style({
  display: 'flex',
  justifyContent: 'flex-end',
  columnGap: unit(2),
});
