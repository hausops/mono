import clsx from 'clsx';
import {ComponentPropsWithRef, forwardRef} from 'react';
import * as s from './Radio.css';

type RadioProps = ComponentPropsWithRef<'input'>;

// Radio is a custom styled radio input to render consistently in all browsers.
export const Radio = forwardRef<HTMLInputElement, RadioProps>(
  function Radio(props, ref) {
    return (
      <input
        {...props}
        type="radio"
        className={clsx(s.Radio, props.className)}
        ref={ref}
      />
    );
  }
);
