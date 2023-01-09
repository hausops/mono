import {unit} from '@/volto/spacing.css';
import {style} from '@vanilla-extract/css';

export const DueDate = style({
  display: 'flex',
  justifyContent: 'space-between',
  alignItems: 'center',
  columnGap: unit(2),
});
