import {color} from '@/volto/color.css';
import {unit} from '@/volto/spacing.css';
import {font} from '@/volto/typography.css';
import {style} from '@vanilla-extract/css';

export const Content = style({
  backgroundColor: color.neutral[35],
  color: color.neutral[98],
  fontSize: font.size[14],
  paddingBlock: unit(2),
  paddingInline: unit(4),
  // width: 'max-content',
});

export const Copied = style({
  display: 'flex',
  alignItems: 'center',
  gap: unit(1),
});

export const CopiedIcon = style({
  display: 'inline-flex',
  fill: 'currentcolor',
  height: unit(5),
  width: unit(5),
});
