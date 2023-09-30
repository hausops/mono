/**
 * CopyToClipboardTooltip is a specialized, styled Tooltip to use with
 * buttons or links to provide copy-to-clipboard functionality.
 */

import {Card} from '@/volto/Card';
import {CheckIcon} from '@/volto/icons';
import {useEffect, useState} from 'react';
import * as s from './CopyToClipboardTooltip.css';
import {Tooltip, TooltipProps} from './Tooltip';
import {TooltipState} from './useTooltipState';

type CopyToClipboardTooltipProps = {
  isCopied: boolean;
  onClose: () => void;
  state: TooltipState;
} & Pick<TooltipProps, 'id' | 'placement' | 'targetRef'>;

export function CopyToClipboardTooltip(props: CopyToClipboardTooltipProps) {
  const {isCopied, onClose, state, ...tooltipProps} = props;
  return (
    <Tooltip {...tooltipProps} isOpen={state.isOpen}>
      <Card>
        <div className={s.Content}>
          {isCopied ? (
            <Copied timeoutMs={1500} onTimeout={onClose} />
          ) : (
            'Copy to clipboard'
          )}
        </div>
      </Card>
    </Tooltip>
  );
}

function Copied({
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
    <div className={s.Copied}>
      <span className={s.CopiedIcon}>
        <CheckIcon />
      </span>
      Copied
    </div>
  );
}

type CopyState = {
  isCopied: boolean;
  copyToClipboard(): Promise<void>;
  handleTooltipClose(): void;
};

export function useCopyToClipboardState(
  content: string,
  tooltipState: TooltipState,
): CopyState {
  const [isCopied, setCopied] = useState(false);
  return {
    isCopied,

    async copyToClipboard() {
      if (navigator) {
        await navigator.clipboard.writeText(content);
        setCopied(true);
      }
    },

    handleTooltipClose() {
      tooltipState.close();
      setCopied(false);
    },
  };
}
