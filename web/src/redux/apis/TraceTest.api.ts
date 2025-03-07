import {createApi, fetchBaseQuery} from '@reduxjs/toolkit/query/react';
import {TracetestApiTags} from 'constants/Test.constants';
import {TTestApiEndpointBuilder} from 'types/Test.types';
import ConfigEndpoint from './endpoints/Config.endpoint';
import EnvironmentEndpoint from './endpoints/Environment.endpoint';
import ExpressionEndpoint from './endpoints/Expression.endpoint';
import ResourceEndpoint from './endpoints/Resource.endpoint';
import TestEndpoint from './endpoints/Test.endpoint';
import TestRunEndpoint from './endpoints/TestRun.endpoints';
import TransactionEndpoint from './endpoints/Transaction.endpoint';
import TransactionRunEndpoint from './endpoints/TransactionRun.endpoint';

const PATH = `${document.baseURI}api/`;

const TraceTestAPI = createApi({
  reducerPath: 'tests',
  baseQuery: fetchBaseQuery({
    baseUrl: PATH,
  }),
  tagTypes: Object.values(TracetestApiTags),
  endpoints(builder: TTestApiEndpointBuilder) {
    return {
      ...TransactionEndpoint(builder),
      ...TransactionRunEndpoint(builder),
      ...TestRunEndpoint(builder),
      ...TestEndpoint(builder),
      ...EnvironmentEndpoint(builder),
      ...ExpressionEndpoint(builder),
      ...ResourceEndpoint(builder),
      ...ConfigEndpoint(builder),
    };
  },
});

export const {
  useCreateTestMutation,
  useGetTestByIdQuery,
  useGetTestVersionByIdQuery,
  useGetTestListQuery,
  useRunTestMutation,
  useDeleteTestByIdMutation,
  useGetRunByIdQuery,
  useGetRunListQuery,
  useGetSelectedSpansQuery,
  useLazyGetSelectedSpansQuery,
  useReRunMutation,
  useLazyGetRunListQuery,
  useDryRunMutation,
  useDeleteRunByIdMutation,
  useGetJUnitByRunIdQuery,
  useLazyGetJUnitByRunIdQuery,
  useEditTestMutation,
  useGetEnvironmentsQuery,
  useCreateEnvironmentMutation,
  useUpdateEnvironmentMutation,
  useDeleteEnvironmentMutation,
  useCreateTransactionMutation,
  useGetTransactionByIdQuery,
  useDeleteTransactionByIdMutation,
  useEditTransactionMutation,
  useGetTransactionRunsQuery,
  useLazyGetTransactionRunsQuery,
  useGetTransactionRunByIdQuery,
  useDeleteTransactionRunByIdMutation,
  useParseExpressionMutation,
  useGetResourcesQuery,
  useRunTransactionMutation,
  useGetTransactionVersionByIdQuery,
  useGetResourceDefinitionQuery,
  useLazyGetResourceDefinitionQuery,
  useGetDataStoresQuery,
  useCreateDataStoreMutation,
  useUpdateDataStoreMutation,
  useDeleteDataStoreMutation,
  useTestConnectionMutation,
  useLazyGetTransactionVersionByIdQuery,
} = TraceTestAPI;
export const {endpoints} = TraceTestAPI;

export default TraceTestAPI;
