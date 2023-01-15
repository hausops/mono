import {
  createContext,
  forwardRef,
  PropsWithChildren,
  useContext,
  useState,
} from 'react';
import * as s from './Tooltip.css';

type TooltipsManager = {
  portalsContainer: HTMLElement | null;
};

export const TooltipsManagerContext = createContext<TooltipsManager>({
  portalsContainer: null,
});

type TooltipsManagerProviderProps = PropsWithChildren<{}>;

export function TooltipsManagerProvider(props: TooltipsManagerProviderProps) {
  const {children} = props;
  // Cannot use useRef here because it does not cause re-render
  // thus the context is not updated after the container mounted.
  const [portalsContainer, setPortalsContainer] = useState<HTMLElement | null>(
    null
  );

  const tooltipsManager = {
    portalsContainer,
  };

  return (
    <TooltipsManagerContext.Provider value={tooltipsManager}>
      {children}
      <PortalsContainer ref={setPortalsContainer} />
    </TooltipsManagerContext.Provider>
  );
}

const PortalsContainer = forwardRef<HTMLDivElement>(function PortalsContainer(
  _,
  ref
) {
  return (
    <div id="tooltips-container" className={s.PortalsContainer} ref={ref} />
  );
});

export function useTooltipsManager() {
  const tooltipsManager = useContext(TooltipsManagerContext);
  return tooltipsManager;
}
