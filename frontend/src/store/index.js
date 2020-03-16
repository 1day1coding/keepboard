import Vue from 'vue';
import Vuex from 'vuex';
import comments from './modules/comments';

Vue.use(Vuex);

const debug = process.env.NODE_ENV !== 'production';

export default new Vuex.Store({
  strict: debug,
  modules: {
    comments,
  },
});
