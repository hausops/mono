import {color} from '@/volto/color.css';
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

export const ContactSeparator = style({
  color: color.text.muted,
  display: 'inline-flex',
  '::before': {
    content: '•',
  },
});
