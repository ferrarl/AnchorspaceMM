import React from 'react';
import PropTypes from 'prop-types';

export default class Panel extends React.Component {
  render() {
    const cls = this.props.alert ? 'card bg-danger' : 'card bg-info'
    return (
      <div className={cls}>
        <div className="card-body">
          <h1>{this.props.title}</h1>
          {this.props.children}
        </div>
      </div>
    )
  }
}

