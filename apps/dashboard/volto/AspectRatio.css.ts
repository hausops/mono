import {style, styleVariants} from '@vanilla-extract/css';

const container = style({
  // overflow: 'hidden',
  position: 'relative',
});

export const ratio = styleVariants(
  {
    '1:1': '1',
    '2:1': '2',
    '3:2': '3 / 2',
    '4:3': '4 / 3',
    '16:9': '16 / 9',
  },
  (r) => [container, {aspectRatio: r}]
);
