import m from 'mithril';
import input from './input';
import Debug from './debug';

const TalkForm = {
  view: vnode => m('div', [
    m(Debug, vnode.attrs.talk),
    m('form', [
      input({
        label: 'Name',
        value: vnode.attrs.talk.name,
        onchange: (value) => { vnode.attrs.talk.name = value; },
        size: 'medium',
      }),
      input({
        label: 'Description',
        value: vnode.attrs.talk.description,
        onchange: (value) => { vnode.attrs.talk.description = value; },
        multiline: true,
      }),
      input({
        label: 'Speaker Name',
        value: vnode.attrs.talk.speakerName,
        onchange: (value) => { vnode.attrs.talk.speakerName = value; },
      }),
      input({
        label: 'Speaker Title',
        value: vnode.attrs.talk.speakerTitle,
        onchange: (value) => { vnode.attrs.talk.speakerTitle = value; },
      }),
      input({
        label: 'Track',
        value: vnode.attrs.talk.track,
        onchange: (value) => { vnode.attrs.talk.track = value; },
      }),
      input({
        label: 'When',
        date: true,
        time: true,
        value: vnode.attrs.talk.when,
        onchange: (value) => { vnode.attrs.talk.when = value; },
      }),
    ]),
  ]),
};

export default TalkForm;
