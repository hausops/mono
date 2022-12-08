import Button from '@/volto/Button';
import {useRef, useState} from 'react';
import * as s from './MultiFamilyForm.css';
import {RentalUnit} from './RentalUnit';
import {UnitEntry} from './UnitEntry';

export type MultiFamilyFormState = {
  units: RentalUnit[];

  addUnit(unit: RentalUnit): void;
  insertUnitAfter(index: number, unit: RentalUnit): void;
  updateUnit(index: number, unit: RentalUnit): void;
  removeUnit(index: number): void;
};

export function MultiFamilyForm({state}: {state: MultiFamilyFormState}) {
  return (
    <div className={s.MultiFamilyForm}>
      <h3 className={s.Title}>Units</h3>
      <ul className={s.UnitEntries}>
        {state.units.map((unit, i) => (
          <UnitEntry
            key={unit.id}
            index={i}
            state={unit}
            onChange={(unit) => state.updateUnit(i, unit)}
            onClone={() => {
              const unitCopy = {...state.units[i]};
              state.insertUnitAfter(i, unitCopy);
            }}
            onRemove={
              // do not allow removing the last unit (ensure at least one)
              state.units.length > 1 ? () => state.removeUnit(i) : undefined
            }
          />
        ))}
      </ul>
      <footer>
        <Button variant="text" onClick={() => state.addUnit(emptyUnit())}>
          + Add unit
        </Button>
      </footer>
    </div>
  );
}

export function useMultiFamilyFormState(): MultiFamilyFormState {
  // Starts with one unit with no data.
  const [units, setUnits] = useState<RentalUnit[]>([{...emptyUnit(), id: '0'}]);

  // Keep track of unique ids instead of using array index as key
  // to avoid reconciliation bug when items are added or removed
  // in the middle of the array.
  //
  // This cannot be randomized (i.e. using nanoid) because of SSR.
  const nextIdRef = useRef(1);
  const nextId = (): string => `${nextIdRef.current++}`;

  return {
    units,
    addUnit(unit) {
      unit = {...unit, id: nextId()};
      setUnits([...units, unit]);
    },
    insertUnitAfter(i, unit) {
      const next = [...units];
      unit = {...unit, id: nextId()};
      next.splice(i + 1, 0, unit);
      setUnits(next);
    },
    updateUnit(i, unit) {
      const next = [...units];
      next[i] = unit;
      setUnits(next);
    },
    removeUnit(i) {
      const next = [...units];
      next.splice(i, 1);
      setUnits(next);
    },
  };
}

function emptyUnit(): RentalUnit {
  return {id: '', number: '', size: '', rentAmount: ''};
}
