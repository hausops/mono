import {LeaseModel} from './LeaseModel';

export interface LeaseService {
  getByUnitId(unitId: string): Promise<LeaseModel | undefined>;
}
