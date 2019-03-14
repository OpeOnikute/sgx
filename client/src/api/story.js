// const BASE_URL = process.env.REACT_APP_SERVER_BASE_URL;

function createStory({ name, email, storyTitle }) {
  return Promise.resolve({
    data: {
      _id: '5c8836752a40710004605dac',
      title: 'my story',
      inviteCode: '-L_o4AnWPyQq1hz1OEiy',
      playerOne: {
        uid: '-L_o4AnWPyQq1hz1OEix',
        name: 'ope',
        email: 'ope@checl-dc.com',
      },
      playerTwo: { uid: '', name: '', email: '' },
      content: null,
      status: 'open',
      created: '2019-03-12T22:45:09.985879148Z',
    },
  });
  // return fetch(`${BASE_URL}/story`, {
  //   method: 'POST',
  //   body: JSON.stringify({ name, email, storyTitle }),
  //   headers: { 'Content-Type': 'application/json' },
  // }).then((data) => data.json());
}

function getStory({ inviteLink }) {
  return Promise.resolve({
    data: {
      _id: '5c8836752a40710004605dac',
      title: 'my story',
      inviteCode: '-L_o4AnWPyQq1hz1OEiy',
      playerOne: {
        uid: '-L_o4AnWPyQq1hz1OEix',
        name: 'ope',
        email: 'ope@checl-dc.com',
      },
      playerTwo: { uid: '', name: '', email: '' },
      content: null,
      status: 'open',
      created: '2019-03-12T22:45:09.985879148Z',
    },
  });
}

export { createStory, getStory };
