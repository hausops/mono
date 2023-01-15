import {Avatar} from '@/volto/Avatar';
import {Card} from '@/volto/Card';
import {useOverlayPosition, AriaPositionProps} from '@react-aria/overlays';
import Link from 'next/link';
import {PropsWithChildren, useRef, useState} from 'react';
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
  const tooltipState = useTooltipState();
  const tooltipTriggerRef = useRef<HTMLButtonElement>(null);
  return (
    <>
      <button
        className={s.Contact}
        onMouseEnter={tooltipState.open}
        onMouseLeave={tooltipState.close}
        ref={tooltipTriggerRef}
      >
        {children}
      </button>
      <Tooltip
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
  isOpen: boolean;
}> &
  Pick<AriaPositionProps, 'boundaryElement' | 'placement' | 'targetRef'>;

// Tooltip.Overlay
function Tooltip(props: TooltipProps) {
  const {boundaryElement, children, isOpen, placement, targetRef} = props;

  const overlayRef = useRef<HTMLDivElement>(null);
  const {overlayProps} = useOverlayPosition({
    // boundaryElement,
    isOpen,
    overlayRef,
    placement,
    targetRef,
  });

  if (!isOpen) {
    return null;
  }

  const position = overlayProps.style ?? {};
  return createPortal(
    <div
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
    document.body
  );
}

function useTooltipState(initialOpen: boolean = false) {
  const [isOpen, setOpen] = useState(initialOpen);
  return {
    isOpen,
    open() {
      setOpen(true);
    },
    close() {
      setOpen(false);
    },
    toggle() {
      setOpen(!isOpen);
    },
  };
}
