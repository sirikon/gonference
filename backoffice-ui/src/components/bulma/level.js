import m from 'mithril';

function levelItems(value) {
  let elements = value;
  if (!Array.isArray(elements)) {
    elements = [elements];
  }
  return elements.map(e => m('div', { class: 'level-item' }, e));
}

function level(left, right) {
  return m('nav', { class: 'level' }, [
    m('div', { class: 'level-left' }, [
      levelItems(left),
    ]),
    m('div', { class: 'level-right' }, [
      levelItems(right),
    ]),
  ]);
}

export default level;
