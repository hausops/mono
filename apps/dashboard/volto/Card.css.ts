import {vars} from '@/volto/root.css';
import {style} from '@vanilla-extract/css';

export const Card = style({
  backgroundColor: 'white', // surface
  borderRadius: vars.border.radius,
  overflow: 'auto',
});
