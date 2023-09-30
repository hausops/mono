import {useEffect, useState} from 'react';
import {useTooltipsManager} from './TooltipsManager';

export type TooltipState = {
  isOpen: boolean;
  open(): void;
  close(): void;
};

export function useTooltipState(
  // id is used to identify the tooltip with the TooltipsManager
  id: string,
  initialOpen: boolean = false,
): TooltipState {
  const [isOpen, setOpen] = useState(initialOpen);
  const {visibility} = useTooltipsManager();

  function open(): void {
    visibility.closeAll();
    visibility.registerCloseTooltip(id, close);
    setOpen(true);
  }

  function close(): void {
    visibility.deregisterCloseTooltip(id);
    setOpen(false);
  }

  useEffect(
    () => () => {
      visibility.deregisterCloseTooltip(id);
    },
    [id, visibility],
  );

  return {
    isOpen,
    open,
    close,
  };
}
