import _ from 'lodash';
import API from '@/api';

const getters = {
  comments: (state) => state.comments,
};

const actions = {
  getComments: async () => API.getComments(),
  // setComments: async({ commit }, payload) => {
  //   await API.setComments(payload);
  //   commit('setComments', {});
  // },
};

const mutations = {
  setComments(state, { comments }) {
    _.assign(state.comments, comments);
  },
};

const state = { comments: [] };

export default {
  actions,
  mutations,
  getters,
  state,
};
