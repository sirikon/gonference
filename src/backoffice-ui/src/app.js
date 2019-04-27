import m from 'mithril';
import Navbar from './components/navbar';

const App = {
  view: vnode => m('div', [
    m(Navbar),
    m('div', { class: 'container x-padded-container' }, [
      m('div', { id: 'content' }, vnode.children),
    ]),
  ]),
};

export default App;
