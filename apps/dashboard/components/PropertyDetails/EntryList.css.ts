import {unit} from '@/volto/spacing.css';
import {style} from '@vanilla-extract/css';

export const EntryList = style({
  display: 'grid',
  gap: unit(2),
});

export const Entry = style({
  display: 'grid',
  gridTemplateColumns: '1fr 1fr',
});
