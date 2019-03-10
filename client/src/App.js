import React, { Component } from 'react';
import './App.scss';
import HomePage from './components/Pages/HomePage';
import { Route } from 'react-router-dom';
import InvitePage from './components/Pages/InvitePage';

class App extends Component {
  render() {
    return (
      <div className="App">
        <Route path="/" exact component={HomePage} />
        <Route path="/invite" exact component={InvitePage} />
      </div>
    );
  }
}

export default App;
