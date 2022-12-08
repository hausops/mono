import {RentalUnit} from '@/services/property';
import Button from '@/volto/Button';
import {useState} from 'react';
import * as s from './MultiFamilyForm.css';
import {UnitEntry} from './UnitEntry';

export type MultiFamilyFormState = {
  units: RentalUnit[];

  addUnit(unit: RentalUnit): void;
  insertUnitAfter(index: number, unit: RentalUnit): void;
  updateUnit(index: number, unit: RentalUnit): void;
  removeUnit(index: number): void;
};

// Bug 1: React doesn't update the DOM for UnitEntry correctly
//        when removing items due to using index as key

// Bug 2: numeric TextField won't remove 0

export function MultiFamilyForm({state}: {state: MultiFamilyFormState}) {
  return (
    <div className={s.MultiFamilyForm}>
      <h3 className={s.Title}>Units</h3>
      <ul className={s.UnitEntries}>
        {state.units.map((unit, i) => (
          <UnitEntry
            key={i}
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
        <Button variant="text" onClick={() => state.addUnit({})}>
          + Add unit
        </Button>
      </footer>
    </div>
  );
}

// starts with one unit with no data
const initialUnits: RentalUnit[] = [{}];

export function useMultiFamilyFormState(): MultiFamilyFormState {
  const [units, setUnits] = useState(initialUnits);
  return {
    units,
    addUnit(unit) {
      setUnits([...units, unit]);
    },
    insertUnitAfter(i, unit) {
      const next = [...units];
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
