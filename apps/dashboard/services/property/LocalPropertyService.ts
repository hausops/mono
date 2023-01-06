import {nanoid} from 'nanoid';
import {NewPropertyData, PropertyModel} from './PropertyModel';
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

  async add(newPropertyData: NewPropertyData): Promise<PropertyModel> {
    const id = nanoid();
    const p = withId(newPropertyData);
    const property: PropertyModel =
      p.type === 'single-family'
        ? {...p, unit: withId(p.unit)}
        : {...p, units: p.units.map(withId)};
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

function withId<T>(o: T): T & {id: string} {
  return {...o, id: nanoid()};
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
    unit: {
      id: '1029599-0',
      bedrooms: 3,
      bathrooms: 2.5,
      size: 1024,
    },
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
    unit: {
      id: '2724749-0',
      activeListing: {id: '8d5f11ed'},
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
    unit: {id: '3288102-0'},
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
    unit: {id: '9999990'},
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
    unit: {id: '9999991'},
  },
  {
    id: '9999992',
    type: 'single-family',
    address: {
      line1: '290 County Rd',
      line2: '#2011',
      city: 'Vista',
      state: 'CA',
      zip: '92081',
    },
    unit: {id: '9999992'},
  },
];
