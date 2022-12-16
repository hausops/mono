import {Close, Copy} from '@/volto/icons';
import {Select, toOption} from '@/volto/Select';
import {TextField} from '@/volto/TextField';
import {RentalUnit} from './RentalUnit';
import * as s from './UnitEntry.css';

type UnitEntryProps = {
  index: number;
  state: RentalUnit;
  onChange(state: RentalUnit): void;
  onClone(index: number): void;
  onRemove?(index: number): void;
};

// TODO: refactor
const bedsOptions = [
  {label: 'Studio', value: 0},
  ...[1, 2, 3, 4, 5].map(toOption),
];

// TODO: refactor
const bathsOptions = [
  {label: 'None', value: 0},
  ...[1, 1.5, 2, 2.5, 3, 3.5, 4].map(toOption),
];

export function UnitEntry({
  index,
  state,
  onChange,
  onClone,
  onRemove,
}: UnitEntryProps) {
  return (
    <li>
      <header className={s.Header}>
        <h3 className={s.Title}>Unit {index + 1}</h3>
        <div className={s.Actions}>
          <button className={s.ActionButton} onClick={() => onClone(index)}>
            <span className={s.ActionButtonIcon}>
              <Copy />
            </span>
            Clone
          </button>
          {onRemove && (
            <button className={s.ActionButton} onClick={() => onRemove(index)}>
              <span className={s.ActionButtonIcon}>
                <Close />
              </span>
              Remove
            </button>
          )}
        </div>
      </header>
      <div className={s.Form}>
        <TextField
          label="Unit #"
          name="number"
          value={state.number}
          onChange={(e) => onChange({...state, number: e.target.value})}
        />
        <Select
          label="Beds"
          name="bedrooms"
          options={bedsOptions}
          value={state.bedrooms}
          onChange={(e) => onChange({...state, bedrooms: +e.target.value})}
        />
        <Select
          label="Baths"
          name="bathrooms"
          options={bathsOptions}
          value={state.bathrooms}
          onChange={(e) => onChange({...state, bathrooms: +e.target.value})}
        />
        <TextField
          type="number"
          label="Size"
          name="size"
          value={state.size}
          onChange={(e) => onChange({...state, size: e.target.value})}
        />
        <TextField
          type="number"
          label="Rent"
          name="rentAmount"
          value={state.rentAmount}
          onChange={(e) => onChange({...state, rentAmount: e.target.value})}
        />
      </div>
    </li>
  );
}
