import {color} from '@/volto/color.css';
import {unit} from '@/volto/spacing.css';
import {style} from '@vanilla-extract/css';

export const UpcomingRent = style({
  display: 'grid',
  gap: unit(2),
});

export const Entry = style({
  display: 'grid',
  gridTemplateColumns: '1fr 1fr',
});

export const EntryLabel = style({
  // color: color.text.muted,
});

export const EntryValue = style({});

export const DueDate = style({
  display: 'flex',
  justifyContent: 'space-between',
  alignItems: 'center',
  columnGap: unit(2),
});
