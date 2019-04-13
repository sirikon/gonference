import m from 'mithril';

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
      m('nav', { class: 'level' }, [
        m('div', { class: 'level-left' }, [
          m('div', { class: 'level-item' }, [
            m('h1', { class: 'title is-3' }, 'Talks'),
          ]),
        ]),
        m('div', { class: 'level-right' }, [
          m('div', { class: 'level-item' }, [
            m('button', { class: 'button is-primary' }, 'New'),
          ]),
        ]),
      ]),
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
