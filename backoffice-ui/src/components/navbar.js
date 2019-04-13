import m from 'mithril';
import Router from '../router';

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
  return Router.getRoutes().map((route) => {
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

export default Navbar;
