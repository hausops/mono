import {Avatar} from '@/volto/Avatar';
import {
  CopyToClipboardTooltip,
  useCopyToClipboardState,
  useTooltip,
} from '@/volto/Tooltip';
import Link from 'next/link';
import {useRef} from 'react';
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
  const {isCopied, copyToClipboard, handleTooltipClose} =
    useCopyToClipboardState(children, tooltip.state);

  return (
    <>
      <button
        aria-describedby={tooltip.state.isOpen ? tooltip.id : undefined}
        className={s.Contact}
        onClick={copyToClipboard}
        onFocus={tooltip.state.open}
        onBlur={handleTooltipClose}
        onMouseEnter={tooltip.state.open}
        onMouseLeave={handleTooltipClose}
        ref={tooltipTriggerRef}
      >
        {children}
      </button>
      <CopyToClipboardTooltip
        id={tooltip.id}
        isCopied={isCopied}
        onClose={handleTooltipClose}
        placement="bottom"
        state={tooltip.state}
        targetRef={tooltipTriggerRef}
      />
    </>
  );
}
