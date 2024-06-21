module.exports = {
  root: true,
  parserOptions: {
    parser: '@babel/eslint-parser',
    sourceType: 'module'
  },
  env: {
    browser: true,
    node: true,
    es6: true
  },
  // extends: ['plugin:vue/recommended', 'eslint:recommended'],
  extends: [],
  rules: {
    // "vue/max-attributes-per-line": 0,
    // "vue/no-v-model-argument": 0,
    'vue/multi-word-component-names': 'off'
  }
}
