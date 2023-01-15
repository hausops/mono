import {Avatar} from '@/volto/Avatar';
import {Card} from '@/volto/Card';
import {useTooltipsManager} from '@/volto/Tooltip';
import {AriaPositionProps, useOverlayPosition} from '@react-aria/overlays';
import Link from 'next/link';
import {PropsWithChildren, useEffect, useId, useRef, useState} from 'react';
import {createPortal} from 'react-dom';
import * as s from './TenantProfile.css';

type TenantProfileProps = {
  name: string;
  imageUrl: string;
  email: string;
  phone?: string;
};

export function TenantProfile({
  name,
  imageUrl,
  email,
  phone,
}: TenantProfileProps) {
  return (
    <article className={s.TenantProfile}>
      <Link className={s.Avatar} href="#">
        <Avatar name={name} imageUrl={imageUrl} />
      </Link>
      <div className={s.Details}>
        <Link className={s.Name} href="#">
          {name}
        </Link>
        <div className={s.ContactInfo}>
          <Contact>{email}</Contact>
          {phone && (
            <>
              <i className={s.ContactSeparator} />
              <Contact>{phone}</Contact>
            </>
          )}
        </div>
      </div>
    </article>
  );
}

function Contact({children}: {children: string}) {
  const tooltipId = useId();
  const tooltipState = useTooltipState(tooltipId);
  const tooltipTriggerRef = useRef<HTMLButtonElement>(null);
  return (
    <>
      <button
        aria-describedby={tooltipState.isOpen ? tooltipId : undefined}
        className={s.Contact}
        onFocus={tooltipState.open}
        onBlur={tooltipState.close}
        onMouseEnter={tooltipState.open}
        onMouseLeave={tooltipState.close}
        ref={tooltipTriggerRef}
      >
        {children}
      </button>
      <Tooltip
        id={tooltipId}
        isOpen={tooltipState.isOpen}
        placement="bottom"
        targetRef={tooltipTriggerRef}
      >
        <Card>
          <p className={s.ContactTooltipContent}>Copy to clipboard</p>
        </Card>
      </Tooltip>
    </>
  );
}

// Tooltip is an element to provide a short context or hint about an UI element.

type TooltipProps = PropsWithChildren<{
  id: string;
  isOpen: boolean;
}> &
  Pick<AriaPositionProps, 'placement' | 'targetRef'>;

// Tooltip.Overlay
function Tooltip(props: TooltipProps) {
  const {children, id, isOpen, placement, targetRef} = props;

  const {portalsContainer} = useTooltipsManager();
  const overlayRef = useRef<HTMLDivElement>(null);
  const {overlayProps} = useOverlayPosition({
    boundaryElement: portalsContainer || undefined,
    isOpen,
    overlayRef,
    placement,
    targetRef,
  });

  if (!isOpen || !portalsContainer) {
    return null;
  }

  const position = overlayProps.style ?? {};
  return createPortal(
    <div
      id={id}
      className={s.TooltipContainer}
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
    portalsContainer
  );
}

function useTooltipState(id: string, initialOpen: boolean = false) {
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
