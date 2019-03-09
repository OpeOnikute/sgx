import React, { Component } from 'react';
import './App.scss';
import HomePage from './components/Pages/HomePage';
import { Route } from 'react-router-dom';

class App extends Component {
  render() {
    return (
      <div className="App">
        <Route path="/" component={HomePage} />
      </div>
    );
  }
}

export default App;
