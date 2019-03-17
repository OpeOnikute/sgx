import { SERVER_BASE_URL } from '../config';

export function createStory({ name, email, storyTitle }) {
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

export function getStoryByInviteLink({ inviteLink }) {
  return fetch(`${SERVER_BASE_URL}/story?f=invitecode&v=${inviteLink}`).then(
    (response) => response.json().then((jsonData) => ({ response, jsonData })),
  );
}

export function joinStory({ inviteCode, playerName, playerEmail }) {
  return fetch(`${SERVER_BASE_URL}/story/join`, {
    method: 'POST',
    body: JSON.stringify({ code: inviteCode, playerName, playerEmail }),
    headers: { 'Content-Type': 'application/json' },
  }).then((response) =>
    response.json().then((jsonData) => ({ response, jsonData })),
  );
}
