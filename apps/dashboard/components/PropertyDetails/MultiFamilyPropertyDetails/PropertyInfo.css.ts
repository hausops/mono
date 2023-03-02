import {unit} from '@/volto/spacing.css';
import {style} from '@vanilla-extract/css';

export const Address = style({
  display: 'flex',
  alignItems: 'center',
  columnGap: unit(1),
});

export const AddressIcon = style({
  display: 'inline-flex',
  flexShrink: 0,
  fill: 'currentcolor',
  width: '1rem',
  height: '1rem',
});

export const EditActions = style({
  display: 'flex',
  justifyContent: 'flex-end',
  columnGap: unit(2),
});
