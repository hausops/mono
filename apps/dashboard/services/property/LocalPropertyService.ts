import {nanoid} from 'nanoid';
import type {NewPropertyData, PropertyModel} from './PropertyModel';
import {PropertyNotFoundErr, type PropertyService} from './PropertyService';

export class LocalPropertyService implements PropertyService {
  private readonly properties: Map<string, PropertyModel> = new Map(
    DEMO_PROPERTIES.map((p) => [p.id, p]),
  );

  async getAll(): Promise<PropertyModel[]> {
    return [...this.properties.values()];
  }

  async getById(id: string): Promise<PropertyModel | undefined> {
    return this.properties.get(id);
  }

  async add(newPropertyData: NewPropertyData): Promise<PropertyModel> {
    const p = withId(newPropertyData);
    const property: PropertyModel =
      p.type === 'single-family'
        ? {...p, unit: withId(p.unit)}
        : {...p, units: p.units.map(withId)};
    this.properties.set(p.id, property);
    return property;
  }

  async update<T extends PropertyModel>(
    id: string,
    updateProperty: Partial<T>,
  ): Promise<T> {
    const previous = this.properties.get(id);

    if (!previous) {
      throw new PropertyNotFoundErr(id);
    }

    if (updateProperty.type && updateProperty.type !== previous.type) {
      throw new Error(
        `Mismatched property type: was=${previous.type}, got=${updateProperty.type}.`,
      );
    }

    // TODO: figure out how to handle type of previous safer in Typescript
    const updated = {...(previous as T), ...updateProperty};
    this.properties.set(id, updated);
    return updated;
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
    id: '4375210',
    type: 'multi-family',
    coverImageUrl: '/images/pexels-quintin-gellar-612949.jpg',
    address: {
      line1: '10 Rosa Street',
      city: 'San Francisco',
      state: 'CA',
      zip: '94107',
    },
    units: [
      {
        id: '4375210-1',
        number: '201',
        bedrooms: 0,
        bathrooms: 1,
        size: 524,
        rentAmount: 2075,
        activeListing: {id: '8d5f11ed'},
      },
      {
        id: '4375210-2',
        number: '301',
        bedrooms: 2,
        bathrooms: 2,
        size: 950,
        rentAmount: 3850,
      },
      {
        id: '4375210-3',
        number: '302',
        bedrooms: 2,
        bathrooms: 2,
        size: 982,
        rentAmount: 4000,
      },
      {
        id: '4375210-4',
        number: '303',
        bedrooms: 2,
        bathrooms: 2,
        size: 982,
        rentAmount: 4000,
      },
    ],
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
