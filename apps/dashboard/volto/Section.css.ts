import {unit} from '@/volto/spacing.css';
import {font} from '@/volto/typography.css';
import {style} from '@vanilla-extract/css';

export const Container = style({
  display: 'flex',
  flexDirection: 'column',
  padding: unit(8),
  rowGap: unit(8),
});

export const Header = style({
  display: 'flex',
  alignItems: 'baseline',
});

export const Title = style({
  flexGrow: 1,
  fontSize: font.size[16],
  fontWeight: font.weight.semibold,
  lineHeight: 1,
});

export const Actions = style({
  flexShrink: 0,
});
