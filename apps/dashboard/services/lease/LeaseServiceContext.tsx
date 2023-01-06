import {createContext, PropsWithChildren, useContext} from 'react';
import {LeaseService} from './LeaseService';

export function LeaseServiceProvider({
  service,
  children,
}: PropsWithChildren<{service: LeaseService}>) {
  return (
    <LeaseServiceContext.Provider value={service}>
      {children}
    </LeaseServiceContext.Provider>
  );
}

export function useLeaseService(): LeaseService {
  const svc = useContext(LeaseServiceContext);
  if (!svc) {
    throw new Error('LeaseService is not provided from the context.');
  }
  return svc;
}

const LeaseServiceContext = createContext<LeaseService | null>(null);
