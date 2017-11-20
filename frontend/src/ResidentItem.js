import React, { Component } from 'react';
import { Button, Table } from 'semantic-ui-react';
import axios from 'axios';
import { Link } from 'react-router-dom';

class ResidentItem extends Component {
    constructor(props) {
        super(props);
        this.state = { openChart: false };
        this.renderChart = this.renderChart.bind(this);
    }

    renderChart() {
        axios
        .get("/events")
        .then((result) => {
           this.setState({ residents: result.data });
        });
        this.setState({ openChart: true });
    }

    render() {
        return (
            <Table.Row>
                <Table.Cell>{this.props.id}</Table.Cell>
                <Table.Cell>{this.props.name}</Table.Cell>
                <Table.Cell>{this.props.address}</Table.Cell>
                <Table.Cell>{this.props.cars}</Table.Cell>
                <Table.Cell>
                    <Link to="/residents/:id/events/">
                        <Button 
                        color='violet' 
                        >
                        Ver Historial
                        </Button>
                    </Link>
                </Table.Cell>
            </Table.Row>
        );
    }
}

export default ResidentItem;