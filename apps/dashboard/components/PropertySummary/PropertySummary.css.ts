import {color} from '@/volto/color.css';
import {unit} from '@/volto/spacing.css';
import {font} from '@/volto/typography.css';
import {style} from '@vanilla-extract/css';

export const PropertyDetail = style({});

export const Cover = style({
  backgroundColor: '#ececec', // TODO
});

export const CoverImage = style({
  objectFit: 'cover',
});

export const NoImage = style({
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
  height: '100%',
});

export const Body = style({
  display: 'flex',
  justifyContent: 'space-between',
  alignItems: 'center',
  padding: unit(4),
});

export const Title = style({
  color: color.primary[50],
  fontWeight: font.weight.medium,
  transition: 'color 240ms',
  ':hover': {
    color: 'inherit',
  },
});

export const SkeletonBody = style({
  padding: unit(4),
});
