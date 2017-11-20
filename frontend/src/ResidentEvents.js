import React, { Component } from 'react';
import LineChart from './LineChart';

class ResidentEvents extends Component {
    constructor() {
        super();

        this.state = {
            data: {
              labels: ['Ene', 'Feb', 'Mar', 'Abr',
                'May', 'Jun', 'Jul', 'Ago',
                'Sep', 'Oct', 'Nov', 'Dic'
              ],
              datasets: [
                {
                  title: 'Some Data',
                  color: 'light-blue',
                  values: [25, 40, 30, 35, 8, 52, 17, -4]
                }
              ]
            }
        };
    }
    
    render() {
        return (
            <LineChart title="Eventos" data={this.state.data} height={250} />
        );
    }
};

export default ResidentEvents;
