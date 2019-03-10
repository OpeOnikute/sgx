import React, { Component } from 'react';
import './HomePage.scss';
import CreateStoryForm from './HomePage/CreateStoryForm';

class HomePage extends Component {
  onCreateStory = ({ story }) => {
    this.props.history.push({ pathname: '/invite', state: { story } });
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
            <div className="Page__center-card__body">
              <CreateStoryForm onCreateStory={this.onCreateStory} />
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default HomePage;
