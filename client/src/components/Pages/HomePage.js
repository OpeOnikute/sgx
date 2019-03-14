import React, { Component } from 'react';
import { createStory } from '../../api/story';
import buttons from '../../scss/buttons.module.scss';
import forms from '../../scss/forms.module.scss';
import pages from '../../scss/pages.module.scss';
import './HomePage.scss';

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
      <div className={`${pages.Page} ${pages.BluePage}`}>
        <div className={pages.Body}>
          <div className={pages.CenterCard}>
            <div className={pages.CenterCardHeading}>
              <div className={pages.CenterCardHeadingMain}>Welcome!</div>
              <div className={pages.CenterCardHeadingSub}>
                No long talk.
                <br />
                Write a cool story with a friend as mad as you.
              </div>
            </div>
            <div className={pages.CenterCardBody}>
              <div className={pages.CenterCardWrapper}>
                <form className={forms.Form} onSubmit={this.onSubmitForm}>
                  {formError && (
                    <div className={`${forms.Prompt} ${forms.ErrorPrompt}`}>
                      {formError}
                    </div>
                  )}
                  <div className={forms.InputGroup}>
                    <label htmlFor="name" className={forms.Label}>
                      What's your name?
                    </label>
                    <input
                      type="text"
                      name="name"
                      className={forms.Input}
                      required
                      onChange={this.handleUserInput}
                      autoFocus
                    />
                  </div>
                  <div className={forms.InputGroup}>
                    <label htmlFor="email" className={forms.Label}>
                      What's your email? (optional)
                    </label>
                    <input
                      type="email"
                      name="email"
                      className={forms.Input}
                      onChange={this.handleUserInput}
                    />
                  </div>
                  <div className={forms.InputGroup}>
                    <label htmlFor="storyTitle" className={forms.Label}>
                      What's the title of the story?
                    </label>
                    <input
                      type="text"
                      name="storyTitle"
                      className={forms.Input}
                      required
                      onChange={this.handleUserInput}
                    />
                  </div>

                  <div className={forms.InputGroup}>
                    <button type="submit" className={buttons.Button}>
                      Start Session
                    </button>
                  </div>
                </form>
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default HomePage;
