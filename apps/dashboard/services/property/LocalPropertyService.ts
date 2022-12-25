import {nanoid} from 'nanoid';
import {PropertyData, PropertyModel} from './PropertyModel';
import {PropertyNotFoundErr, PropertyService} from './PropertyService';

export class LocalPropertyService implements PropertyService {
  private readonly properties: Map<string, PropertyModel> = new Map(
    DEMO_PROPERTIES.map((p) => [p.id, p])
  );

  async getAll(): Promise<PropertyModel[]> {
    return [...this.properties.values()];
  }

  async get(id: string): Promise<PropertyModel | undefined> {
    return this.properties.get(id);
  }

  async add(newPropertyData: PropertyData): Promise<PropertyModel> {
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

const DEMO_PROPERTIES: PropertyModel[] = [
  {
    id: '1029599',
    type: 'single-family',
    coverImageUrl: '/images/pexels-scott-webb-1029599.jpg',
    address: {
      line1: '527 Bridle Street',
      city: 'Flowery Branch',
      state: 'GA',
      zip: '30542',
    },
    bedrooms: 3,
    bathrooms: 2.5,
    size: 1024,
  },
  {
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
  {
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
  {
    id: '9999990',
    type: 'single-family',
    address: {
      line1: '9189 South Argyle Dr.',
      city: 'Natchez',
      state: 'MS',
      zip: '39120',
    },
  },
  {
    id: '9999991',
    type: 'single-family',
    address: {
      line1: '9189 South Argyle Dr.',
      city: 'Natchez',
      state: 'MS',
      zip: '39120',
    },
  },
  {
    id: '9999992',
    type: 'single-family',
    address: {
      line1: '9189 South Argyle Dr.',
      city: 'Natchez',
      state: 'MS',
      zip: '39120',
    },
  },
];
