import {vars} from '@/volto/root.css';
import {unit} from '@/volto/spacing.css';
import {style, styleVariants} from '@vanilla-extract/css';
import {color} from './color.css';

export const Badge = style({
  display: 'inline-flex',
  alignItems: 'center',
  borderRadius: vars.border.radius,
  paddingInline: unit(1),
  fontSize: '0.75rem',
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
