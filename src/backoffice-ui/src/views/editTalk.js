import m from 'mithril';
import TalkForm from '../components/talkForm';
import level from '../components/bulma/level';
import ErrorBox from '../components/errorBox';
import {objectToFormData} from "../utils/utils";

const CONFIRM_DELETE_TEXT = 'Esta acción eliminará la charla y toda su información. No puede revertirse. Está totalmente seguro/a?';

export default function EditTalk(initialVnode) {
  const { talkId } = initialVnode.attrs;

  let talk = {
    id: '',
    name: '',
    description: '',
    speakerName: '',
    speakerTitle: '',
    speakerImage: null,
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

  function deleteTalk(talkId) {
    if (!confirm(CONFIRM_DELETE_TEXT)) { return; }
    return m.request({
      method: 'DELETE',
      url: `/api/talks/${talkId}`,
    })
        .then(() => {
          window.history.back();
        });
  }

  function save() {
    error = '';
    const formData = objectToFormData(talk);
    m.request({
      method: 'PUT',
      url: `/api/talks/${talkId}`,
      body: formData,
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
    oninit: () => fetch(),
    view: () => m('div', [
      m(ErrorBox, { error, onclose: () => { error = null; } }),
      level(
        m('h1', { class: 'title is-3' }, 'Edit Talk'),
        [
          m('button', { class: 'button', onclick: () => { window.history.back(); } }, 'Cancel'),
        ],
      ),
      m(TalkForm, { talk }),
      level([], [
        m('button.button.is-danger', { onclick: () => deleteTalk(talk.id) }, 'Delete'),
        m('button.button.is-primary', { onclick: () => save() }, 'Update'),
      ])
    ]),
  };
}
