import m from 'mithril';
import input from '../components/input';
import level from '../components/bulma/level';

function ChangePasswordView() {
  const data = {
    currentPassword: '',
    newPassword: '',
    repeatNewPassword: '',
  };

  function save() {
    m.request({
      method: 'POST',
      url: '/api/me/change-password',
      body: data,
      withCredentials: true
    })
      .then(() => {
        window.history.back();
      })
      .catch((err) => {
        // eslint-disable-next-line no-console
        console.log('Error', err);
      });
  }

  return {
    view: () => m('div', [
      level(
        m('h1', { class: 'title is-3' }, 'Change Password'),
        [
          m('button', { class: 'button is-primary', onclick: () => save() }, 'Change password'),
        ],
      ),
      m('form', [
        input({
          label: 'Current password',
          value: data.currentPassword,
          secure: true,
          onchange: (value) => { data.currentPassword = value; },
        }),
        input({
          label: 'New password',
          value: data.newPassword,
          secure: true,
          onchange: (value) => { data.newPassword = value; },
        }),
        input({
          label: 'Repeat new password',
          value: data.repeatNewPassword,
          secure: true,
          onchange: (value) => { data.repeatNewPassword = value; },
        }),
      ]),
    ]),
  };
}

export default ChangePasswordView;
