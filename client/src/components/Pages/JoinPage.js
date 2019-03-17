import PropTypes from 'prop-types';
import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { getStoryByInviteLink, joinStory } from '../../api/story';
import buttons from '../../scss/buttons.module.scss';
import forms from '../../scss/forms.module.scss';
import pages from '../../scss/pages.module.scss';
import placeholders from '../../scss/placeholders.module.scss';
import { RotatingCircleLoader } from '../Utilities/Loader';
import './JoinPage.scss';

class JoinPage extends Component {
  static propTypes = {
    match: PropTypes.shape({
      params: PropTypes.shape({ inviteLink: PropTypes.string.isRequired }),
    }),
  };

  state = {
    story: null,
    formError: null,
    getStoryError: null,
    requestingJoinStory: false,
    name: null,
    email: null,
  };

  async componentWillMount() {
    await this.fetchStory();
  }

  fetchStory = async () => {
    const { inviteLink } = this.props.match.params;

    try {
      const { response, jsonData } = await getStoryByInviteLink({ inviteLink });
      if (!response.ok) {
        this.setState({ getStoryError: jsonData.message });
        return;
      }

      this.setState({ story: jsonData.data });
    } catch (error) {
      this.setState({
        getStoryError:
          'An error occured. Please check your internet connection and try again.',
      });
    }
  };

  onSubmitForm = async (evt) => {
    evt.preventDefault();
    const formError = this.validateForm();
    if (formError) {
      return this.setState({ formError });
    }

    await this.joinStory();
  };

  joinStory = async () => {
    const { story, name, email } = this.state;

    try {
      const { response, jsonData } = await joinStory({
        inviteCode: story.inviteCode,
        playerName: name,
        playerEmail: email,
      });

      if (!response.ok) {
        this.setState({ formError: jsonData.message });
        return;
      }

      this.props.history.push({
        pathname: '/write',
        state: { story: jsonData.data },
      });
    } catch (error) {
      this.setState({
        formError:
          'An error occured. Please check your internet connection and try again.',
        requestingCreateStory: false,
      });
    }
  };

  validateForm() {
    const { name, email } = this.state;

    if (!name) {
      return 'Please enter your name';
    }
    if (email) {
      const re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      if (!re.test(email.toLowerCase())) {
        return 'Please enter a valid email address';
      }
    }

    return false;
  }

  handleUserInput = (evt) => {
    this.setState({ [evt.target.name]: evt.target.value });
  };

  render() {
    const { story, formError, requestingJoinStory } = this.state;

    const submitButton = requestingJoinStory ? (
      <button type="submit" className={buttons.Button} disabled>
        <RotatingCircleLoader
          className={buttons.Loader}
          height={17}
          width={17}
        />
      </button>
    ) : (
      <button type="submit" className={buttons.Button}>
        Join Session
      </button>
    );

    return (
      <div className={`${pages.Page} ${pages.BluePage} Join-Page`}>
        <div className={pages.Body}>
          {story &&
            (story.status === 'open' ? (
              <div className={pages.CenterCard}>
                <div className={pages.CenterCardHeading}>
                  <div className={pages.CenterCardHeadingSub}>
                    Join {story.playerOne.name} to write
                  </div>
                  <div className={pages.CenterCardHeadingMain}>
                    {story.title}
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
                          className={forms.Input}
                          name="name"
                          type="text"
                          required
                          onInput={this.handleUserInput}
                        />
                      </div>
                      <div className={forms.InputGroup}>
                        <label htmlFor="email" className={forms.Label}>
                          What's your email? (optional)
                        </label>
                        <input
                          className={forms.Input}
                          name="email"
                          type="email"
                          onInput={this.handleUserInput}
                        />
                      </div>
                      <div className={forms.InputGroup}>{submitButton}</div>
                    </form>
                  </div>
                </div>
              </div>
            ) : (
              <div className={pages.CenterCard}>
                <div className={pages.CenterCardHeading}>
                  <div className={pages.CenterCardHeadingSub}>
                    {story.title}, by {story.playerOne.name}
                  </div>
                  <div className={pages.CenterCardHeadingMain}>
                    Sorry, this story is closed.
                  </div>
                </div>
                <div className={pages.CenterCardBody}>
                  <div className={placeholders.Banner} />
                  <div className={pages.CenterCardButtonWrapper1}>
                    <Link to="/" className={buttons.Button}>
                      Start a new story
                    </Link>
                  </div>
                </div>
              </div>
            ))}
        </div>
      </div>
    );
  }
}

export default JoinPage;
