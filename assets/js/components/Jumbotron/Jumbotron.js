import React, { Component } from 'react'
import { Jumbotron, Button, ButtonDropdown, DropdownToggle, DropdownMenu, DropdownItem } from 'reactstrap'

class Hero extends Component {
  constructor (props) {
    super(props)
    this.toggle = this.toggle.bind(this)
    this.state = {
      dropdownOpen: false
    }
  }

  toggle () {
    this.setState({ dropdownOpen: !this.state.dropdownOpen })
  }

  render() {
    let dropdownItems = this.props.neighborhoods.map(neighborhood =>
      <DropdownItem href={`/listings?neighborhood=${neighborhood.id}`}>{neighborhood.name}</DropdownItem>
    )

    return (
      <Jumbotron className='jumbotronHero mb-0' style={{ backgroundImage: '/assets/images/office-bg.jpg' }}>
        <div className='jumbotronHeroContainer container'>
          <div className='row justify-content-center jumbotronHeroBody'>
            <div className='col-md-10 col-lg-8'>
              <h1 className='jumbotronHeroTitle'>Find Your Space</h1>
              <div className='jumbotronHeroText lead'>
                <ButtonDropdown isOpen={this.state.dropdownOpen} toggle={this.toggle}>
                  <DropdownToggle caret size='lg' color='white'>
                    View Neighborhood Listings
                  </DropdownToggle>
                  <DropdownMenu>
                    {dropdownItems}
                  </DropdownMenu>
                </ButtonDropdown>
              </div>
            </div>
          </div>
        </div>
      </Jumbotron>
    )
  }
}

export default Hero
