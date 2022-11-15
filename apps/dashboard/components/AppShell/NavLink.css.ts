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
  // backgroundColor: color.brand,
  borderBottomRightRadius: unit(1),
  borderTopRightRadius: unit(1),
});

export const link = style({
  display: 'flex',
  alignItems: 'center',
  gap: unit(2),
  paddingInline: unit(2),
  borderRadius: unit(1),

  lineHeight: unit(8),
  fontSize: font.size[14],
  fontWeight: font.weight.medium,

  ':hover': {
    backgroundColor: color.background.hovered,
  },
});

export const state = styleVariants({
  active: {
    backgroundColor: color.background.selected,
    color: color.text.primary.$,
    // color: color.brand,
  },
});

export const icon = style({
  display: 'inline-flex',
  fill: 'currentcolor',
  height: unit(5),
  width: unit(5),
});
