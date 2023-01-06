import {LeaseModel} from './LeaseModel';
import {LeaseService} from './LeaseService';

export class LocalLeaseService implements LeaseService {
  private readonly byUnitId = new Map<string, LeaseModel>(
    DEMO_LEASES.map((lease) => [lease.unitId, lease])
  );

  async getByUnitId(unitId: string) {
    return this.byUnitId.get(unitId);
  }
}

const DEMO_LEASES: LeaseModel[] = [
  {
    id: '8563400',
    unitId: '1029599-0', // see LocalPropertyService
    tenant: {
      id: 'u0001',
      firstName: 'Jane',
      lastName: 'Doe',
      imageUrl: '/images/michael-dam-mEZ3PoFGs_k-unsplash-avatar.jpg',
      email: 'jane.doe@example.com',
      phone: '(555) 123-4567',
    },
    upcomingRent: {
      dueDate: '2023-02-01T00:00:00',
      status: 'outstanding',
      amount: 4000,
      paid: 3000,
      paymentPending: 1000,
    },
    pastPayments: [
      {
        dueDate: '2023-01-01T00:00:00',
        status: 'overdue',
        amount: 4000,
        paid: 0,
      },
      {
        dueDate: '2022-12-01T00:00:00',
        status: 'fully-paid',
        amount: 4000,
        paid: 4000,
      },
      {
        dueDate: '2022-11-01T00:00:00',
        status: 'fully-paid',
        amount: 4000,
        paid: 4000,
      },
    ],
  },
];
