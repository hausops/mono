import {LocalLeaseService} from './LocalLeaseService';

// Temporary: local-only until we have a way to configure and
// inject different versions of the LeaseService.
export const leaseSvc = new LocalLeaseService();

export * from './LeaseModel';
export * from './LeaseService';
// export * from './LeaseServiceContext';
export * from './LocalLeaseService';
