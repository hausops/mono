import {useOverlayPosition, type AriaPositionProps} from '@react-aria/overlays';
import {useEffect, useRef, type PropsWithChildren} from 'react';
import {createPortal} from 'react-dom';
import * as s from './Tooltip.css';
import {useTooltipsManager} from './TooltipsManager';

export type TooltipProps = PropsWithChildren<{
  id: string;
  isOpen: boolean;
}> &
  Pick<AriaPositionProps, 'placement' | 'targetRef'>;

export function Tooltip(props: TooltipProps) {
  const {children, id, isOpen, placement = 'top', targetRef} = props;

  const {portalsContainer} = useTooltipsManager();
  const overlayRef = useRef<HTMLDivElement>(null);
  const {overlayProps, updatePosition} = useOverlayPosition({
    boundaryElement: portalsContainer || undefined, // satisfy the type
    isOpen,
    overlayRef,
    placement,
    targetRef,
  });

  // update position when content (children) changes
  useEffect(() => {
    if (isOpen) {
      updatePosition();
    }
  }, [children, isOpen, updatePosition]);

  if (!isOpen || !portalsContainer) {
    return null;
  }

  const position = overlayProps.style ?? {};
  return createPortal(
    <div
      id={id}
      className={s.Tooltip}
      ref={overlayRef}
      role="tooltip"
      style={{
        top: position.top,
        left: position.left,
        bottom: position.bottom,
        right: position.right,
      }}
    >
      {children}
    </div>,
    portalsContainer,
  );
}
