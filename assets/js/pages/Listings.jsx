import React, { Component } from 'react'
import { Link } from 'react-router-dom'
import NavigationAlt from '../components/NavigationAlt/NavigationAlt.js'
import Cards from '../components/Cards/Cards.js'
import JumbotronFluid from '../components/Jumbotron/JumbotronFluid.js'
import Select from '../components/Select/Select.js'
import Footer from '../components/Footer.jsx';
import { PageWrapper } from '../styles'
import { FetchListings, FetchNeighborhoods } from '../utils/networkHelpers.js'

export default class Page extends Component {
  constructor (props) {
    super(props)
    console.log(props)
    this.state = {
      loading: true,
      listings: [],
      neighborhoods: [],
      loadingNeighborhoods: true,
      neighborhoodID: '',
      currentNeighborhood: null
    }
    this.handleNeighborhoodUpdate = this.handleNeighborhoodUpdate.bind(this)
    this.loadListings = this.loadListings.bind(this)
  }

  componentDidMount () {
    this.loadNeighborhoods()
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
      let neighborhood = this.paramNeighborhood()
      if (neighborhood != null) {
        this.setState({
          currentNeighborhood: neighborhood
        })
        this.loadListings(neighborhood.id)
      } else {
        this.loadListings()
      }
    })
  }

  paramNeighborhood () {
    let url = new URL(window.location.href)
    let id = url.searchParams.get('neighborhood')
    return this.neighborByID(id)
  }

  neighborByID (id) {
    var i = 0
    for (i in this.state.neighborhoods) {
      if (this.state.neighborhoods[i].id === id) {
        return this.state.neighborhoods[i]
      }
    }
    return null
  }

  loadListings (neighborhoodID) {
    this.setState({
      loading: true
    })

    var params = new URLSearchParams()
    if (neighborhoodID != null && neighborhoodID !== '') {
      params.set('NeighborhoodID', neighborhoodID)
    }

    FetchListings(params).then(json => {
      this.setState({
        loading: false,
        listings: json
      })
    })
  }

  handleNeighborhoodUpdate (e) {
    this.setState({
      currentNeighborhood: this.neighborByID(e.target.value)
    })
    this.loadListings(e.target.value)
  }

  sortedNeighborhoods () {
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

  render () {
    let currentNeighborhood = this.state.currentNeighborhood

    let title = currentNeighborhood == null ? 'Neighborhood Listings' : `${currentNeighborhood.name} Listings`

    return (
      <PageWrapper>
        <NavigationAlt />
        <Select
          options={this.sortedNeighborhoods()}
          handleChange={this.handleNeighborhoodUpdate}
          title='Select Neighborhood (show all)'
          displayField='name'
          selectedID={currentNeighborhood == null ? '' : currentNeighborhood.id} />
        {
          this.state.listings.length > 0 ? <Cards listings={this.state.listings} title={title} /> : <div></div>
        }
        <JumbotronFluid />
        <Footer />
      </PageWrapper>
    )
  }
}

Page.displayName = 'ListingsPage'
