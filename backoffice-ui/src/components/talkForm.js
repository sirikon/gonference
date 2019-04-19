import m from 'mithril';
import input from './input';

const TalkForm = {
  view: vnode => m('div', [
    m('form', [
      input({
        label: 'Name',
        value: vnode.attrs.talk.name,
        oninput: (e) => { vnode.attrs.talk.name = e.target.value; },
        size: 'medium',
      }),
      input({
        label: 'Description',
        value: vnode.attrs.talk.description,
        oninput: (e) => { vnode.attrs.talk.description = e.target.value; },
        multiline: true,
      }),
      input({
        label: 'Speaker Name',
        value: vnode.attrs.talk.speakerName,
        oninput: (e) => { vnode.attrs.talk.speakerName = e.target.value; },
      }),
      input({
        label: 'Speaker Title',
        value: vnode.attrs.talk.speakerTitle,
        oninput: (e) => { vnode.attrs.talk.speakerTitle = e.target.value; },
      }),
      input({
        label: 'Track',
        value: vnode.attrs.talk.track,
        oninput: (e) => { vnode.attrs.talk.track = e.target.value; },
      }),
      input({
        label: 'When',
        value: vnode.attrs.talk.when,
        oninput: (e) => { vnode.attrs.talk.when = e.target.value; },
      }),
    ]),
  ]),
};

export default TalkForm;
