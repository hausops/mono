import {PropsWithChildren} from 'react';

export function SvgIcon({children}: PropsWithChildren<{}>) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 24 24"
      focusable={false}
      aria-hidden={true}
    >
      {children}
    </svg>
  );
}
