import React from 'react';
import ReactDOM from 'react-dom';

import App from './App';
import injections from './injections';
import registerServiceWorker from './registerServiceWorker';

injections.logger.info(`App build: ${process.env.REACT_APP_BUILD_VERSION}`);

const rootEl = document.getElementById('root');

ReactDOM.render(<App />, rootEl);

if (module.hot) {
  module.hot.accept('./App', () => {
    const NextApp = require('./App').default;
    ReactDOM.render(<NextApp />, rootEl);
  });
}

// see caveats at `src/registerServiceWorker.js`
registerServiceWorker();
