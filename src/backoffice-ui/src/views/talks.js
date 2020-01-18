import m from 'mithril';
import level from '../components/bulma/level';

function Talks() {
  let talks = [];

  function loadTalks() {
    return m.request({
      method: 'GET',
      url: '/api/talks',
    }).then((result) => {
      talks = result;
    });
  }

  return {
    view: () => m('div', [
      level(
        m('h1', { class: 'title is-3' }, 'Talks'),
        m('a[href=#!/talks/new]', { oncreate: m.route.link }, [
          m('button', { class: 'button is-primary' }, 'New'),
        ]),
      ),
      m('table', { class: 'table is-fullwidth is-striped' }, [
        m('thead', [
          m('tr', [
            m('th', 'Name'),
            // m('th', 'Description'),
            m('th', 'Speaker Name'),
            // m('th', 'Speaker Title'),
            // m('th', 'Track'),
            // m('th', 'When'),
            m('th', 'Actions'),
          ]),
        ]),
        m('tbody', [
          talks.map(talk => m('tr', { key: talk.id }, [
            m('td', talk.name),
            // m('td', talk.description),
            m('td', talk.speakerName),
            // m('td', talk.speakerTitle),
            // m('td', talk.track),
            // m('td', talk.when),
            m('td', [
              m(`a[href=#!/talks/${talk.id}].button.is-small`, 'Edit'),
            ]),
          ])),
        ]),
      ]),
    ]),
    oninit: () => loadTalks(),
  };
}

export default Talks;
