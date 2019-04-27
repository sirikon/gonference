import m from 'mithril';
import TalkForm from '../components/talkForm';
import level from '../components/bulma/level';
import ErrorBox from '../components/errorBox';

function EditTalk(initialVnode) {
  const { talkId } = initialVnode.attrs;

  let talk = {
    id: '',
    name: '',
    description: '',
    speakerName: '',
    speakerTitle: '',
    track: '',
    when: new Date(),
  };

  let error = null;

  function fetch() {
    error = '';
    m.request({
      method: 'GET',
      url: `/api/talks/${talkId}`,
    })
      .then((result) => {
        talk = result;
        talk.when = new Date(talk.when);
      })
      .catch(() => {
        error = 'There was a problem while getting the talk information.';
      });
  }

  function save() {
    error = '';
    m.request({
      method: 'PUT',
      url: `/api/talks/${talkId}`,
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
    oninit: () => fetch(),
    view: () => m('div', [
      m(ErrorBox, { error, onclose: () => { error = null; } }),
      level(
        m('h1', { class: 'title is-3' }, 'Edit Talk'),
        [
          m('button', { class: 'button', onclick: () => { window.history.back(); } }, 'Cancel'),
          m('button', { class: 'button is-primary', onclick: () => save() }, 'Update'),
        ],
      ),
      m(TalkForm, { talk }),
    ]),
  };
}

export default EditTalk;
