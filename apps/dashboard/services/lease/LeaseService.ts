import type {LeaseModel} from './LeaseModel';

export interface LeaseService {
  getByUnitId(unitId: string): Promise<LeaseModel | undefined>;
  getManyByUnitIds(unitIds: string[]): Promise<LeaseModel[]>;
}
