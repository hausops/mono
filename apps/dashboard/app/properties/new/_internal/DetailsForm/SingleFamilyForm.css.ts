import {unit} from '@/volto/spacing.css';
import {style} from '@vanilla-extract/css';

export const SingleFamilyForm = style({
  display: 'grid',
  gridTemplateColumns: 'repeat(3, 1fr)',
  gap: unit(4),
});
