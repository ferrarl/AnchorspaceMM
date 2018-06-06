import React from 'react';
import PropTypes from 'prop-types';

import Navigation from './Navigation/Navigation.js';
import Jumbotron from './Jumbotron/';
import { Header, Title } from './styles';

/** Header is an example component from create-react-app. */
const Component = ({ logo, title }) => (
  <Header>
    <Navigation/>
  </Header>
);

Component.displayName = 'Header';

Component.propTypes = {
  /** Path to the logo image (probably resolved with Webpack). */
  logo: PropTypes.string,
  /** Short and important text to display. */
  title: PropTypes.string,
};

Component.defaultProps = {
  title: 'Welcome to React',
};

export default Component;
