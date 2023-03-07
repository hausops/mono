import {PropertyType} from '@/services/property';
import {Radio, RadioGroupState, useRadio, useRadioGroup} from '@/volto/Radio';
import {useFocusRing} from '@react-aria/focus';
import clsx from 'clsx';
import {useId, useRef} from 'react';
import * as s from './PropertyTypeSelector.css';

type PropertyTypeSelectorProps = {
  state: RadioGroupState<PropertyType>;
};

export function PropertyTypeSelector({state}: PropertyTypeSelectorProps) {
  const {labelProps, radioGroupProps, radioProps} = useRadioGroup({
    name: 'property-type',
    orientation: 'horizontal',
  });

  return (
    <div {...radioGroupProps}>
      <h3 {...labelProps} className={s.Label}>
        Property type
      </h3>
      <div className={s.Options}>
        <PropertyTypeOption
          {...radioProps}
          value="single-family"
          label="Single-family property"
          description="A property with one renter associates with one address such as a house or a townhouse."
          state={state}
        />
        <PropertyTypeOption
          {...radioProps}
          value="multi-family"
          label="Multi-family property"
          description="A building with multiple rental units such as a duplex, an apartment, or a condo."
          state={state}
        />
      </div>
    </div>
  );
}

type PropertyTypeOptionProps = {
  name: string;
  value: PropertyType;
  label: string;
  description: string;
  state: RadioGroupState<PropertyType>;
};

function PropertyTypeOption({state, ...props}: PropertyTypeOptionProps) {
  const titleId = useId();
  const descriptionId = useId();

  const ref = useRef<HTMLInputElement>(null);
  const {clickableZoneProps, inputProps, isSelected} = useRadio(
    props,
    state,
    ref
  );
  const {isFocusVisible, focusProps} = useFocusRing();

  return (
    <div
      {...clickableZoneProps}
      className={clsx(s.Option, {
        [s.OptionState.selected]: isSelected,
        [s.OptionState.focusVisible]: isFocusVisible,
      })}
    >
      <header className={s.OptionHeader}>
        <Radio
          {...inputProps}
          {...focusProps}
          aria-labelledby={titleId}
          aria-describedby={descriptionId}
          ref={ref}
        />
        <h4 id={titleId} className={s.OptionTitle}>
          {props.label}
        </h4>
      </header>
      <p id={descriptionId} className={s.OptionDescription}>
        {props.description}
      </p>
    </div>
  );
}
