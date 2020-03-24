import API from '@/api';

const getters = {
  comments: (state) => state.comments,
};

const actions = {
  getComments: async () => API.getComments(),
};

const mutations = {
  addComment: async (_, payload) => {
    await API.addComment(payload);
    console.log(payload);
  },
};

const state = {
  comments: [],
  comment: '',
};

export default {
  actions,
  mutations,
  getters,
  state,
};
