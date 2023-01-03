import {color} from '@/volto/color.css';
import {unit} from '@/volto/spacing.css';
import {style, styleVariants} from '@vanilla-extract/css';

export const Avatar = style({
  position: 'relative', // for Image fill props
  borderRadius: '50%',
  overflow: 'hidden',
});

export const sizes = {
  small: '2rem',
  medium: '3rem',
  large: '4rem',
};

export const size = styleVariants(sizes, (v) => ({
  height: v,
  width: v,
}));

export const Image = style({
  objectFit: 'cover',
});

export const Fallback = style({
  backgroundColor: color.neutral[95],
  fill: color.neutral[80],
  padding: unit(1),
});
