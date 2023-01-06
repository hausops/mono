export type LeaseModel = {
  id: string;
  unitId: string; // FK - a property unit ID
  tenant: Tenant;
  upcomingRent?: RentPayment;
  pastPayments: RentPayment[];
};

type Tenant = {
  id: string;
  imageUrl: string;
  firstName: string;
  lastName: string;
  email: string;
  phone?: string;
};

type AmountUSD = number; // just for documentation

export type RentPayment = {
  dueDate: string;
  amount: AmountUSD;
  status: 'outstanding' | 'overdue' | 'fully-paid';
  paid: AmountUSD;
  paymentPending?: AmountUSD;
};
