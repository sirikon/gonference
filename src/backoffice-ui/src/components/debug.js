import m from 'mithril';

const Debug = {
  view: vnode => m('pre', JSON.stringify(vnode.attrs, null, 2)),
};

export default Debug;
