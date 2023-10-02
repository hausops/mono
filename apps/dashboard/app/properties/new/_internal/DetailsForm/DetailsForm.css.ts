import {color} from '@/volto/color.css';
import {vars} from '@/volto/root.css';
import {unit} from '@/volto/spacing.css';
import {style, styleVariants} from '@vanilla-extract/css';

export const PropertyTypeButtonGroup = style({
  display: 'flex',
  padding: unit(1),
  gap: unit(1),
  backgroundColor: color.neutral[95],
  borderRadius: vars.border.radius,
  width: 'fit-content',
});

export const PropertyTypeButton = styleVariants({
  default: {
    backgroundColor: 'transparent',
    ':hover': {backgroundColor: color.neutral[98]},
  },
  selected: {
    backgroundColor: color.neutral[100],
  },
});
