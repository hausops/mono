import type {AddressModel} from './AddressModel';

export class Address {
  static from(model: AddressModel): Address {
    return new Address(model);
  }

  constructor(private readonly model: AddressModel) {}

  format(): [string, string] {
    const {line1, line2, city, state, zip} = this.model;
    const street = [line1, line2].filter((s) => !!s).join(' ');
    const region = `${city}, ${state} ${zip}`;
    return [street, region];
  }

  toString(): string {
    const [street, region] = this.format();
    return `${street}, ${region}`;
  }
}
