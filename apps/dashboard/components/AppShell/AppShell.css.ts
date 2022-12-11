import {color} from '@/volto/color.css';
import {font} from '@/volto/typography.css';
import {style} from '@vanilla-extract/css';

export const AppShell = style({
  display: 'flex',
  // alignItems: 'stretch',
  height: '100vh',
  fontSize: font.size[14],
  color: color.text.$,
  backgroundColor: color.background.$,
});

export const Main = style({
  flexGrow: 1,
  overflowY: 'auto',
});
