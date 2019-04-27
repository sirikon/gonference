import m from 'mithril';

const ErrorBox = {
  view: vnode => (vnode.attrs.error ? m('div', { class: 'notification is-danger' }, [
    m('button', { class: 'delete', onclick: () => vnode.attrs.onclose() }),
    vnode.attrs.error,
  ]) : null),
};

export default ErrorBox;
