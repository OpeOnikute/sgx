import React, { Component } from 'react';
import { Redirect } from 'react-router-dom';
import './InvitePage.scss';

const APP_INVITE_URL = 'https://sgx.com/link';

class InvitePage extends Component {
  state = {
    userHasCopiedText: false,
  };

  onClickInviteLink = (evt) => {
    this.copyTextToClipboard(evt.target.innerText);
  };

  fallbackCopyTextToClipboard = (text) => {
    const textArea = document.createElement('textarea');
    textArea.value = text;
    document.body.appendChild(textArea);
    textArea.focus();
    textArea.select();

    try {
      const successful = document.execCommand('copy');
      if (successful) {
        this.setState({ userHasCopiedText: true });
      }
    } catch (err) {}
    document.body.removeChild(textArea);
  };

  copyTextToClipboard = async (text) => {
    if (!navigator.clipboard) {
      this.fallbackCopyTextToClipboard(text);
      return;
    }

    try {
      await navigator.clipboard.writeText(text);
      this.setState({ userHasCopiedText: true });
    } catch (error) {}
  };

  render() {
    const { story } = this.props.location.state;

    if (!story) {
      return <Redirect to="/" />;
    }

    const { userHasCopiedText } = this.state;

    return (
      <div className="Page Page--blue">
        <div className="Page__body">
          <div className="Page__center-card">
            <div className="Page__center-card__body">
              <div className="Invite">
                <div className="Invite__header" />
                <div className="Invite__message">
                  You're almost ready!
                  <br />
                  Share this link to invite your friend.
                </div>
                <div className="Invite__link__message Invite__link__message--success">
                  {userHasCopiedText && 'Copied!'}
                </div>
                <button
                  className="Invite__link"
                  onClick={this.onClickInviteLink}
                >
                  {APP_INVITE_URL + '/' + story.inviteID}
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default InvitePage;
