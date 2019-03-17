import { SERVER_BASE_URL } from '../config';

function createStory({ name, email, storyTitle }) {
  return fetch(`${SERVER_BASE_URL}/story`, {
    method: 'POST',
    body: JSON.stringify({
      playerName: name,
      playerEmail: email,
      title: storyTitle,
    }),
    headers: { 'Content-Type': 'application/json' },
  }).then((data) => data.json());
}

function getStory({ inviteLink }) {
  return fetch(`${SERVER_BASE_URL}/story?f=invitecode&v=${inviteLink}`).then(
    (data) => data.json(),
  );
}

export { createStory, getStory };
