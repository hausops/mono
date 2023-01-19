import {useId} from 'react';
import {TooltipState, useTooltipState} from './useTooltipState';

type TriggerProps = {
  'aria-describedby'?: string;
  onFocus: () => void;
  onBlur: () => void;
  onMouseEnter: () => void;
  onMouseLeave: () => void;
};

export function useTooltip(): {
  id: string;
  state: TooltipState;
  triggerProps: TriggerProps;
} {
  const id = useId();
  const state = useTooltipState(id);
  const triggerProps = {
    'aria-describedby': state.isOpen ? id : undefined,
    onFocus: state.open,
    onBlur: state.close,
    onMouseEnter: state.open,
    onMouseLeave: state.close,
  };

  return {
    id,
    state,
    triggerProps,
  };
}
