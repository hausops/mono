import {color} from '@/volto/color.css';
import {vars} from '@/volto/root.css';
import {font} from '@/volto/typography.css';
import {style, styleVariants} from '@vanilla-extract/css';

export const Skeleton = style({});

export const variants = styleVariants({
  text: {
    borderRadius: vars.border.radius,
    backgroundColor: color.neutral[95],
    height: font.size[14],
    width: '100%',
  },
});

export const TextSkeletonLine = style({
  display: 'flex',
  alignItems: 'center',
  height: `${(14 * 1.5) / 16}rem`,
});

export const TextSkeletonContent = style({
  display: 'block',
  borderRadius: vars.border.radius,
  backgroundColor: color.neutral[95],
  height: font.size[14],
  width: '100%',
});
