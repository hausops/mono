import * as color from '@/volto/color.css';
import {unit} from '@/volto/spacing.css';
import {font} from '@/volto/typography.css';
import {style, styleVariants} from '@vanilla-extract/css';

export const root = style({
  position: 'relative',
  paddingInline: unit(2),
});

export const activeMarkerContainer = style({
  position: 'absolute',
  top: 0,
  bottom: 0,
  left: 0,
  display: 'flex',
  alignItems: 'center',
});

export const activeMarker = style({
  height: '80%',
  width: 3,
  backgroundColor: color.text.primary.$,
  borderBottomRightRadius: unit(1),
  borderTopRightRadius: unit(1),
});

export const link = style({
  display: 'flex',
  alignItems: 'center',
  gap: unit(2),
  paddingBlock: unit(1),
  paddingInline: unit(2),
  borderRadius: unit(1),

  fontSize: '0.875rem',
  fontWeight: font.weight.medium,
  color: color.text.$,

  ':hover': {
    backgroundColor: color.background.hovered,
  },
});

export const state = styleVariants({
  active: {
    color: color.text.primary.$,
    backgroundColor: color.background.selected,
  },
});

export const icon = style({
  display: 'inline-flex',
  fontSize: unit(5),
});
