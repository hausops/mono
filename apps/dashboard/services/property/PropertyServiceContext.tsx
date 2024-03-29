import {createContext, useContext, type PropsWithChildren} from 'react';
import type {PropertyService} from './PropertyService';

export function PropertyServiceProvider({
  service,
  children,
}: PropsWithChildren<{service: PropertyService}>) {
  return (
    <PropertyServiceContext.Provider value={service}>
      {children}
    </PropertyServiceContext.Provider>
  );
}

export function usePropertyService(): PropertyService {
  const svc = useContext(PropertyServiceContext);
  if (!svc) {
    throw new Error('PropertyService is not provided in the context.');
  }
  return svc;
}

const PropertyServiceContext = createContext<PropertyService | null>(null);
