import {style} from '@vanilla-extract/css';
import {unit} from './spacing.css';

export const ButtonGroup = style({
  display: 'flex',
  justifyItems: 'flex-end',
  gap: unit(2),
});
