import React, { Component } from 'react';
import ResidentsList from './ResidentsList';
import CarsList from './CarsList';
import HomeComponent from './HomeComponent';
import { Button } from 'semantic-ui-react';
import { Switch, Route } from 'react-router-dom';
import Header from './Header';
import ResidentEvents from './ResidentEvents';

const App = () => (
  <div>
    <Header/>
    <Switch>
      <Route exact path="/" component={HomeComponent} />
      <Route exact path="/residents" component={ResidentsList} />
      <Route path="/cars" component={CarsList} />
      <Route path="/residents/:id/events" component={ResidentEvents} />
      <div id="chart"/>
    </Switch>
  </div>
);

export default App;
