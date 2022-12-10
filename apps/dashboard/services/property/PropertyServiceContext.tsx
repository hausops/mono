import {createContext, PropsWithChildren, useContext} from 'react';
import {PropertyService} from './PropertyService';

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
    throw new Error('PropertyService is not provided from the context.');
  }
  return svc;
}

const PropertyServiceContext = createContext<PropertyService | null>(null);
