import Logo from 'assets/Logo.svg';
import {FC} from 'react';
import {Link, useLocation} from 'react-router-dom';
import HomeAnalyticsService from '../../services/Analytics/HomeAnalytics.service';
import * as S from './Header.styled';
import {HeaderMenu} from './HeaderMenu';

const {onGuidedTourClick} = HomeAnalyticsService;

export const testRunPageRegex = /\/test\/(.*)\/run\/(.*)/;
const Header: FC = () => {
  const {pathname} = useLocation();
  return pathname.match(testRunPageRegex) ? null : (
    <S.Header>
      <Link to="/">
        <img alt="tracetest_log" data-cy="logo" src={Logo} />
      </Link>
      <HeaderMenu onGuidedTourClick={onGuidedTourClick} pathname={pathname} />
    </S.Header>
  );
};

export default Header;
