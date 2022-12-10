export interface AddressService {
  getAllStates(): StateEntry[];
}

type StateEntry = {
  code: string;
  name: string;
};
