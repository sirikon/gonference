import m from 'mithril';
import App from './app';
import Router from './router';

m.route(
  document.getElementById('app'),
  Router.getDefaultPath(),
  Router.getMithrilRoutes(App),
);
