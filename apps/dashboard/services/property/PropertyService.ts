import {PropertyData, PropertyModel} from './PropertyModel';

export interface PropertyService {
  getAll(): Promise<PropertyModel[]>;

  // creates a new property from newPropertyData.
  // The service will assign the property id.
  add(newPropertyData: PropertyData): Promise<PropertyModel>;

  // deletes the property matching id.
  // If success, resolves to the id of the property being deleted
  // If a property with the id is not found, it will reject with an error.
  delete(id: string): Promise<string>;
}

export class PropertyNotFoundErr extends Error {
  constructor(readonly id: string) {
    super(`Cannot find property with id=${id}.`);
  }
}
