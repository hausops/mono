import {createGlobalTheme} from '@vanilla-extract/css';
import * as border from './border.css';
import * as color from './color.css';

export const vars = createGlobalTheme(':root', {
  border: {
    divider: `${border.width[1]} solid ${color.divider}`,
    radius: '0.375rem',
  },
  size: {
    36: '2.25rem',
    64: '4rem',
  },
});
