import m from 'mithril';
import TalkForm from '../components/talkForm';
import level from '../components/bulma/level';
import ErrorBox from '../components/errorBox';

function NewTalk() {
  const talk = {
    name: '',
    description: '',
    speakerName: '',
    speakerTitle: '',
    track: '',
    when: '',
  };

  let error = null;

  function save() {
    error = '';
    m.request({
      method: 'POST',
      url: '/api/talks',
      data: talk,
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
