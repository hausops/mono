import {LocalAddressService} from './LocalAddressService';

// Temporary: local-only until we have a way to configure and
// inject different versions of the AddressService.
export const addressSvc = new LocalAddressService();

export * from './Address';
export * from './AddressModel';
export * from './AddressService';
// export * from './AddressServiceContext';
export * from './LocalAddressService';
