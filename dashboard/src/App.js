import React from 'react';
import { Button } from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import { gql } from '@apollo/client';
import { ApolloClient, InMemoryCache } from '@apollo/client';
import './App.css';

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    padding: theme.spacing(2),
    textAlign: 'center',
    color: theme.palette.text.primary,
  },
}));

const client = new ApolloClient({
  uri: 'http://localhost:8080/query',
  cache: new InMemoryCache()
});


function getShipInfo() {
  client
  .query({
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
  .then(result => console.log(result));
}


function App() {
  const classes = useStyles();
  return (
    <div className="App">
      <header className="App-header">
      <Grid container spacing={3}>
        <Grid item sm={12} md={6} >
          <p>
            Planet Express!
          </p>
          <Button color="primary" variant="contained" onClick={getShipInfo}>Get Ship Information</Button>
        </Grid>
        <Grid item sm={12} md={6} >
          <Paper className={classes.paper}>xs=6</Paper>
        </Grid>
      </Grid>
      </header>
    </div>
  );
}

export default App;
