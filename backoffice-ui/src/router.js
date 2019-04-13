import m from 'mithril';
import Talks from './views/talks';
import Settings from './views/settings';

const routes = [
  {
    name: 'Talks',
    path: '/talks',
    component: Talks,
  },
  {
    name: 'Settings',
    path: '/settings',
    component: Settings,
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
          return () => m(layoutComponent, m(component));
        }(r.component)),
      };
    });
    return mithrilRoutes;
  },
};

export default Router;
