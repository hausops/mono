import {
  createContext,
  forwardRef,
  PropsWithChildren,
  useContext,
  useRef,
  useState,
} from 'react';
import * as s from './TooltipsManager.css';

// TooltipsManager manages states across different tooltips such as ensuring
// that only one tooltip is open at a given time.
//
// It also provides the portal element for tooltips.
type TooltipsManager = {
  portalsContainer: HTMLElement | null;
  visibility: VisibilityManager;
};

type TooltipsManagerProviderProps = PropsWithChildren<{}>;

export function TooltipsManagerProvider(props: TooltipsManagerProviderProps) {
  const {children} = props;
  // Cannot use useRef here because it does not cause a re-render
  // thus the context is not updated after the container mounted.
  const [portalsContainer, setPortalsContainer] = useState<HTMLElement | null>(
    null,
  );

  const {current: visibility} = useRef(new VisibilityManager());

  const tooltipsManager: TooltipsManager = {
    portalsContainer,
    visibility,
  };

  return (
    <TooltipsManagerContext.Provider value={tooltipsManager}>
      {children}
      <PortalsContainer ref={setPortalsContainer} />
    </TooltipsManagerContext.Provider>
  );
}

export function useTooltipsManager(): TooltipsManager {
  const tooltipsManager = useContext(TooltipsManagerContext);
  if (!tooltipsManager) {
    throw new Error('TooltipsManager is not provided via context.');
  }
  return tooltipsManager;
}

const TooltipsManagerContext = createContext<TooltipsManager | null>(null);

class VisibilityManager {
  private closeById = new Map<string, () => void>();

  registerCloseTooltip(id: string, close: () => void): void {
    this.closeById.set(id, close);
  }

  deregisterCloseTooltip(id: string): void {
    this.closeById.delete(id);
  }

  closeAll(): void {
    for (const [, close] of this.closeById) {
      close();
    }
    this.closeById.clear();
  }
}

const PortalsContainer = forwardRef<HTMLDivElement>(
  function PortalsContainer(_, ref) {
    return (
      <div id="tooltips-container" className={s.PortalsContainer} ref={ref} />
    );
  },
);
