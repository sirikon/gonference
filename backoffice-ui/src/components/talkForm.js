import m from 'mithril';
import level from './bulma/level';

function TalkForm() {
  const data = {
    name: '',
    description: '',
  };

  let error = null;

  function save() {
    error = '';
    m.request({
      method: 'POST',
      url: '/api/talks',
      data,
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
      error ? m('div', { class: 'notification is-danger' }, [
        m('button', { class: 'delete', onclick: () => { error = null; } }),
        error,
      ]) : null,
      level(
        m('h1', { class: 'title is-3' }, 'New Talk'),
        [
          m('button', { class: 'button', onclick: () => { window.history.back(); } }, 'Cancel'),
          m('button', { class: 'button is-primary', onclick: () => save() }, 'Save'),
        ],
      ),
      m('form', [
        m('div', { class: 'field' }, [
          m('label', { class: 'label' }, 'Name'),
          m('div', { class: 'control' }, [
            m('input', {
              class: 'input is-medium',
              type: 'text',
              placeholder: 'Name',
              value: data.name,
              oninput: (e) => { data.name = e.target.value; },
            }),
          ]),
        ]),
        m('div', { class: 'field' }, [
          m('label', { class: 'label' }, 'Description'),
          m('div', { class: 'control' }, [
            m('textarea', {
              class: 'textarea',
              placeholder: 'Description',
              value: data.description,
              oninput: (e) => { data.description = e.target.value; },
            }),
          ]),
        ]),
      ]),
    ]),
  };
}

export default TalkForm;
