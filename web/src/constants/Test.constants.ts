export interface IKeyValue {
  value: string;
  key: string;
}

export const DEFAULT_HEADERS: IKeyValue[] = [{key: 'Content-Type', value: 'application/json'}];

export enum TriggerTypes {
  http = 'http',
  grpc = 'grpc',
  traceid = 'traceid',
}

export enum SortBy {
  Created = 'created',
  LastRun = 'last_run',
  Name = 'name',
}

export enum SortDirection {
  Asc = 'asc',
  Desc = 'desc',
}

export const sortOptions = [
  {
    label: 'Recently Created',
    value: 'created',
    params: {sortDirection: SortDirection.Desc, sortBy: SortBy.Created},
  },
  {
    label: 'Last Run',
    value: 'last_run',
    params: {sortDirection: SortDirection.Desc, sortBy: SortBy.LastRun},
  },
  {
    label: 'Name, A to Z',
    value: 'name_asc',
    params: {sortDirection: SortDirection.Asc, sortBy: SortBy.Name},
  },
  {
    label: 'Name, Z to A',
    value: 'name_desc',
    params: {sortDirection: SortDirection.Desc, sortBy: SortBy.Name},
  },
] as const;

export enum TracetestApiTags {
  ENVIRONMENT = 'environment',
  TRANSACTION = 'transaction',
  TRANSACTION_RUN = 'transactionRun',
  TEST = 'test',
  TEST_DEFINITION = 'testDefinition',
  TEST_RUN = 'testRun',
  SPAN = 'span',
  EXPRESSION = 'expression',
  RESOURCE = 'resource',
  CONFIG = 'config',
}
