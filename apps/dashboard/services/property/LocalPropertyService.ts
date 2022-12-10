import {nanoid} from 'nanoid';
import {PropertyData, PropertyModel} from './PropertyModel';
import {PropertyNotFoundErr, PropertyService} from './PropertyService';

export class LocalPropertyService implements PropertyService {
  private readonly properties: Map<string, PropertyModel> = new Map();

  async getAll(): Promise<PropertyModel[]> {
    return Object.values(this.properties);
  }

  async create(newPropertyData: PropertyData): Promise<PropertyModel> {
    const id = nanoid();
    const property = {...newPropertyData, id};
    this.properties.set(id, property);
    return property;
  }

  async delete(id: string): Promise<string> {
    const existed = this.properties.delete(id);
    if (!existed) {
      throw new PropertyNotFoundErr(id);
    }
    return id;
  }
}
