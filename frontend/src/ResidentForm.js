import React, { Component } from 'react';
import { Form, Message } from 'semantic-ui-react'
import axios from 'axios';

class ResidentForm extends Component {
    state = { 
        nombre: '', 
        direccion: '', 
        patente: '', 
        modelo: '' 
    }
    
    handleChange = (e, { name, value }) => this.setState({ [name]: value });

    handleSubmit = () => {
        axios
        .post("/resident", {
            name: this.state.nombre,
            address: this.state.direccion,
            plate: this.state.patente,
            model: this.state.modelo
        })
        .then((result) => {
           this.setState({success: true})
        })
        .catch((error) => {
            console.log(error);
        });
    }

    render() {
        const { nombre, direccion, patente, modelo } = this.state
        return (
            <Form onSubmit={this.handleSubmit} success={this.state.success ? true : false}>
                <Form.Group widths='equal'>
                    <Form.Input 
                    label='Nombre completo'
                    placeholder='Nombre completo'
                    name="nombre"
                    value={nombre} 
                    onChange={this.handleChange}
                    />
                    <Form.Input 
                    label='Dirección'
                    placeholder='Dirección'
                    name="direccion"
                    value={direccion}
                    onChange={this.handleChange}
                    />
                </Form.Group>
                <Form.Group widths='equal'>
                    <Form.Input
                    label='Patente'
                    placeholder='Patente'
                    name="patente" 
                    value={patente}
                    onChange={this.handleChange}
                    />
                    <Form.Input
                    label='Modelo'
                    placeholder='Modelo'
                    name="modelo"
                    value={modelo}
                    onChange={this.handleChange}
                    />
                </Form.Group>
                <Message
                success
                header='Agregado!'
                content="Se creó correctamente el nuevo residente"
                />
                <Form.Button color="green">Enviar</Form.Button>
            </Form>
        );
    }
}

export default ResidentForm;