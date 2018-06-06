import { createStore, compose, applyMiddleware } from 'redux';
import { routerMiddleware } from 'react-router-redux';
import createSagaMiddleware from 'redux-saga';
import history from 'history/createHashHistory';

import injections from './injections';
import rootReducer from './reducers';
import rootSaga from './sagas';

const sagaMiddleware = createSagaMiddleware();

const configureStore = (preloadedState = null) => {
  const hasWindow = typeof window !== 'undefined';
  /* eslint-disable no-underscore-dangle */
  const devToolsCompose = hasWindow && window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__;
  /* eslint-enable */
  const composeEnhancers = devToolsCompose || compose;

  const middleware = composeEnhancers(applyMiddleware(sagaMiddleware, routerMiddleware(history)));

  let store;

  if (preloadedState) {
    store = createStore(rootReducer, preloadedState, middleware);
  } else {
    store = createStore(rootReducer, middleware);
  }

  sagaMiddleware.run(rootSaga(injections));

  if (module.hot) {
    module.hot.accept('./reducers', changed => {
      window.console.debug(`[hmr:store.js] accepting reducer changes made in ${changed}`);
      const nextRootReducer = require('./reducers').default;
      store.replaceReducer(nextRootReducer);
    });
  }

  return store;
};

export default configureStore;
