import React from 'react';
import { Jumbotron, Button, Container } from 'reactstrap';

const JumbotronFluid = props => {
  return (
    <div>
      <Jumbotron fluid className="jumbotronFluid">
        <Container className="jumbotronFluidContainer">
        <div className="row">
          <div className="col-xs-12 col-lg-9">
          <div className="jumbotronFluidBody">
            <h1 className="jumbotronFluidTitle">List Your Space</h1>
            <p className="jumbotronFluidText">Looking to gain income from your empty or unused space? List with Anchorspace to start leasing with many flexible options.
            </p>
          </div>
          </div>
          <div className="col-xs-12 col-lg-3">
          <p className="jumbotronFluidAction lead">
            <Button className="btn-gradient" color="primary" size="lg" href="/listings/new">List Your Space</Button>
          </p>
          </div>
          </div>
        </Container>
      </Jumbotron>
    </div>
  );
};

export default JumbotronFluid;
