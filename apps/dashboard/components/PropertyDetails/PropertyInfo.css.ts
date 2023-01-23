import {unit} from '@/volto/spacing.css';
import {style} from '@vanilla-extract/css';

export const EditingAttributeList = style({
  gap: unit(8),
});

export const EditActions = style({
  display: 'flex',
  justifyContent: 'flex-end',
  columnGap: unit(2),
});
