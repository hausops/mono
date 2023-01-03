import {border} from '@/volto/border.css';
import {color} from '@/volto/color.css';
import {vars} from '@/volto/root.css';
import {unit} from '@/volto/spacing.css';
import {font} from '@/volto/typography.css';
import {style} from '@vanilla-extract/css';

export const RecentPayments = style({});

export const Table = style({
  minWidth: '100%',
});

export const TableHeader = style({
  borderBottom: border.solid(1, color.neutral[95]),
});

export const HeaderLabel = style({
  fontSize: 12,
  fontWeight: font.weight.medium,
  textTransform: 'uppercase',
  color: color.text.muted,
});

export const TableCell = style({
  paddingInline: unit(2),
  paddingBlock: unit(1),
});

export const BadgeCell = style([
  TableCell,
  {
    verticalAlign: 'middle',
    width: '0px',
  },
]);
