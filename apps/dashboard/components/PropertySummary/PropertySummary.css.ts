import * as color from '@/volto/color.css';
import {unit} from '@/volto/spacing.css';
import {font} from '@/volto/typography.css';
import {style} from '@vanilla-extract/css';

export const PropertyDetail = style({});

export const CoverImage = style({
  // need width & height 100% so object-fit works
  width: '100%',
  height: '100%',
  objectFit: 'cover',
  objectPosition: 'center',
});

export const NoImage = style({
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
  height: '100%',
  backgroundColor: '#ececec', // TODO
});

export const Body = style({
  display: 'flex',
  justifyContent: 'space-between',
  alignItems: 'center',
  padding: unit(4),
});

export const Title = style({
  color: color.primaryPallete[50],
  // fontSize: '1rem',
  fontWeight: font.weight.medium,
  transition: 'color 240ms',
  ':hover': {
    color: 'inherit',
  },
});
