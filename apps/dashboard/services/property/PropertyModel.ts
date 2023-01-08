import {AddressModel} from '@/services/address';

export type PropertyType = 'single-family' | 'multi-family';
export type PropertyModel = SingleFamilyProperty | MultiFamilyProperty;

export type SingleFamilyProperty = {
  id: string;
  type: 'single-family';
  coverImageUrl?: string;
  address: AddressModel;
  builtYear?: number;
  unit: RentalUnit;
};

export type MultiFamilyProperty = {
  id: string;
  type: 'multi-family';
  coverImageUrl?: string;
  address: AddressModel;
  builtYear?: number;
  units: RentalUnit[];
};

export type RentalUnit = {
  id: string;
  number?: string;
  bedrooms?: number;
  bathrooms?: number;
  size?: number; // in sq.m.
  rentAmount?: number;
  activeListing?: {
    id: string;
  };
};

export type NewPropertyData =
  | {
      type: 'single-family';
      coverImageUrl?: string;
      address: AddressModel;
      builtYear?: number;
      unit: NewRentalUnit;
    }
  | {
      type: 'multi-family';
      coverImageUrl?: string;
      address: AddressModel;
      builtYear?: number;
      units: NewRentalUnit[];
    };

type NewRentalUnit = {
  number?: string;
  bedrooms?: number;
  bathrooms?: number;
  size?: number; // in sq.m.
  rentAmount?: number;
};
