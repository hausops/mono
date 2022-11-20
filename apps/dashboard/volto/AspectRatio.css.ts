import {style, styleVariants} from '@vanilla-extract/css';

const container = style({
  position: 'relative',
});

export const ratio = styleVariants(
  {
    '1:1': 100,
    '2:1': 50,
    '3:2': 66.67,
    '4:3': 75,
    '16:9': 56.25,
  },
  (percent) => [container, {paddingTop: `${percent}%`}]
);

export const media = style({
  position: 'absolute',
  inset: 0,
  overflow: 'hidden',
});
