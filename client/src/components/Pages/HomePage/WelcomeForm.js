import React, { Component } from 'react';
import './WelcomeForm.scss';

class WelcomeForm extends Component {
  state = {
    name: '',
    email: '',
    storyTitle: '',
    formError: '',
  };

  onSubmit = (evt) => {
    evt.preventDefault();

    const { name, email, storyTitle } = this.state;
    if (!this.validateForm()) return;
  };

  validateForm() {
    const { name, email, storyTitle } = this.state;

    if (!name) {
      this.setState({ formError: 'Please enter your name' });
      return;
    }
    if (email) {
      const re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
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
      <form className="Form Form--center-card" onSubmit={this.onSubmit}>
        {formError && (
          <div className="Form__prompt Form__prompt--error">{formError}</div>
        )}
        <div className="Form__input-group">
          <label for="name" className="Form__label">
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
          <label for="email" className="Form__label">
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
          <label for="storyTitle" className="Form__label">
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
    );
  }
}

export default WelcomeForm;
