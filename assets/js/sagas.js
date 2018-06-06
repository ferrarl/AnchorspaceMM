import { call, takeEvery } from 'redux-saga/effects';
import { LOCATION_CHANGE } from 'react-router-redux';

const example = injections =>
  function* exampleSaga(action) {
    yield call([injections.logger, 'log'], 'Location changed. Action: ', action);
  };

export default function root(injections) {
  return function* rootSaga() {
    yield takeEvery(LOCATION_CHANGE, example(injections));
  };
}
