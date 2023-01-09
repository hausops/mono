import {color} from '@/volto/color.css';
import {vars} from '@/volto/root.css';
import {unit} from '@/volto/spacing.css';
import {font} from '@/volto/typography.css';
import {style, styleVariants} from '@vanilla-extract/css';

export const Badge = style({
  display: 'inline-flex',
  alignItems: 'center',
  borderRadius: vars.border.radius,
  paddingInline: unit(1),
  fontSize: font.size[12],
  // lineHeight: '1.25rem',
});

export const status = styleVariants({
  default: {
    backgroundColor: color.neutral[95],
  },
  attention: {
    backgroundColor: '#ffe359',
  },
});
