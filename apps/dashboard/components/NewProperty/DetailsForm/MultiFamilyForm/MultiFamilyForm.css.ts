import {unit} from '@/volto/spacing.css';
import {font} from '@/volto/typography.css';
import {style} from '@vanilla-extract/css';

export const MultiFamilyForm = style({
  display: 'flex',
  flexDirection: 'column',
  rowGap: unit(8),
});

export const Title = style({
  fontWeight: font.weight.semibold,
});

export const UnitEntries = style({
  display: 'flex',
  flexDirection: 'column',
  gap: unit(8),
  paddingLeft: unit(8),
});
