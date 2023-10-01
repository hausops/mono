import {LocalPropertyService} from './LocalPropertyService';

// Temporary: local-only until we have a way to configure and
// inject different versions of the PropertyService.
export const propertySvc = new LocalPropertyService();

export * from './LocalPropertyService';
export * from './PropertyModel';
export * from './PropertyService';
// export * from './PropertyServiceContext';
