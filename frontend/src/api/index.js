import axios from 'axios';

const endpoint = 'http://localhost:8888';

export default {
  async getComments() {
    const url = `${endpoint}/api/comments/`;
    return axios.get(url);
  },
};
