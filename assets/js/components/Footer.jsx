import React from 'react';
import {HandleLogout} from '../utils/networkHelpers'

class Footer extends React.Component {
  constructor(props) {
    super(props);

    this.handleSignout = this.handleSignout.bind(this);
  }

  handleSignout(e){
    e.preventDefault();
    console.log("signout");
    HandleLogout(this.props.params, () => window.location.reload(false), (err) => console.log("error:",err) );
    return false;
  }

  render() {
    return (
  <footer className="footer">
    <div className="footerContainer">
      <div className="footerBody">
        <section className="footerBrand">
          <div><img src='/assets/images/logo-muted.png' alt='Anchorspace Logo Muted' /></div>
        </section>
        <section className="footerSection footerCompany">
          <ul className="footerSection__list">
            <li className="footerSection__title">Company</li>
            <li className="footerSection__list-item">About Us</li>
            <li className="footerSection__list-item">Get Help</li>
            <li className="footerSection__list-item">Resources</li>
            <li className="footerSection__list-item">News &amp; Events</li>
            <li className="footerSection__list-item">Contact</li>
            <li className="footerSection__list-item">Blog</li>
          </ul>
        </section>
        <section className="footerSection footerAccount">
          <ul className="footerSection__list">
            <li className="footerSection__title">Account</li>
            <li className="footerSection__list-item"><a href="/admin/panel">Space Admin</a></li>
            <li className="footerSection__list-item">Space Owner</li>
            <li className="footerSection__list-item">Space Renter</li>
            <li className="footerSection__list-item"><a href="" onClick={this.handleSignout}>Sign Out</a></li>
          </ul>
        </section>
        <section className="footerSection footerExplore">
          <ul className="footerSection__list">
            <li className="footerSection__title">Explore</li>
            <li className="footerSection__list-item">View Available Spaces</li>
            <li className="footerSection__list-item">View Upcoming Spaces</li>
            <li className="footerSection__list-item">List New Space</li>
          </ul>
        </section>
      </div>
      <footer className="footerBottom">
      <ul>
        <li>&copy; Anchorspace</li>
        <li>Terms &amp; Privacy</li>
        <li>Site Map</li>
      </ul>
      </footer>
    </div>
  </footer>
    )
  }
}

export default Footer;
