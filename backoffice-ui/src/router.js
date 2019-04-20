import m from 'mithril';
import Talks from './views/talks';
import NewTalk from './views/newTalk';
import Settings from './views/settings';
import EditTalk from './views/editTalk';

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
          return vnode => m(layoutComponent, m(component, vnode.attrs));
        }(r.component)),
      };
    });
    return mithrilRoutes;
  },
};

export default Router;
