import {noop} from 'lodash';
import {createContext, useContext, useEffect, useMemo} from 'react';
import {useAppSelector} from 'redux/hooks';

import {useTestRun} from 'providers/TestRun/TestRun.provider';
import TestSpecsSelectors from 'selectors/TestSpecs.selectors';
import AssertionResults, {TAssertionResultEntry} from 'models/AssertionResults.model';
import {TTestSpecEntry} from 'models/TestSpecs.model';
import {ISuggestion} from 'types/TestSpecs.types';
import useTestSpecsCrud from './hooks/useTestSpecsCrud';
import {useTest} from '../Test/Test.provider';

interface IContext {
  revert: (originalSelector: string) => void;
  add(spec: TTestSpecEntry): void;
  update(selector: string, spec: TTestSpecEntry): void;
  remove(selector: string): void;
  publish(): void;
  cancel(): void;
  dryRun(definitionList: TTestSpecEntry[]): void;
  updateIsInitialized(isInitialized: boolean): void;
  selectedTestSpec?: TAssertionResultEntry;
  assertionResults?: AssertionResults;
  specs: TTestSpecEntry[];
  isLoading: boolean;
  isError: boolean;
  isDraftMode: boolean;
  setSelectedSpec(selector?: string): void;
  setSelectorSuggestions(selectorSuggestions: ISuggestion[]): void;
  setPrevSelector(selector: string): void;
}

export const Context = createContext<IContext>({
  add: noop,
  revert: () => noop,
  update: noop,
  remove: noop,
  publish: noop,
  dryRun: noop,
  cancel: noop,
  isLoading: false,
  isError: false,
  isDraftMode: false,
  specs: [],
  setSelectedSpec: noop,
  updateIsInitialized: noop,
  setSelectorSuggestions: noop,
  setPrevSelector: noop,
});

interface IProps {
  testId: string;
  runId: string;
  children: React.ReactNode;
}

export const useTestSpecs = () => useContext(Context);

const TestSpecsProvider = ({children, testId, runId}: IProps) => {
  const {test} = useTest();
  const {run} = useTestRun();

  const assertionResults = useAppSelector(state => TestSpecsSelectors.selectAssertionResults(state));
  const specs = useAppSelector(state => TestSpecsSelectors.selectSpecs(state));
  const isDraftMode = useAppSelector(state => TestSpecsSelectors.selectIsDraftMode(state));
  const isLoading = useAppSelector(state => TestSpecsSelectors.selectIsLoading(state));
  const isInitialized = useAppSelector(state => TestSpecsSelectors.selectIsInitialized(state));

  const selectedSpec = useAppSelector(TestSpecsSelectors.selectSelectedSpec);
  const selectedTestSpec = useAppSelector(state => TestSpecsSelectors.selectAssertionBySelector(state, selectedSpec!));

  const {
    add,
    cancel,
    publish,
    remove,
    dryRun,
    update,
    init,
    reset,
    revert,
    updateIsInitialized,
    setSelectedSpec,
    setSelectorSuggestions,
    setPrevSelector,
  } = useTestSpecsCrud({
    assertionResults,
    test,
    testId,
    runId,
    isDraftMode,
  });

  useEffect(() => {
    if (run.state === 'FINISHED') init(run.result);
  }, [init, run.result, run.state]);

  useEffect(() => {
    return () => {
      reset();
    };
  }, [reset]);

  useEffect(() => {
    if (isInitialized && run.state === 'FINISHED') dryRun(specs);
  }, [dryRun, specs, isInitialized, run.state]);

  const value = useMemo<IContext>(
    () => ({
      add,
      remove,
      update,
      isLoading,
      isError: false,
      isDraftMode,
      publish,
      dryRun,
      assertionResults,
      specs,
      selectedTestSpec,
      cancel,
      setSelectedSpec,
      revert,
      updateIsInitialized,
      setSelectorSuggestions,
      setPrevSelector,
    }),
    [
      add,
      remove,
      update,
      isLoading,
      isDraftMode,
      publish,
      dryRun,
      assertionResults,
      specs,
      selectedTestSpec,
      cancel,
      setSelectedSpec,
      revert,
      updateIsInitialized,
      setSelectorSuggestions,
      setPrevSelector,
    ]
  );

  return <Context.Provider value={value}>{children}</Context.Provider>;
};

export default TestSpecsProvider;
