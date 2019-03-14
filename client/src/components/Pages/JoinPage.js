import PropTypes from 'prop-types';
import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { getStory } from '../../api/story';
import buttons from '../../scss/buttons.module.scss';
import forms from '../../scss/forms.module.scss';
import pages from '../../scss/pages.module.scss';
import placeholders from '../../scss/placeholders.module.scss';
import './JoinPage.scss';

class JoinPage extends Component {
  static propTypes = {
    match: PropTypes.shape({
      params: PropTypes.shape({ inviteLink: PropTypes.string.isRequired }),
    }),
  };

  state = { story: null, formError: null };

  async componentWillMount() {
    const { inviteLink } = this.props.match.params;
    const storyData = await getStory({ inviteLink });
    this.setState({ story: storyData.data });
  }

  render() {
    const { story, formError } = this.state;
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
                    <form className={forms.Form}>
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
                        />
                      </div>
                      <div className={pages.CenterCardButtonWrapper1}>
                        <button className={buttons.Button} type="submit">
                          Join Session
                        </button>
                      </div>
                    </form>
                  </div>
                </div>
              </div>
            ) : (
              <div className={pages.CenterCard}>
                <div className={pages.CenterCardHeading}>
                  <div className={pages.CenterCardHeadingSub}>
                    {story.title} by {story.playerOne.name}
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
