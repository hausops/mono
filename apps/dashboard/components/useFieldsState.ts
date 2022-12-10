import {useState} from 'react';

export type FieldsState<T> = T & {
  updateField: <K extends keyof T>(key: K, value: T[K]) => void;
};

export function useFieldsState<T>(initialState: T): FieldsState<T> {
  const [state, setState] = useState<T>(initialState);
  return {
    ...state,
    updateField(key, value) {
      setState({...state, [key]: value});
    },
  };
}
