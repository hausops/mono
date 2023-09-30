import {Button} from '@/volto/Button';
import type {ComponentPropsWithoutRef, ReactElement} from 'react';
import * as s from './ButtonGroup.css';

type ButtonGroupProps = {
  children?: ButtonGroupChildren;
};

type ButtonProps = ComponentPropsWithoutRef<typeof Button>;
type ButtonGroupChildren =
  | null
  | ReactElement<ButtonProps>
  | ReactElement<ButtonProps>[];

export function ButtonGroup(props: ButtonGroupProps) {
  return <div className={s.ButtonGroup}>{props.children}</div>;
}
