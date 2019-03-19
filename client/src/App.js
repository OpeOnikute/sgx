import React, { Component } from 'react';
import './App.scss';
import HomePage from './components/Pages/HomePage';
import { Route } from 'react-router-dom';
import InvitePage from './components/Pages/InvitePage';
import JoinPage from './components/Pages/JoinPage';

class App extends Component {
  render() {
    return (
      <div className="App">
        <Route path="/" exact component={HomePage} />
        <Route path="/invite" exact component={InvitePage} />
        <Route path="/join/:inviteLink" component={JoinPage} />
      </div>
    );
  }
}

export default App;
