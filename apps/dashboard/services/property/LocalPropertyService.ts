import {nanoid} from 'nanoid';
import {PropertyData, PropertyModel} from './PropertyModel';
import {PropertyNotFoundErr, PropertyService} from './PropertyService';

export class LocalPropertyService implements PropertyService {
  private readonly properties: Map<string, PropertyModel> = new Map(
    Object.entries(DEMO_PROPERTIES)
  );

  async getAll(): Promise<PropertyModel[]> {
    return [...this.properties.values()];
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

const DEMO_PROPERTIES: {[id: string]: PropertyModel} = {
  '1029599': {
    id: '1029599',
    type: 'single-family',
    coverImageUrl: '/images/pexels-scott-webb-1029599.jpg',
    address: {
      line1: '527 Bridle Street',
      city: 'Flowery Branch',
      state: 'GA',
      zip: '30542',
    },
  },
  '2724749': {
    id: '2724749',
    type: 'single-family',
    coverImageUrl: '/images/pexels-mark-mccammon-2724749.jpg',
    address: {
      line1: '495 Ohio Street',
      city: 'Harleysville',
      state: 'PA',
      zip: '19438',
    },
  },
  '3288102': {
    id: '3288102',
    type: 'single-family',
    coverImageUrl: '/images/pexels-curtis-adams-3288102.jpg',
    address: {
      line1: '9026 Washington Dr.',
      city: 'Orland Park',
      state: 'IL',
      zip: '60462',
    },
  },
  '9999990': {
    id: '9999990',
    type: 'single-family',
    address: {
      line1: '9189 South Argyle Dr.',
      city: 'Natchez',
      state: 'MS',
      zip: '39120',
    },
  },
  '9999991': {
    id: '9999991',
    type: 'single-family',
    address: {
      line1: '9189 South Argyle Dr.',
      city: 'Natchez',
      state: 'MS',
      zip: '39120',
    },
  },
  '9999992': {
    id: '9999992',
    type: 'single-family',
    address: {
      line1: '9189 South Argyle Dr.',
      city: 'Natchez',
      state: 'MS',
      zip: '39120',
    },
  },
};
