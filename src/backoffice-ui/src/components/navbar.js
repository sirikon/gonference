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
          navbarBurger(),
        ]),
        navbarMenu(),
        state.username ? navbarProfile(state.username) : null,
      ]),
    ]),
  };
}

export default Navbar;
