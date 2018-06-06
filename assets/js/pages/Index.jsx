import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import Header from '../components/Header.jsx';
import Jumbotron from '../components/Jumbotron';
import Cards from '../components/Cards/Cards.js';
import JumbotronFluid from '../components/Jumbotron/JumbotronFluid.js';
import Footer from '../components/Footer.jsx';
import { PageWrapper } from '../styles';

import { FetchListings, FetchNeighborhoods } from "../utils/networkHelpers.js"

export default class Page extends Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: true,
      listings: [],
      loadingNeighborhoods: true,
      neighborhoods: []
    };
  }

  componentDidMount () {
    this.loadNeighborhoods()
    this.loadListings()
  }

  loadNeighborhoods () {
    this.setState({
      loadingNeighborhoods: true
    })

    FetchNeighborhoods().then(json => {
      this.setState({
        loadingNeighborhoods: false,
        neighborhoods: json
      })
    })
  }

  sortedNeighborhoods() {
    var neighborhoods = this.state.neighborhoods
    neighborhoods.sort((a, b) => {
      if (a.name < b.name) {
        return -1
      } else if (a.name > b.name) {
        return 1
      }
      return 0
    })
    return neighborhoods
  }

  loadListings() {
    this.setState({
      loading: true
    });

    FetchListings().then(json => {
      this.setState({
        loading: false,
        listings: json
      });
    });
  }

  render() {
    return (
      <PageWrapper>
        <Header />
        <Jumbotron neighborhoods={this.sortedNeighborhoods()} />
        {this.state.loading ? (
          <Loading />
        ) : (
          <Cards listings={this.state.listings} numberOfRows={2} />
        )}
        <JumbotronFluid />
        <Footer />
      </PageWrapper>
    );
  }
}

Page.displayName = "IndexPage";

const Loading = () => (
  <section className="container-fluid">
    <h1>Loading...</h1>
  </section>
);
