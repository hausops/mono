import {createContext, useContext, type PropsWithChildren} from 'react';
import type {AddressService} from './AddressService';

export function AddressServiceProvider({
  service,
  children,
}: PropsWithChildren<{service: AddressService}>) {
  return (
    <AddressServiceContext.Provider value={service}>
      {children}
    </AddressServiceContext.Provider>
  );
}

export function useAddressService(): AddressService {
  return useContext(AddressServiceContext);
}

const emptyAddressService: AddressService = {
  getAllStates() {
    return [];
  },
};

const AddressServiceContext = createContext(emptyAddressService);
