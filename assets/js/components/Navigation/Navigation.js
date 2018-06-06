import React from 'react';
import { Link } from 'react-router-dom';

class Navigation extends React.Component {
  render() {
    return (
      <div>
        <nav className="navbar navbar-expand-sm w-100">
          <a className="navbar-brand" href="/"><img src="/assets/images/logo.png" alt="Logo"/></a>
            <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbar5">
                <span className="navbar-toggler-icon"></span>
            </button>
            <div className="navbar-collapse collapse" id="navbar5">
                <ul className="navbar-nav ml-auto">
                    <li className="nav-item active">
                        <a className="nav-link" href="/listings">List Your Space</a>
                    </li>
                    <li className="nav-item">
                        <a className="nav-link" href="/users/new/">Sign Up</a>
                    </li>
                    <li className="nav-item">
                      <a className="nav-link" href="/signin">Sign In</a>
                    </li>
                    <li className="nav-item">
                        <a className="nav-link" href="/">Blog</a>
                    </li>
                    <li className="nav-item">
                        <a className="nav-link" href="/">Contact</a>
                    </li>
                </ul>
            </div>
        </nav>
      </div>
    )
  }
}

export default Navigation;
