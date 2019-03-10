const BASE_URL = 'http://127.0.0.1:8000';

function createStory({ name, email, storyTitle }) {
  return Promise.resolve({
    _id: 'knf2pun34m3kfnp238h432infwqejknf',
    link: 'https://sgx.com/link/erknfp23uehf3fr23jfno32u4',
    initiatorName: name,
    initiatorEmail: email,
    title: storyTitle,
  });
  // return fetch(`${BASE_URL}/story`, {
  //   method: 'POST',
  //   body: JSON.stringify({ name, email, storyTitle }),
  //   headers: { 'Content-Type': 'application/json' },
  // }).then((data) => data.json());
}

export default createStory;
