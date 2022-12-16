import {color} from '@/volto/color.css';
import {vars} from '@/volto/root.css';
import {font} from '@/volto/typography.css';
import {style} from '@vanilla-extract/css';

const REM_BASE = 16;
const LINE_HEIGHT = 1.5;

export const line = style({
  display: 'flex',
  alignItems: 'center',
  height: `${(14 / REM_BASE) * LINE_HEIGHT}rem`,
});

export const text = style({
  display: 'block',
  borderRadius: vars.border.radius,
  backgroundColor: color.neutral[95],
  width: '100%',
  height: font.size[14],
});
