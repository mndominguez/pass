import React, { Component } from 'react';
import ChartBuilder from './ChartBuilder';

class LineChart extends Component {
    render() {
        return <ChartBuilder {...this.props} />
    }
};

export default LineChart;
