import React, { Component } from 'react';
import Chart from 'frappe-charts/dist/frappe-charts.min.esm'
import 'frappe-charts/dist/frappe-charts.min.css'

class ChartBuilder extends Component {
  componentDidMount () {
    const { title, data, type, height } = this.props
    this.c = new Chart({ parent: this.chart, title, data, type, height });
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps && !nextProps.data || this.props.data !== nextProps.data) {
      const { title, data, type, height } = nextProps;
      this.c = new Chart({ parent: this.chart, title, data, type, height });
    }
  }

  render () {
    return (<div ref={ chart => (this.chart = chart) } />)
  }
}

export default ChartBuilder;
