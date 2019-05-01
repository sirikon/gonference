export default {
  input: 'src/main.js',
  output: {
    file: '../web/assets/backoffice/js/app.js',
    format: 'iife',
    globals: {
        'mithril': 'm'
    }
  },
  external: ['mithril'],
  watch: {
    include: 'src/**'
  }
};
