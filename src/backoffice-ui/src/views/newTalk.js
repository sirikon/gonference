import m from 'mithril';
import TalkForm from '../components/talkForm';
import level from '../components/bulma/level';
import ErrorBox from '../components/errorBox';

function now() {
  const d = new Date();
  d.setSeconds(0);
  d.setMilliseconds(0);
  return d;
}

function NewTalk() {
  const talk = {
    name: '',
    description: '',
    speakerName: '',
    speakerTitle: '',
    track: '',
    when: now(),
  };

  let error = null;

  function save() {
    error = '';
    m.request({
      method: 'POST',
      url: '/api/talks',
      body: talk,
      withCredentials: true
    })
      .then(() => {
        window.history.back();
      })
      .catch(() => {
        error = 'There was a problem while saving the talk.';
      });
  }

  return {
    view: () => m('div', [
      m(ErrorBox, { error, onclose: () => { error = null; } }),
      level(
        m('h1', { class: 'title is-3' }, 'New Talk'),
        [
          m('button', { class: 'button', onclick: () => { window.history.back(); } }, 'Cancel'),
          m('button', { class: 'button is-primary', onclick: () => save() }, 'Save'),
        ],
      ),
      m(TalkForm, { talk }),
    ]),
  };
}

export default NewTalk;
