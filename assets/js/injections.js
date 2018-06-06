import browserHistory from 'history/createBrowserHistory';
import hashHistory from 'history/createHashHistory';
import memoryHistory from 'history/createMemoryHistory';

let createHistory;

switch ((process.env.REACT_APP_ROUTER_HISTORY || '').toLowerCase()) {
  case 'browser':
    createHistory = browserHistory;
    break;
  case 'hash':
    createHistory = hashHistory;
    break;
  case 'memory':
    createHistory = memoryHistory;
    break;
  default:
    createHistory = browserHistory;
}

export default {
  createHistory,
  logger: console,
};
