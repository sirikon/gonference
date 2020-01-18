import m from 'mithril';
import Router from '../router';

const displayedRoutePaths = [
  '/talks',
  '/settings',
];

function getDisplayedRoutes() {
  return Router.getRoutes()
    .filter(r => displayedRoutePaths.indexOf(r.path) >= 0);
}

function navbarTitle(text) {
  return m('a', { class: 'navbar-item', href: '#!/' }, [
    m('h1', { class: 'title is-4' }, text),
  ]);
}

function navbarBurger(state) {
  console.log(state);
  return m('a', {
    role: 'button',
    class: `navbar-burger burger ${state.menuOpen ? 'is-active' : ''}`,
    'aria-label': 'menu',
    'aria-expanded': 'false',
    'data-target': 'navbarBasicExample',
    onclick: () => state.menuOpen = !state.menuOpen
  }, [
    m('span', { 'aria-hidden': 'true' }),
    m('span', { 'aria-hidden': 'true' }),
    m('span', { 'aria-hidden': 'true' }),
  ]);
}

function navbarMenuItems() {
  return getDisplayedRoutes().map((route) => {
    const isActiveClass = route.path === m.route.get() ? 'is-active' : '';
    return m('div', { class: `navbar-start` }, [
      m(`a[href=#!${route.path}]`, {
        class: `navbar-item ${isActiveClass}`,
        oncreate: m.route.link,
      }, route.name),
    ]);
  });
}

function navbarMenu(state) {
  return m('div', { id: 'navbarBasicExample', class: `navbar-menu ${state.menuOpen ? 'is-active' : ''}` }, [
    m('div', { class: 'navbar-start' }, navbarMenuItems(state)),
    state.username ? navbarProfile(state.username) : null,
  ]);
}

function navbarProfile(username) {
  return m('div.navbar-end', [
    m('div.navbar-item.has-dropdown.is-hoverable', [
      m('a.navbar-link is-arrowless', `Hello, ${username}!`),
      m('div.navbar-dropdown', [
        m('a.navbar-item[href="#!/change-password"]', 'Change password'),
        m('a.navbar-item[href="/logout"]', 'Log out'),
      ]),
    ]),
  ]);
}

function Navbar() {
  const state = {
    username: null,
    menuOpen: false
  };

  function fetchUser() {
    m.request({
      method: 'GET',
      url: '/api/me',
    })
      .then((result) => {
        state.username = result.username;
      })
      .catch(() => {});
  }

  return {
    oninit: () => fetchUser(),
    view: () => m('div', { class: 'navbar is-light', role: 'navigation', 'aria-label': 'main navigation' }, [
      m('div', { class: 'container' }, [
        m('div', { class: 'navbar-brand' }, [
          navbarTitle('Gonference'),
          navbarBurger(state),
        ]),
        navbarMenu(state),
      ]),
    ]),
  };
}

export default Navbar;
