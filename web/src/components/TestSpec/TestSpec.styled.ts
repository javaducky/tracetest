import {ApartmentOutlined} from '@ant-design/icons';
import {Typography} from 'antd';
import styled from 'styled-components';

export const ActionsContainer = styled.div`
  align-items: center;
  display: flex;
  flex: 1;
  justify-content: flex-end;
`;

export const Container = styled.div<{$isDeleted: boolean}>`
  align-items: center;
  background-color: ${({theme}) => theme.color.white};
  border: ${({theme}) => `1px solid ${theme.color.border}`};
  cursor: pointer;
  display: flex;
  gap: 12px;
  padding: 16px;

  > div:first-child {
    opacity: ${({$isDeleted}) => ($isDeleted ? 0.5 : 1)};
  }
`;

export const Column = styled.div`
  display: flex;
  flex-direction: column;
`;

export const HeaderDetail = styled(Typography.Text)`
  color: ${({theme}) => theme.color.text};
  font-size: ${({theme}) => theme.size.sm};
  margin-right: 8px;
`;

export const HeaderDot = styled.span<{$passed: boolean}>`
  background-color: ${({$passed, theme}) => ($passed ? theme.color.success : theme.color.error)};
  height: 10px;
  width: 10px;
  display: inline-block;
  margin-right: 4px;
  line-height: 0;
  vertical-align: -0.1em;
  border-radius: 50%;
`;

export const HeaderSpansIcon = styled(ApartmentOutlined)`
  margin-right: 4px;
`;

export const Selector = styled.div`
  display: flex;
  flex-direction: column;
`;
