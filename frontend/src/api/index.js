import axios from 'axios';

const endpoint = 'http://localhost:8888';

export default {
  async getComments() {
    const url = `${endpoint}/api/comment/`;
    return axios.get(url);
  },
  async addComment({ comment }) {
    const url = `${endpoint}/api/comment/`;
    return axios.post(url, { comment });
  },
};
