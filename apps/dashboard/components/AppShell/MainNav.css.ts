import {unit} from '@/volto/spacing.css';
import {style} from '@vanilla-extract/css';

export const MainNav = style({
  display: 'flex',
  flexDirection: 'column',
  gap: unit(4),
  minWidth: unit(60),
  paddingBlock: unit(4),
  paddingInline: unit(2),

  borderRight: '0.0625rem solid #e1e3e5',
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
