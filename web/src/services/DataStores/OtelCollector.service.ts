import {SupportedDataStores, TDataStoreService} from 'types/Config.types';

const OtelCollectorService = (): TDataStoreService => ({
  getRequest(draft, dataStoreType = SupportedDataStores.OtelCollector) {
    return Promise.resolve({
      type: dataStoreType,
      name: dataStoreType,
    });
  },
  validateDraft() {
    return Promise.resolve(true);
  },
  getInitialValues(draft, dataStoreType = SupportedDataStores.OtelCollector) {
    return {
      dataStore: {
        name: dataStoreType,
        type: dataStoreType,
      },
      dataStoreType,
    };
  },
});

export default OtelCollectorService();
