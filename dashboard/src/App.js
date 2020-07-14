import React, { Component } from 'react';
import { Button } from '@material-ui/core';
import Grid from '@material-ui/core/Grid';
import { gql } from '@apollo/client';
import { ApolloClient, InMemoryCache } from '@apollo/client';
import './App.css';
import Ship from './Ship'

const client = new ApolloClient({
  uri: 'http://localhost:8080/query',
  cache: new InMemoryCache()
});



export default class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      ship: null
    };
    this.setupShip = this.setupShip.bind(this);
  }

  async setupShip () {
    const ship = await this.getShipInfo()
    this.setState({ ship: ship })
  }

  async getShipInfo() {
    return new Promise(resolve => {
      client.query({
        query: gql`
          query GetShip {
              ship {
                  name
                  location
                  fuelLevel
                  deliveries {
                      numberOfPackages
                      uuid
                      deliveryDate
                  }
                  crew {
                      crewName
                      crewMembers {
                          role
                          name
                      }
                  }
              }
          }
        `
      })
      .then((result) => {
        console.log(result)
        resolve(result.data.ship)
      });
    })
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
        <Grid container spacing={3}>
          <Grid item sm={12} md={6} >
            <p>
              Planet Express!
            </p>
            <Button color="primary" variant="contained" onClick={this.setupShip}>Get Ship Information</Button>
          </Grid>
          <Grid item sm={12} md={6} >
            <Ship ship={this.state.ship} />
          </Grid>
        </Grid>
        </header>
      </div>
    );
  }
}

