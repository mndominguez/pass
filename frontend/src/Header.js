import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { Menu, Segment } from 'semantic-ui-react';

export default class Header extends Component {
    state = { activeItem: 'home' }
    
    handleItemClick = (e, { name }) => this.setState({ activeItem: name });

    render() {
        const { activeItem } = this.state;

        return (
            <div>
                <Menu pointing secondary>
                    <Link to='/'><Menu.Item name="home" active={activeItem === 'home'} onClick={this.handleItemClick} /></Link>
                    <Link to='/residents'><Menu.Item name="residents" active={activeItem === 'residents'} onClick={this.handleItemClick} /></Link>
                    <Link to='/cars'><Menu.Item name="cars" active={activeItem === 'cars'} onClick={this.handleItemClick} /></Link>
                </Menu>
            </div>
        );
    }
};

