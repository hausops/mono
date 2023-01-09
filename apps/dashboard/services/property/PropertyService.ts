import {NewPropertyData, PropertyModel} from './PropertyModel';

export interface PropertyService {
  getAll(): Promise<PropertyModel[]>;

  get(id: string): Promise<PropertyModel | undefined>;

  // creates a new property from newPropertyData.
  // The service will assign the property id.
  add(newPropertyData: NewPropertyData): Promise<PropertyModel>;

  // patches the property matching id with updatePropertyData.
  // If success, resolves to the updated PropertyModel.
  // If a property with the id is not found, it will reject with an error.
  update<T extends PropertyModel>(
    id: string,
    updateProperty: Partial<T>
  ): Promise<T>;

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
