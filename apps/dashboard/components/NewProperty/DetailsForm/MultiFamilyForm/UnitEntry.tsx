import {BedroomsSelect, BathroomsSelect} from '@/components/PropertyForm';
import {MiniTextButton} from '@/volto/Button';
import {Close, Copy} from '@/volto/icons';
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
          <MiniTextButton icon={<Copy />} onClick={() => onClone(index)}>
            Clone
          </MiniTextButton>
          {onRemove && (
            <MiniTextButton icon={<Close />} onClick={() => onRemove(index)}>
              Remove
            </MiniTextButton>
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
        <BedroomsSelect
          name="bedrooms"
          value={state.bedrooms}
          onChange={(e) => onChange({...state, bedrooms: +e.target.value})}
        />
        <BathroomsSelect
          name="bathrooms"
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
