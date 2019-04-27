export default {
  input: 'src/main.js',
  output: {
    file: '../http/assets/backoffice/js/app.js',
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
