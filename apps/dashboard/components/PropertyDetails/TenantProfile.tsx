import {Avatar} from '@/volto/Avatar';
import {Card} from '@/volto/Card';
import {Check} from '@/volto/icons';
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

  const [isCopied, setCopied] = useState(false);

  async function handleClick() {
    if (navigator) {
      await navigator.clipboard.writeText(children);
      setCopied(true);
    }
  }

  function handleTooltipClose() {
    tooltipState.close();
    setCopied(false);
  }

  return (
    <>
      <button
        aria-describedby={tooltipState.isOpen ? tooltipId : undefined}
        className={s.Contact}
        onClick={handleClick}
        onFocus={tooltipState.open}
        onBlur={handleTooltipClose}
        onMouseEnter={tooltipState.open}
        onMouseLeave={handleTooltipClose}
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
          <div className={s.ContactTooltipContent}>
            {isCopied ? (
              <CopiedContact timeoutMs={1500} onTimeout={handleTooltipClose} />
            ) : (
              'Copy to clipboard'
            )}
          </div>
        </Card>
      </Tooltip>
    </>
  );
}

function CopiedContact({
  timeoutMs,
  onTimeout,
}: {
  timeoutMs: number;
  onTimeout: () => void;
}) {
  useEffect(() => {
    if (timeoutMs < 0) {
      return;
    }
    const timer = setTimeout(onTimeout, timeoutMs);
    return () => clearTimeout(timer);
  }, [timeoutMs, onTimeout]);

  return (
    <div className={s.CopiedContact}>
      <span className={s.CopiedContactIcon}>
        <Check />
      </span>
      Copied
    </div>
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
  const {overlayProps, updatePosition} = useOverlayPosition({
    boundaryElement: portalsContainer || undefined,
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
