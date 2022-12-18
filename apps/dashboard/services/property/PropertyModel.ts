import {AddressModel} from '@/services/address';

export type PropertyModel = PropertyData & {id: string};

export type PropertyType = 'single-family' | 'multi-family';
export type PropertyData = SingleFamilyProperty | MultiFamilyProperty;

export type SingleFamilyProperty = {
  type: 'single-family';
  coverImageUrl?: string;
  address: AddressModel;
  builtYear?: number;
  // ...
  bedrooms?: number;
  bathrooms?: number;
  size?: number; // in sq.m.
  rentAmount?: number;
};

export type MultiFamilyProperty = {
  type: 'multi-family';
  coverImageUrl?: string;
  address: AddressModel;
  builtYear?: number;
  //
  units: RentalUnit[];
};

export type RentalUnit = {
  number?: string;
  bedrooms?: number;
  bathrooms?: number;
  size?: number; // in sq.m.
  rentAmount?: number;
};
