import {color} from '@/volto/color.css';
import {vars} from '@/volto/root.css';
import {style} from '@vanilla-extract/css';

export const Card = style({
  backgroundColor: color.surface.$,
  borderRadius: vars.border.radius,
  overflow: 'auto',
});
