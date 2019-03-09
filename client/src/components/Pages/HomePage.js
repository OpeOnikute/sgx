import React, { Component } from 'react';
import './HomePage.scss';
import WelcomeForm from './HomePage/WelcomeForm';

class HomePage extends Component {
  onSubmit = (evt) => {
    evt.preventDefault();
  };

  render() {
    return (
      <div className="Page Page--blue">
        <div className="Page__body">
          <div className="Page__center-card">
            <div className="Page__center-card__heading">
              <div className="Page__center-card__heading--main">Welcome!</div>
              <div className="Page__center-card__heading--sub">
                No long talk.
                <br />
                Write a cool story with a friend as mad as you.
              </div>
            </div>
            <WelcomeForm />
          </div>
        </div>
      </div>
    );
  }
}

export default HomePage;
