function createStory({ name, email, storyTitle }) {
  return Promise.resolve({
    _id: 'knf2pun34m3kfnp238h432infwqejknf',
    inviteID: 'erknfp23uehf3fr23jfno32u4',
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
