import {style} from '@vanilla-extract/css';

export const Stack = style({
  display: 'flex',
  flexDirection: 'column',
  // alignItems: 'stretch',
});

export const direction = {
  column: style({flexDirection: 'column'}),
  row: style({flexDirection: 'row'}),
};
