import React, { Component } from 'react';
import { Provider } from 'react-redux';
import { Redirect, Route, Switch } from 'react-router-dom';
import { ConnectedRouter } from 'react-router-redux';

import IndexPage from './pages/Index.jsx';
import ListingsPage from './pages/Listings.jsx';
import FourOhFourPage from './pages/404.jsx';

import Header from './components/Header.jsx';
import Cards from './components/Cards/Cards.js';
import JumbotronFluid from './components/Jumbotron/JumbotronFluid.js'

import configureStore from './store';
import injections from './injections';
import { setGlobalStyles, Root } from './styles';

const store = configureStore();
const history = injections.createHistory();

class App extends Component {
  componentWillMount() {
    setGlobalStyles();
  }

  render() {
    return (
      <Provider store={store}>
        <Root>
          <ConnectedRouter history={history}>
            <Switch>
              <Route exact path="/" component={IndexPage} />
              <Route path="/listings" component={ListingsPage} />
              <Route component={FourOhFourPage} />
            </Switch>
          </ConnectedRouter>
        </Root>
      </Provider>
    );
  }
}

export default App;
