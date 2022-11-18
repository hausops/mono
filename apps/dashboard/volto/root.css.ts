import {createGlobalTheme} from '@vanilla-extract/css';
import * as border from './border.css';
import * as color from './color.css';

export const root = createGlobalTheme(':root', {
  border: {
    divider: `${border.width[1]} solid ${color.divider}`,
  },
});
