import {Reducer, useCallback, useReducer} from 'react';

export type FieldsState<T> = {
  readonly fields: T;
  updateField<K extends keyof T>(key: K, value: T[K]): void;
};

export function useFieldsState<T>(initialFields: T): FieldsState<T> {
  const [fields, dispatch] = useReducer<Reducer<T, Action<T>>>(
    fieldsReducer,
    initialFields,
  );

  return {
    fields,
    updateField: useCallback((key, value) => {
      dispatch({key, value});
    }, []),
  };
}

type Action<T> = {
  key: keyof T;
  value: T[keyof T];
};

function fieldsReducer<T>(fields: T, {key, value}: Action<T>): T {
  return {...fields, [key]: value};
}
