import {layer} from '@/volto/layer.css';
import {unit} from '@/volto/spacing.css';
import {style} from '@vanilla-extract/css';

export const Tooltip = style({
  paddingBlock: unit(2),
  position: 'absolute',
  zIndex: layer.tooltip,
});
