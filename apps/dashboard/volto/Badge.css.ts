import {color} from '@/volto/color.css';
import {vars} from '@/volto/root.css';
import {unit} from '@/volto/spacing.css';
import {font} from '@/volto/typography.css';
import {style, styleVariants} from '@vanilla-extract/css';

export const Badge = style({
  display: 'inline-flex',
  alignItems: 'center',
  borderRadius: vars.border.radius,
  paddingInline: unit(1),
  fontSize: font.size[12],
  // lineHeight: '1.25rem',
});

export const status = styleVariants({
  default: {
    backgroundColor: color.neutral[95],
  },
  attention: {
    backgroundColor: '#ffe359',
  },
});

export const LivenessBadge = style({
  alignItems: 'center',
  display: 'inline-flex',
  gap: unit(2),
  fontWeight: font.weight.medium,
});

export const LivenessBadgeStatus = styleVariants(
  {
    idle: {
      foreground: color.neutral[90],
      background: color.neutral[98],
    },
    live: {
      foreground: '#36ac52',
      background: '#36ac5233', // 0.2 alpha
    },
  },
  ({foreground, background}) => ({
    backgroundColor: background,
    borderRadius: '50%',
    display: 'inline-flex',
    flexShrink: 0,
    padding: unit(1),
    '::before': {
      backgroundColor: foreground,
      borderRadius: '50%',
      content: '',
      height: unit(2),
      width: unit(2),
    },
  })
);
