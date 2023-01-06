import {color} from '@/volto/color.css';
import {vars} from '@/volto/root.css';
import {font} from '@/volto/typography.css';
import {style, styleVariants} from '@vanilla-extract/css';

export const Skeleton = style({});

export const variants = styleVariants({
  title: {
    borderRadius: vars.border.radius,
    backgroundColor: color.neutral[95],
    height: font.size[16],
    width: '100%',
  },
  text: {
    borderRadius: vars.border.radius,
    backgroundColor: color.neutral[95],
    height: font.size[14],
    width: '100%',
  },
});
