import React from 'react';
import ReactDOM from 'react-dom';

import App from './App';

jest.mock('./injections');

it('renders without crashing', () => {
  const div = document.createElement('div');
  ReactDOM.render(<App />, div);
});
