import m from 'mithril';
import input from './input';

const TalkForm = {
  view: vnode => m('div', [
    m('form', [
      input({
        label: 'Name',
        value: vnode.attrs.talk.name,
        onchange: (value) => { vnode.attrs.talk.name = value; },
        size: 'medium',
      }),
      input({
        label: 'Slug',
        value: vnode.attrs.talk.slug,
        onchange: (value) => { vnode.attrs.talk.slug = value; },
        size: 'small',
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
        label: 'Speaker Image',
        file: true,
        onchange: (value) => { vnode.attrs.talk.speakerImage = value; console.log(value); },
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
