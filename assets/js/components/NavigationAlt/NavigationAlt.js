import React from 'react';
import Navigation from '../Navigation/Navigation.js';

class NavigationAlt extends React.Component {
  render() {
    return (
      <div>
      <nav class="navbar navbar-expand-sm w-100 navbar-light bg-light py-4">
	<a className="navbar-brand" href="/"><img src='/assets/images/logo-alt.png' alt="Logo"/></a>
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbar5">
        <span class="navbar-toggler-icon"></span>
    </button>
    <div class="navbar-collapse collapse" id="navbar5">
        <ul class="navbar-nav ml-auto">
            <li class="nav-item active">
                <a class="nav-link" href="/">Space</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="/">Dashboard</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="/">Blog</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="/">Contact</a>
            </li>
        </ul>
    </div>
</nav>
</div>
    )
  }
}

export default NavigationAlt;
