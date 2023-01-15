import {useId} from 'react';
import {TooltipState, useTooltipState} from './useTooltipState';

export function useTooltip(): {id: string; state: TooltipState} {
  const id = useId();
  const state = useTooltipState(id);
  return {
    id,
    state,
  };
}
