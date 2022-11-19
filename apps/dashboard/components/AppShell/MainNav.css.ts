import {unit} from '@/volto/spacing.css';
import {vars} from '@/volto/root.css';
import {style} from '@vanilla-extract/css';

export const MainNav = style({
  display: 'flex',
  flexDirection: 'column',
  flexShrink: 0,
  gap: unit(4),
  minWidth: unit(60),
  paddingBlock: unit(4),
  borderRight: vars.border.divider,
});

export const Header = style({
  flexShrink: 0,
});

export const Footer = style({
  flexShrink: 0,
});

export const NavSection = style({
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'stretch',
  flexGrow: 1,
  overflow: 'auto',
});
