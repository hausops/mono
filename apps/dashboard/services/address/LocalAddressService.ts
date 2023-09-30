import type {AddressService} from './AddressService';
import data from './data.local.json';

// LocalAddressService keep data on the machine
export class LocalAddressService implements AddressService {
  getAllStates() {
    return data.states;
  }
}
