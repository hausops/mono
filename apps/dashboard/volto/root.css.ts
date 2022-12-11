import {border} from '@/volto/border.css';
import * as color from '@/volto/color.css';
import {createGlobalTheme} from '@vanilla-extract/css';

export const vars = createGlobalTheme(':root', {
  border: {
    divider: border.solid(1, color.divider),
    radius: '0.375rem',
  },
  size: {
    36: '2.25rem',
    64: '4rem',
  },
});
