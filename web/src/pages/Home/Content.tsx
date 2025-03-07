import ConfigCTA from './ConfigCTA';
import Resources from './Resources';

interface IProps {
  isLoading: boolean;
  shouldDisplayConfigSetup: boolean;
  skipConfigSetup(): void;
}

const Content = ({isLoading, shouldDisplayConfigSetup, skipConfigSetup}: IProps) => {
  if (isLoading) return null;

  return shouldDisplayConfigSetup ? <ConfigCTA onSkip={skipConfigSetup} /> : <Resources />;
};

export default Content;
