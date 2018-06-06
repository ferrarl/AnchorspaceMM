import React, { Component } from 'react'
import Page from '../pages/Index.jsx';

class Hello extends Component {
  render () {
    return (
    <div className="postCssTester">
    <Page />
      <h1 className="postCssTester__title">Hello <span className="postCssTester__title-react">React</span>!</h1>
      <div className="colorMaps">
        <div className="colorMaps__container">
          <div className="colorMaps__item colorMaps__item--color blue-dark"></div>
          <div className="colorMaps__item colorMaps__item--color blue-medium"></div>
          <div className="colorMaps__item colorMaps__item--color blue"></div>
          <div className="colorMaps__item colorMaps__item--color blue-light"></div>
          <div className="colorMaps__item colorMaps__item--color blue-lighter"></div>
        </div>
      </div>
    </div>
    )
  }
}

export default Hello
