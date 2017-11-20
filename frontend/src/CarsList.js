import React, { Component } from 'react';
import { Grid, Table } from 'semantic-ui-react'
import axios from 'axios';
import CarItem from './CarItem';

class CarsList extends Component {
    constructor(props) {
        super(props);
        this.state = { cars: [] };
      }
    
      componentDidMount() {
        this.serverRequest =
          axios
            .get("/cars")
            .then((result) => {
               this.setState({ cars: result.data });
            });
      }
    
      render() {
        const cars = this.state.cars.map((car, i) => {
          return (
            <CarItem key={i} id={car.Id} plate={car.plate} model={car.model} resident={car.resident} />
          );
        });
    
        return (
          <Grid columns={1} padded>
            <Grid.Column>
              <Table celled>
                <Table.Header>
                  <Table.Row>
                    <Table.HeaderCell>Id</Table.HeaderCell>
                    <Table.HeaderCell>Plate</Table.HeaderCell>
                    <Table.HeaderCell>Model</Table.HeaderCell>
                    <Table.HeaderCell>Resident</Table.HeaderCell>
                  </Table.Row>
                </Table.Header>
                <Table.Body>
                  {cars}
                </Table.Body>
              </Table>
            </Grid.Column>
          </Grid>
        );
      }
}

export default CarsList;
