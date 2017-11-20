import React, { Component } from 'react';
import ResidentItem from './ResidentItem';
import axios from 'axios';
import { Grid, Table, Button } from 'semantic-ui-react';
import ResidentForm from './ResidentForm';

class ResidentsList extends Component {
    constructor(props) {
      super(props);
      this.state = { residents: [] };
      this.handleClickNewResident = this.handleClickNewResident.bind(this);
    }
  
    componentDidMount() {
      this.serverRequest =
        axios
          .get("/residents")
          .then((result) => {
             this.setState({ residents: result.data });
          });
    }

    handleClickNewResident() {
      this.setState({selected: true});
    }
  
    render() {
      const residents = this.state.residents.map((resident, i) => {
        return (
          <ResidentItem key={i} id={resident.Id} name={resident.name} address={resident.address} cars={resident.cars[0].plate.toString()} />
        );
      });
  
      return (
        <Grid centered columns={1} padded>
          <Grid.Column>
            <Table celled>
              <Table.Header>
                <Table.Row>
                  <Table.HeaderCell>Id</Table.HeaderCell>
                  <Table.HeaderCell>Nombre</Table.HeaderCell>
                  <Table.HeaderCell>Dirección</Table.HeaderCell>
                  <Table.HeaderCell>Vehículos</Table.HeaderCell>
                  <Table.HeaderCell>Historial</Table.HeaderCell>
                </Table.Row>
              </Table.Header>
              <Table.Body>
                {residents}
              </Table.Body>
            </Table>
            <Grid.Row centered columns={1}>
              <Grid.Column>
                <Button onClick={this.handleClickNewResident} color='violet'>Agregar Nuevo</Button>
              </Grid.Column>
            </Grid.Row>    
          </Grid.Column>
          <Grid.Column>
            {this.state.selected == true ? <ResidentForm /> : null}
          </Grid.Column>
        </Grid>
      );
    }
  }

  export default ResidentsList;