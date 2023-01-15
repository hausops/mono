import {color} from '@/volto/color.css';
import {layer} from '@/volto/layer.css';
import {vars} from '@/volto/root.css';
import {unit} from '@/volto/spacing.css';
import {font} from '@/volto/typography.css';
import {style} from '@vanilla-extract/css';

export const TenantProfile = style({
  display: 'inline-flex',
  alignItems: 'center',
  columnGap: unit(3),
});

export const Avatar = style({
  display: 'block',
  flexShrink: 0,
});

export const Details = style({
  flexGrow: 1,
});

export const Name = style({
  fontWeight: font.weight.semibold,
  ':hover': {
    color: color.text.primary.$,
  },
});

export const ContactInfo = style({
  alignItems: 'center',
  columnGap: unit(1),
  display: 'flex',
});

export const TooltipContainer = style({
  position: 'absolute',
  paddingBlock: unit(2),
  zIndex: layer.tooltip,
});

export const ContactTooltipContent = style({
  backgroundColor: color.neutral[35],
  color: color.neutral[98],
  fontSize: font.size[14],
  paddingBlock: unit(2),
  paddingInline: unit(4),
  // width: 'max-content',
});

export const Contact = style({
  backgroundColor: 'transparent',
  border: 0,
  borderRadius: vars.border.radius,
  color: color.text.muted,
  cursor: 'pointer',
  display: 'inline-flex',

  ':hover': {
    backgroundColor: color.neutral[98],
    color: color.text.$,
  },

  ':focus-visible': {
    backgroundColor: color.neutral[98],
    color: color.text.$,
  },
});

export const CopiedContact = style({
  display: 'flex',
  alignItems: 'center',
  gap: unit(1),
  // width: 'max-content',
});

export const CopiedContactIcon = style({
  display: 'inline-flex',
  fill: 'currentcolor',
  height: unit(5),
  width: unit(5),
});

export const ContactSeparator = style({
  color: color.text.muted,
  display: 'inline-flex',
  '::before': {
    content: ' • ',
  },
});
