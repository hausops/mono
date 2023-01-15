import {useEffect, useId, useState} from 'react';
import {useTooltipsManager} from './TooltipsManager';

export function useTooltip(): {id: string; state: TooltipState} {
  const id = useId();
  const state = useTooltipState(id);
  return {
    id,
    state,
  };
}

type TooltipState = {
  isOpen: boolean;
  open(): void;
  close(): void;
};

function useTooltipState(
  id: string,
  initialOpen: boolean = false
): TooltipState {
  const [isOpen, setOpen] = useState(initialOpen);
  const {visibility} = useTooltipsManager();

  useEffect(
    () => () => {
      visibility.removeCloseFunction(id);
    },
    [id, visibility]
  );

  return {
    isOpen,

    open() {
      visibility.closeAll();
      visibility.addCloseFunction(id, () => setOpen(false));
      setOpen(true);
    },

    close() {
      // need to unsubscribe on close because closing a tooltip
      // does not guarantee to cause an unmount.
      visibility.removeCloseFunction(id);
      setOpen(false);
    },
  };
}
