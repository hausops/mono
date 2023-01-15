import {
  createContext,
  forwardRef,
  PropsWithChildren,
  useContext,
  useRef,
  useState,
} from 'react';
import * as s from './Tooltip.css';

type TooltipsManager = {
  portalsContainer: HTMLElement | null;
  visibility: VisibilityManager;
};

type TooltipsManagerProviderProps = PropsWithChildren<{}>;

export function TooltipsManagerProvider(props: TooltipsManagerProviderProps) {
  const {children} = props;
  // Cannot use useRef here because it does not cause re-render
  // thus the context is not updated after the container mounted.
  const [portalsContainer, setPortalsContainer] = useState<HTMLElement | null>(
    null
  );

  const tooltipsManager: TooltipsManager = {
    portalsContainer,
    visibility: useVisibilityManager(),
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

interface VisibilityManager {
  addCloseFunction(id: string, close: () => void): void;
  removeCloseFunction(id: string): void;
  closeAll(): void;
}

function useVisibilityManager(): VisibilityManager {
  const {current: closeById} = useRef(new Map<string, () => void>());
  return {
    addCloseFunction(id, close) {
      closeById.set(id, close);
    },

    removeCloseFunction(id) {
      closeById.delete(id);
    },

    closeAll() {
      for (const [, close] of closeById) {
        close();
      }
      closeById.clear();
    },
  };
}

const PortalsContainer = forwardRef<HTMLDivElement>(function PortalsContainer(
  _,
  ref
) {
  return (
    <div id="tooltips-container" className={s.PortalsContainer} ref={ref} />
  );
});
