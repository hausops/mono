import {border} from '@/volto/border.css';
import {color} from '@/volto/color.css';
import {vars} from '@/volto/root.css';
import {unit} from '@/volto/spacing.css';
import {style} from '@vanilla-extract/css';

export const Images = style({
  display: 'grid',
  gridTemplateColumns: 'repeat(4, 1fr)',
  gap: unit(1),
  userSelect: 'none',
});

export const ImageContainer = style({
  border: border.solid(1, color.neutral[90]),
  borderRadius: vars.border.radius,
  overflow: 'hidden',
});

export const CoverImageContainer = style([
  ImageContainer,
  {
    gridColumn: '1/span 2',
    gridRow: '1/span 2',
  },
]);

export const Image = style({
  objectFit: 'contain',
});

export const NoImage = style({
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
  padding: unit(8),
  border: `${border.width[1]} dashed ${color.neutral[90]}`,
  borderRadius: vars.border.radius,
});

export const NoImageTitle = style({
  color: color.text.muted,
});
