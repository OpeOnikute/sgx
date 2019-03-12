import React, { Component } from 'react';
import './HomePage.scss';
import { createStory } from '../../api/story';

class HomePage extends Component {
  state = {
    name: '',
    email: '',
    storyTitle: '',
    formError: '',
  };

  onSubmitForm = async (evt) => {
    evt.preventDefault();
    if (!this.validateForm()) return;
    const { name, email, storyTitle } = this.state;
    const storyResp = await createStory({ name, email, storyTitle });
    this.props.history.push({
      pathname: '/invite',
      state: { story: storyResp.data },
    });
  };

  validateForm() {
    const { name, email, storyTitle } = this.state;

    if (!name) {
      this.setState({ formError: 'Please enter your name' });
      return;
    }
    if (email) {
      const re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      if (!re.test(email.toLowerCase())) {
        this.setState({ formError: 'Please enter a valid email address' });
        return;
      }
    }
    if (!storyTitle) {
      this.setState({ formError: 'Please enter a story title' });
      return;
    }

    this.setState({ formError: null });
    return true;
  }

  handleUserInput = (evt) => {
    this.setState({ [evt.target.name]: evt.target.value });
  };

  render() {
    const { formError } = this.state;

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
              <form
                className="Form Form--center-card"
                onSubmit={this.onSubmitForm}
              >
                {formError && (
                  <div className="Form__prompt Form__prompt--error">
                    {formError}
                  </div>
                )}
                <div className="Form__input-group">
                  <label htmlFor="name" className="Form__label">
                    What's your name?
                  </label>
                  <input
                    type="text"
                    name="name"
                    className="Form__input"
                    required
                    onChange={this.handleUserInput}
                  />
                </div>
                <div className="Form__input-group">
                  <label htmlFor="email" className="Form__label">
                    What's your email? (optional)
                  </label>
                  <input
                    type="email"
                    name="email"
                    className="Form__input"
                    onChange={this.handleUserInput}
                  />
                </div>
                <div className="Form__input-group">
                  <label htmlFor="storyTitle" className="Form__label">
                    What's the title of the story?
                  </label>
                  <input
                    type="text"
                    name="storyTitle"
                    className="Form__input"
                    required
                    onChange={this.handleUserInput}
                  />
                </div>

                <div className="Form__input-group">
                  <button type="submit" className="Form__button--submit">
                    Start Session
                  </button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default HomePage;
