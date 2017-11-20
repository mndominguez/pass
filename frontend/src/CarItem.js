import React, { Component } from 'react';
import { Button, Table } from 'semantic-ui-react'

class CarItem extends Component {
    render() {
        return (
            <Table.Row>
                <Table.Cell>{this.props.id}</Table.Cell>
                <Table.Cell>{this.props.plate}</Table.Cell>
                <Table.Cell>{this.props.model}</Table.Cell>
                <Table.Cell>{this.props.resident}</Table.Cell>
            </Table.Row>
        );
    }
}

export default CarItem;