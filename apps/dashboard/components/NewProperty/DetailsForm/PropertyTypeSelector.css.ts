import * as border from '@/volto/border.css';
import {boxShadow} from '@/volto/boxShadow.css';
import * as color from '@/volto/color.css';
import {vars} from '@/volto/root.css';
import {unit} from '@/volto/spacing.css';
import {font} from '@/volto/typography.css';
import {style, styleVariants} from '@vanilla-extract/css';

export const Label = style({
  fontWeight: font.weight.semibold,
  marginBottom: unit(2),
});

export const Options = style({
  display: 'flex',
  columnGap: unit(4),
});

export const Option = style({
  border: `${border.width[1]} solid ${color.neutral[90]}`,
  borderRadius: vars.border.radius,
  cursor: 'pointer',
  flexBasis: '50%',
  maxWidth: '25rem',
  outline: 'none',
  padding: unit(4),
});

export const OptionState = styleVariants({
  selected: {
    borderColor: color.primaryPallete[35],
    boxShadow: boxShadow.asBorder(1, color.primaryPallete[35]),
  },
  focusVisible: {
    boxShadow: [
      boxShadow.asBorder(1, color.primaryPallete[35]),
      boxShadow.asBorder(4, color.primaryPallete[90]),
    ].join(','),
  },
});

export const OptionHeader = style({
  display: 'flex',
  alignItems: 'center',
  gap: unit(2),
  marginBottom: unit(2),
});

export const OptionTitle = style({
  flexGrow: 1,
});

export const OptionDescription = style({
  color: color.text.muted,
});
