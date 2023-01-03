import {color} from '@/volto/color.css';
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
  color: color.text.muted,
});

export const Phone = style({
  '::before': {
    content: ' • ',
  },
});
