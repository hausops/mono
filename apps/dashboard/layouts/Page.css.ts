import {unit} from '@/volto/spacing.css';
import {style} from '@vanilla-extract/css';

export const Page = style({
  display: 'flex',
  flexDirection: 'column',
  padding: `${unit(8)} ${unit(16)}`,
  // TBD
  rowGap: unit(4),
  minHeight: '100%',
});
