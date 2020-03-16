<template>
  <div class="hello">
    <label>
      {{this.comments}}
    </label>
  </div>
</template>

<script>
import Swal from 'sweetalert2';

export default {
  name: 'HelloWorld',
  data() {
    return { comments: [] };
  },
  async created() {
    try {
      const { data } = await this.$store.dispatch('getComments');
      this.comments = data;
    } catch (e) {
      await Swal.fire(`${e}`);
    }
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
  h3 {
    margin: 40px 0 0;
  }

  ul {
    list-style-type: none;
    padding: 0;
  }

  li {
    display: inline-block;
    margin: 0 10px;
  }

  a {
    color: #42b983;
  }
</style>
