import React from 'react';
import { Card, Button, CardImg, CardTitle, CardText, CardGroup,
 CardSubtitle, CardBody } from 'reactstrap';
import { ANCHOR_URL } from '../../constants'

const Cards = props => {
  let rowLength = 4;
  var numRows = props.numberOfRows || -1;

  var i = 0;
  var j = 0;
  let totalListings = props.listings.length;

  var listingRows = [];

  var shouldContinue = true;
  if (totalListings > 0) {
    while (shouldContinue) {
      var row = [];
      while (row.length < rowLength) {
        row.push(props.listings[i % totalListings]);
        i++;
      }
      listingRows.push(row);
      j++;
      if (numRows > 0) {
        shouldContinue = j < numRows;
      } else {
        shouldContinue = i < totalListings;
      }
    }
  }

  let rows = listingRows.map(listings => (
    <CardGroup className="cardGroup">
      <ListingRow listings={listings} />
    </CardGroup>
  ));

  return (
    <div>
      <section className="container-fluid">
        <div className="featuredCards">
          <h1 className="featuredCards__title mb-4 pl-4">{props.title || 'Featured Listings'}</h1>
          <div className="py-2">{rows}</div>
        </div>
      </section>
    </div>
  );
};

const ListingRow = props => {
  return props.listings.map(listing => <ListingCard listing={listing} />);
};

const ListingCard = props => {
  let listing = props.listing;
  let url = ANCHOR_URL + "/listings/" + listing.id;

  return (
    <Card key={listing.id} className="cardGroup__card shadow-lg mb-5 bg-white rounded">
      <div className="row no-gutters">
        <div className="col">
          <CardImg
            className="card__image cardGroup__cardImage"
            top
            width="100%"
            src={listing.picture_url}
            alt="Card image cap"
          />
          <CardBody className="card__body cardGroup__cardBody">
            <CardTitle>
              <a href={url}>{listing.name}</a>
            </CardTitle>
            <CardText>
              {listing.address1}
              <br />
              {listing.address2}
              <br />
              {listing.city}, {listing.state} {listing.zipcode}
            </CardText>
          </CardBody>
        </div>
      </div>
    </Card>
  );
};

export default Cards;
