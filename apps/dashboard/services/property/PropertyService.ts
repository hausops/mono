import {nanoid} from 'nanoid';
import {PropertyData, PropertyModel} from './PropertyModel';

export class PropertyService {
  private readonly properties: Map<string, PropertyModel> = new Map();

  createProperty(newPropertyData: PropertyData): void {
    const id = nanoid();
    const property = {...newPropertyData, id};
    this.properties.set(id, property);
  }

  deleteProperty(id: string): void {
    this.properties.delete(id);
  }
}
