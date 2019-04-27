(function (m) {
  'use strict';

  m = m && m.hasOwnProperty('default') ? m['default'] : m;

  function levelItems(value) {
    let elements = value;
    if (!Array.isArray(elements)) {
      elements = [elements];
    }
    return elements.map(e => m('div', { class: 'level-item' }, e));
  }

  function level(left, right) {
    return m('nav', { class: 'level' }, [
      m('div', { class: 'level-left' }, [
        levelItems(left),
      ]),
      m('div', { class: 'level-right' }, [
        levelItems(right),
      ]),
    ]);
  }

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

    function deleteTalk(talkId) {
      return m.request({
        method: 'DELETE',
        url: `/api/talks/${talkId}`,
      })
        .then(() => {
          loadTalks();
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
              m('th', 'Description'),
              m('th', 'Speaker Name'),
              m('th', 'Speaker Title'),
              m('th', 'Track'),
              m('th', 'When'),
              m('th', 'Actions'),
            ]),
          ]),
          m('tbody', [
            talks.map(talk => m('tr', { key: talk.id }, [
              m('td', talk.name),
              m('td', talk.description),
              m('td', talk.speakerName),
              m('td', talk.speakerTitle),
              m('td', talk.track),
              m('td', talk.when),
              m('td', [
                m(`a[href=/talks/${talk.id}].button.is-small`, { oncreate: m.route.link }, 'Edit'),
                m('a.button.is-small.is-danger', { onclick: () => deleteTalk(talk.id) }, 'Delete'),
              ]),
            ])),
          ]),
        ]),
      ]),
      oninit: () => loadTalks(),
    };
  }

  function textElement({
    label, value, onchange, size,
  }) {
    return m('input.input', {
      type: 'text',
      class: size ? `is-${size}` : '',
      placeholder: label,
      value,
      oninput: e => onchange(e.target.value),
    });
  }

  function textareaElement({
    label, value, onchange, size,
  }) {
    return m('textarea.textarea', {
      class: size ? `is-${size}` : '',
      placeholder: label,
      value,
      oninput: e => onchange(e.target.value),
    });
  }

  function dateElement({ value, onchange, size }) {
    return m('input.input', {
      type: 'date',
      class: size ? `is-${size}` : '',
      value,
      oninput: e => onchange(e.target.value),
    });
  }

  function timeElement({ value, onchange, size }) {
    return m('input.input', {
      type: 'time',
      class: size ? `is-${size}` : '',
      value,
      oninput: e => onchange(e.target.value),
    });
  }

  function pad(num, size) {
    let s = `${num}`;
    while (s.length < size) s = `0${s}`;
    return s;
  }

  function getTime(date) {
    const hours = date.getHours();
    const minutes = date.getMinutes();
    return `${pad(hours, 2)}:${pad(minutes, 2)}`;
  }

  function getDay(date) {
    const year = date.getFullYear();
    const month = date.getMonth() + 1;
    const day = date.getDate();
    return `${pad(year, 4)}-${pad(month, 2)}-${pad(day, 2)}`;
  }

  function setTime(currentDate, newTime, callback) {
    const parts = newTime.split(':');
    const hours = parseInt(parts[0], 10);
    const minutes = parseInt(parts[1], 10);
    currentDate.setHours(hours);
    currentDate.setMinutes(minutes);
    callback(currentDate);
  }

  function setDay(currentDate, newDay, callback) {
    const parts = newDay.split('-');
    const year = parseInt(parts[0], 10);
    const month = parseInt(parts[1], 10) - 1;
    const day = parseInt(parts[2], 10);
    currentDate.setFullYear(year);
    currentDate.setMonth(month);
    currentDate.setDate(day);
    callback(currentDate);
  }

  function getTimezoneText(value) {
    let offset = '';
    if (value.getFullYear() > 1900) {
      const timezoneOffset = value.getTimezoneOffset();
      const absoluteOffset = Math.abs(timezoneOffset);
      const offsetHours = Math.floor(absoluteOffset / 60);
      const offsetMinutes = absoluteOffset % 60;
      const sign = timezoneOffset < 0 ? '-' : '';
      offset = ` (${sign}${pad(offsetHours, 2)}:${pad(offsetMinutes, 2)})`;
    }

    const timezone = Intl.DateTimeFormat().resolvedOptions().timeZone;
    return `${timezone}${offset}`;
  }

  function input({
    label, value, onchange, size, multiline, date, time,
  }) {
    let internalInput = null;

    if (multiline) {
      internalInput = textareaElement({
        label, value, onchange, size,
      });
    } else if (date && time) {
      internalInput = m('div.field-body', [
        m('div.field', [
          dateElement({
            value: getDay(value),
            onchange: newDay => setDay(value, newDay, onchange),
            size,
          }),
        ]),
        m('div.field', [
          timeElement({
            value: getTime(value),
            onchange: newTime => setTime(value, newTime, onchange),
            size,
          }),
          m('p.help', getTimezoneText(value)),
        ]),
      ]);
    } else if (date) {
      internalInput = dateElement({ value, onchange, size });
    } else if (time) {
      internalInput = timeElement({ value, onchange, size });
    } else {
      internalInput = textElement({
        label, value, onchange, size,
      });
    }

    return m('div.field', [
      m('label.label', label),
      m('div.control', [
        internalInput,
      ]),
    ]);
  }

  const TalkForm = {
    view: vnode => m('div', [
      m('form', [
        input({
          label: 'Name',
          value: vnode.attrs.talk.name,
          onchange: (value) => { vnode.attrs.talk.name = value; },
          size: 'medium',
        }),
        input({
          label: 'Description',
          value: vnode.attrs.talk.description,
          onchange: (value) => { vnode.attrs.talk.description = value; },
          multiline: true,
        }),
        input({
          label: 'Speaker Name',
          value: vnode.attrs.talk.speakerName,
          onchange: (value) => { vnode.attrs.talk.speakerName = value; },
        }),
        input({
          label: 'Speaker Title',
          value: vnode.attrs.talk.speakerTitle,
          onchange: (value) => { vnode.attrs.talk.speakerTitle = value; },
        }),
        input({
          label: 'Track',
          value: vnode.attrs.talk.track,
          onchange: (value) => { vnode.attrs.talk.track = value; },
        }),
        input({
          label: 'When',
          date: true,
          time: true,
          value: vnode.attrs.talk.when,
          onchange: (value) => { vnode.attrs.talk.when = value; },
        }),
      ]),
    ]),
  };

  const ErrorBox = {
    view: vnode => (vnode.attrs.error ? m('div', { class: 'notification is-danger' }, [
      m('button', { class: 'delete', onclick: () => vnode.attrs.onclose() }),
      vnode.attrs.error,
    ]) : null),
  };

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

  const SettingsView = {
    view: () => m('div', [
      m('h1', { class: 'title is-3' }, 'Settings'),
    ]),
  };

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

  const routes = [
    {
      name: 'Talks',
      path: '/talks',
      component: Talks,
    },
    {
      name: 'New Talk',
      path: '/talks/new',
      component: NewTalk,
    },
    {
      name: 'Edit Talk',
      path: '/talks/:talkId',
      component: EditTalk,
    },
    {
      name: 'Settings',
      path: '/settings',
      component: SettingsView,
    },
  ];

  const Router = {
    getDefaultPath() {
      return routes[0].path;
    },
    getRoutes() {
      return routes;
    },
    getMithrilRoutes(layoutComponent) {
      const mithrilRoutes = {};
      routes.forEach((r) => {
        mithrilRoutes[r.path] = {
          render: (function renderProxy(component) {
            return vnode => m(layoutComponent, m(component, vnode.attrs));
          }(r.component)),
        };
      });
      return mithrilRoutes;
    },
  };

  const displayedRoutePaths = [
    '/talks',
    '/settings',
  ];

  function getDisplayedRoutes() {
    return Router.getRoutes()
      .filter(r => displayedRoutePaths.indexOf(r.path) >= 0);
  }

  function navbarTitle(text) {
    return m('a', { class: 'navbar-item', href: '/admin/' }, [
      m('h1', { class: 'title is-4' }, text),
    ]);
  }

  function navbarBurger() {
    return m('a', {
      role: 'button',
      class: 'navbar-burger burger',
      'aria-label': 'menu',
      'aria-expanded': 'false',
      'data-target': 'navbarBasicExample',
    }, [
      m('span', { 'aria-hidden': 'true' }),
      m('span', { 'aria-hidden': 'true' }),
      m('span', { 'aria-hidden': 'true' }),
    ]);
  }

  function navbarMenuItems() {
    return getDisplayedRoutes().map((route) => {
      const isActiveClass = route.path === m.route.get() ? 'is-active' : '';
      return m('div', { class: 'navbar-start' }, [
        m(`a[href=${route.path}]`, { class: `navbar-item ${isActiveClass}`, oncreate: m.route.link }, route.name),
      ]);
    });
  }

  function navbarMenu() {
    return m('div', { id: 'navbarBasicExample', class: 'navbar-menu' }, [
      m('div', { class: 'navbar-start' }, navbarMenuItems()),
    ]);
  }

  const Navbar = {
    view: () => m('div', { class: 'navbar is-light', role: 'navigation', 'aria-label': 'main navigation' }, [
      m('div', { class: 'container' }, [
        m('div', { class: 'navbar-brand' }, [
          navbarTitle('Gonference'),
          navbarBurger(),
        ]),
        navbarMenu(),
      ]),
    ]),
  };

  const App = {
    view: vnode => m('div', [
      m(Navbar),
      m('div', { class: 'container x-padded-container' }, [
        m('div', { id: 'content' }, vnode.children),
      ]),
    ]),
  };

  m.route(
    document.getElementById('app'),
    Router.getDefaultPath(),
    Router.getMithrilRoutes(App),
  );

}(m));
