import {Avatar} from '@/volto/Avatar';
import {Card} from '@/volto/Card';
import {Check} from '@/volto/icons';
import {Tooltip, useTooltip} from '@/volto/Tooltip';
import Link from 'next/link';
import {useEffect, useRef, useState} from 'react';
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
  const tooltip = useTooltip();
  const tooltipTriggerRef = useRef<HTMLButtonElement>(null);

  const [isCopied, setCopied] = useState(false);

  async function handleClick() {
    if (navigator) {
      await navigator.clipboard.writeText(children);
      setCopied(true);
    }
  }

  function handleTooltipClose() {
    tooltip.state.close();
    setCopied(false);
  }

  return (
    <>
      <button
        aria-describedby={tooltip.state.isOpen ? tooltip.id : undefined}
        className={s.Contact}
        onClick={handleClick}
        onFocus={tooltip.state.open}
        onBlur={handleTooltipClose}
        onMouseEnter={tooltip.state.open}
        onMouseLeave={handleTooltipClose}
        ref={tooltipTriggerRef}
      >
        {children}
      </button>
      <Tooltip
        id={tooltip.id}
        isOpen={tooltip.state.isOpen}
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
