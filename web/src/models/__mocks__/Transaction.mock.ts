import faker from '@faker-js/faker';
import Transaction, {TRawTransaction} from 'models/Transaction.model';
import {IMockFactory} from 'types/Common.types';
import TestMock from './Test.mock';

const TransactionMock: IMockFactory<Transaction, TRawTransaction> = () => ({
  raw(data = {}) {
    const test = TestMock.raw();
    const test2 = TestMock.raw();
    const test3 = TestMock.raw();

    return {
      id: faker.datatype.uuid(),
      name: faker.name.firstName(),
      version: faker.datatype.number(),
      description: faker.company.catchPhraseDescriptor(),
      createdAt: faker.date.past().toISOString(),
      steps: [test, test2, test3],
      env: {
        usename: 'john doe',
      },
      summary: {
        runs: 0,
        hasRuns: false,
        lastRun: {
          time: '',
          passes: 0,
          fails: 0,
        },
      },
      ...data,
    };
  },
  model(data = {}) {
    return Transaction(this.raw(data));
  },
});

export default TransactionMock();
