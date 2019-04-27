import m from 'mithril';

function textElement({
  label, value, onchange, size,
}) {
  return m('input.input', {
    type: 'text',
    class: size ? `is-${size}` : '',
    placeholder: label,
    value,
    oninput: e => onchange(e.target.value),
  });
}

function textareaElement({
  label, value, onchange, size,
}) {
  return m('textarea.textarea', {
    class: size ? `is-${size}` : '',
    placeholder: label,
    value,
    oninput: e => onchange(e.target.value),
  });
}

function dateElement({ value, onchange, size }) {
  return m('input.input', {
    type: 'date',
    class: size ? `is-${size}` : '',
    value,
    oninput: e => onchange(e.target.value),
  });
}

function timeElement({ value, onchange, size }) {
  return m('input.input', {
    type: 'time',
    class: size ? `is-${size}` : '',
    value,
    oninput: e => onchange(e.target.value),
  });
}

function pad(num, size) {
  let s = `${num}`;
  while (s.length < size) s = `0${s}`;
  return s;
}

function getTime(date) {
  const hours = date.getHours();
  const minutes = date.getMinutes();
  return `${pad(hours, 2)}:${pad(minutes, 2)}`;
}

function getDay(date) {
  const year = date.getFullYear();
  const month = date.getMonth() + 1;
  const day = date.getDate();
  return `${pad(year, 4)}-${pad(month, 2)}-${pad(day, 2)}`;
}

function setTime(currentDate, newTime, callback) {
  const parts = newTime.split(':');
  const hours = parseInt(parts[0], 10);
  const minutes = parseInt(parts[1], 10);
  currentDate.setHours(hours);
  currentDate.setMinutes(minutes);
  callback(currentDate);
}

function setDay(currentDate, newDay, callback) {
  const parts = newDay.split('-');
  const year = parseInt(parts[0], 10);
  const month = parseInt(parts[1], 10) - 1;
  const day = parseInt(parts[2], 10);
  currentDate.setFullYear(year);
  currentDate.setMonth(month);
  currentDate.setDate(day);
  callback(currentDate);
}

function getTimezoneText(value) {
  let offset = '';
  if (value.getFullYear() > 1900) {
    const timezoneOffset = value.getTimezoneOffset();
    const absoluteOffset = Math.abs(timezoneOffset);
    const offsetHours = Math.floor(absoluteOffset / 60);
    const offsetMinutes = absoluteOffset % 60;
    const sign = timezoneOffset < 0 ? '-' : '';
    offset = ` (${sign}${pad(offsetHours, 2)}:${pad(offsetMinutes, 2)})`;
  }

  const timezone = Intl.DateTimeFormat().resolvedOptions().timeZone;
  return `${timezone}${offset}`;
}

function input({
  label, value, onchange, size, multiline, date, time,
}) {
  let internalInput = null;

  if (multiline) {
    internalInput = textareaElement({
      label, value, onchange, size,
    });
  } else if (date && time) {
    internalInput = m('div.field-body', [
      m('div.field', [
        dateElement({
          value: getDay(value),
          onchange: newDay => setDay(value, newDay, onchange),
          size,
        }),
      ]),
      m('div.field', [
        timeElement({
          value: getTime(value),
          onchange: newTime => setTime(value, newTime, onchange),
          size,
        }),
        m('p.help', getTimezoneText(value)),
      ]),
    ]);
  } else if (date) {
    internalInput = dateElement({ value, onchange, size });
  } else if (time) {
    internalInput = timeElement({ value, onchange, size });
  } else {
    internalInput = textElement({
      label, value, onchange, size,
    });
  }

  return m('div.field', [
    m('label.label', label),
    m('div.control', [
      internalInput,
    ]),
  ]);
}

export default input;
