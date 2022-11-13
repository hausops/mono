import * as color from '@/volto/color.css';
import {style} from '@vanilla-extract/css';

export const AppShell = style({
  display: 'flex',
  height: '100vh',
  backgroundColor: color.background.$,
});
