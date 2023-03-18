import {AddressModel} from '@/services/address';

export type PropertyType = 'single-family' | 'multi-family';
export type PropertyModel = SingleFamily.Property | MultiFamily.Property;

export type NewPropertyData =
  | SingleFamily.NewPropertyData
  | MultiFamily.NewPropertyData;

export namespace SingleFamily {
  export type Property = {
    id: string;
    type: 'single-family';
    coverImageUrl?: string;
    address: AddressModel;
    yearBuilt?: number;
    unit: Unit;
  };

  export type Unit = {
    id: string;
    bedrooms?: number;
    bathrooms?: number;
    size?: number; // in sq.m.
    rentAmount?: number;
    activeListing?: {
      id: string;
    };
  };

  export type NewPropertyData = {
    type: 'single-family';
    coverImageUrl?: string;
    address: AddressModel;
    yearBuilt?: number;
    unit: {
      bedrooms?: number;
      bathrooms?: number;
      size?: number; // in sq.m.
      rentAmount?: number;
    };
  };
}

export namespace MultiFamily {
  export type Property = {
    id: string;
    type: 'multi-family';
    coverImageUrl?: string;
    address: AddressModel;
    yearBuilt?: number;
    units: Unit[];
  };

  export type Unit = {
    id: string;
    number: string; // unit.number; should be a string
    bedrooms?: number;
    bathrooms?: number;
    size?: number; // in sq.m.
    rentAmount?: number;
    activeListing?: {
      id: string;
    };
  };

  export type NewPropertyData = {
    type: 'multi-family';
    coverImageUrl?: string;
    address: AddressModel;
    yearBuilt?: number;
    units: Array<{
      number: string;
      bedrooms?: number;
      bathrooms?: number;
      size?: number; // in sq.m.
      rentAmount?: number;
    }>;
  };
}
