import m from 'mithril';

function input({
  label, value, oninput, size, multiline,
}) {
  return m('div.field', [
    m('label.label', label),
    m('div.control', [
      multiline ? m('textarea.textarea', {
        class: size ? `is-${size}` : '',
        placeholder: label,
        value,
        oninput: e => oninput(e),
      }) : m('input.input', {
        type: 'text',
        class: size ? `is-${size}` : '',
        placeholder: label,
        value,
        oninput: e => oninput(e),
      }),
    ]),
  ]);
}

export default input;
