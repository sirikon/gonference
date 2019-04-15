import m from 'mithril';
import level from '../components/bulma/level';

function Talks() {
  let talks = [];

  function loadTalks() {
    return m.request({
      method: 'GET',
      url: '/api/talks',
    })
      .then((result) => {
        talks = result;
      });
  }

  return {
    view: () => m('div', [
      level(
        m('h1', { class: 'title is-3' }, 'Talks'),
        m('a[href=/talks/new]', { oncreate: m.route.link }, [
          m('button', { class: 'button is-primary' }, 'New'),
        ]),
      ),
      m('table', { class: 'table is-fullwidth is-striped' }, [
        m('thead', [
          m('tr', [
            m('th', 'Name'),
          ]),
        ]),
        m('tbody', [
          talks.map(talk => m('tr', { key: talk.id }, [
            m('td', talk.name),
          ])),
        ]),
      ]),
    ]),
    oninit: () => loadTalks(),
  };
}

export default Talks;
