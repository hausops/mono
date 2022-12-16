import {Button} from '@/volto/Button';
import {useReducer} from 'react';
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
  const [{units}, dispatch] = useReducer(reducer, initialState);
  return {
    units,
    addUnit(unit) {
      dispatch({type: 'INSERT', index: units.length, unit});
    },
    insertUnitAfter(i, unit) {
      dispatch({type: 'INSERT', index: i + 1, unit});
    },
    updateUnit(index, unit) {
      dispatch({type: 'UPDATE', index, unit});
    },
    removeUnit(index) {
      dispatch({type: 'REMOVE', index});
    },
  };
}

type State = {
  nextId: number;
  units: RentalUnit[];
};

type Action =
  | {type: 'INSERT'; index: number; unit: RentalUnit}
  | {type: 'UPDATE'; index: number; unit: RentalUnit}
  | {type: 'REMOVE'; index: number};

function reducer(state: State, action: Action): State {
  const id = state.nextId;
  let units: RentalUnit[];

  switch (action.type) {
    case 'INSERT':
      const newUnit = {...action.unit, id: `${id}`};
      units = [...state.units];
      units.splice(action.index, 0, newUnit);
      return {
        nextId: id + 1,
        units,
      };

    case 'UPDATE':
      units = [...state.units];
      units[action.index] = action.unit;
      return {...state, units};

    case 'REMOVE':
      units = [...state.units];
      units.splice(action.index, 1);
      return {...state, units};

    default:
      return state;
  }
}

const initialState: State = {
  // Keep track of unique ids instead of using array index as key
  // to avoid reconciliation bug when items are added or removed
  // in the middle of the array.
  //
  // This cannot be randomized (i.e. using nanoid) because of SSR.
  nextId: 1,
  // Starts with one unit with no data.
  units: [{...emptyUnit(), id: '0'}],
};

function emptyUnit(): RentalUnit {
  return {id: '', number: '', size: '', rentAmount: ''};
}
