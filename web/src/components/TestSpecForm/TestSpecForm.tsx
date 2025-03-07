import {Button, Form, Tag} from 'antd';
import {useEffect, useState} from 'react';

import {SELECTOR_LANGUAGE_CHEAT_SHEET_URL} from 'constants/Common.constants';
import {CompareOperator} from 'constants/Operator.constants';
import {useAppSelector} from 'redux/hooks';
import AssertionSelectors from 'selectors/Assertion.selectors';
import SpanSelectors from 'selectors/Span.selectors';
import TestSpecsSelectors from 'selectors/TestSpecs.selectors';
import OperatorService from 'services/Operator.service';
import {TStructuredAssertion} from 'types/Assertion.types';
import {singularOrPlural} from 'utils/Common';
import AssertionCheckList from './AssertionCheckList';
import useOnFieldsChange from './hooks/useOnFieldsChange';
import useOnValuesChange from './hooks/useOnValuesChange';
import SelectorInput from './SelectorInput';
import SelectorSuggestions from './SelectorSuggestions';
import * as S from './TestSpecForm.styled';

export interface IValues {
  assertions?: TStructuredAssertion[];
  selector?: string;
}

interface IProps {
  defaultValues?: IValues;
  isEditing?: boolean;
  onCancel(): void;
  onClearSelectorSuggestions(): void;
  onClickPrevSelector(prevSelector: string): void;
  onSubmit(values: IValues): void;
  runId: string;
  testId: string;
}

const TestSpecForm = ({
  defaultValues: {
    assertions = [
      {
        left: '',
        comparator: OperatorService.getOperatorSymbol(CompareOperator.EQUALS),
        right: '',
      },
    ],
    selector = '',
  } = {},
  isEditing = false,
  onCancel,
  onClearSelectorSuggestions,
  onClickPrevSelector,
  onSubmit,
  runId,
  testId,
}: IProps) => {
  const [form] = Form.useForm<IValues>();
  const [isValid, setIsValid] = useState(false);

  const spanIdList = useAppSelector(SpanSelectors.selectMatchedSpans);
  const attributeList = useAppSelector(state =>
    AssertionSelectors.selectAttributeList(state, testId, runId, spanIdList)
  );

  const onValuesChange = useOnValuesChange({setIsValid});
  const onFieldsChange = useOnFieldsChange();

  useEffect(() => {
    onValuesChange(null, {assertions, selector});
  }, []);

  const selectorSuggestions = useAppSelector(TestSpecsSelectors.selectSelectorSuggestions);
  const prevSelector = useAppSelector(TestSpecsSelectors.selectPrevSelector);

  return (
    <S.AssertionForm>
      <S.AssertionFormHeader>
        <S.AssertionFormTitle>{isEditing ? 'Edit Test Spec' : 'Add Test Spec'}</S.AssertionFormTitle>
      </S.AssertionFormHeader>

      <Form<IValues>
        name="assertion-form"
        form={form}
        initialValues={{
          remember: true,
          assertions,
          selector,
        }}
        onFinish={onSubmit}
        autoComplete="off"
        layout="vertical"
        data-cy="assertion-form"
        onFieldsChange={onFieldsChange}
        onValuesChange={onValuesChange}
      >
        <S.FormSection>
          <S.FormSectionHeaderSelector>
            <S.FormSectionRow1>
              <S.FormSectionTitle $noMargin>1. Select Spans</S.FormSectionTitle>
              <Tag color="blue">{`${spanIdList.length} ${singularOrPlural('span', spanIdList.length)} selected`}</Tag>
            </S.FormSectionRow1>
            <a href={SELECTOR_LANGUAGE_CHEAT_SHEET_URL} target="_blank">
              <S.ReadIcon /> SL cheat sheet
            </a>
          </S.FormSectionHeaderSelector>
          <S.FormSectionRow>
            <S.FormSectionText>Select the spans to which a set of assertions will be applied</S.FormSectionText>
          </S.FormSectionRow>
          <SelectorInput form={form} testId={testId} runId={runId} onValidSelector={setIsValid} />

          <S.SuggestionsContainer>
            {!isEditing && (
              <SelectorSuggestions
                onClick={query => {
                  onClickPrevSelector(form.getFieldValue('selector'));
                  onClearSelectorSuggestions();
                  form.setFieldsValue({
                    selector: query,
                  });
                }}
                onClickPrevSelector={query => {
                  onClickPrevSelector('');
                  onClearSelectorSuggestions();
                  form.setFieldsValue({
                    selector: query,
                  });
                }}
                prevSelector={prevSelector}
                selectorSuggestions={selectorSuggestions}
              />
            )}
          </S.SuggestionsContainer>
        </S.FormSection>

        <S.FormSection>
          <S.FormSectionTitle>2. Add Assertions</S.FormSectionTitle>
          <S.FormSectionRow>
            <S.FormSectionText>Add assertions using the attributes from the selected spans</S.FormSectionText>
          </S.FormSectionRow>
          <Form.List name="assertions">
            {(fields, {add, remove}) => (
              <AssertionCheckList
                form={form}
                fields={fields}
                add={add}
                remove={remove}
                attributeList={attributeList}
                runId={runId}
                testId={testId}
                spanIdList={spanIdList}
              />
            )}
          </Form.List>
        </S.FormSection>

        <S.AssertionFromActions>
          <Button onClick={onCancel}>Cancel</Button>
          <Button type="primary" disabled={!isValid} onClick={form.submit} data-cy="assertion-form-submit-button">
            Save Test Spec
          </Button>
        </S.AssertionFromActions>
      </Form>
    </S.AssertionForm>
  );
};

export default TestSpecForm;
