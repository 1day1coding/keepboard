<template>
  <div class="container">
    <br><br><br>
    <div>
      <table class="table">
        <thead class="thead-dark">
        <th class="table-sharp">
          #
        </th>
        <th class="table-blank">
        </th>
        <th class="table-body">
          comment
        </th>
        </thead>
        <tbody>
        <tr v-for="(comment, idx) in comments" v-bind:key="comment">
          <td class="table-sharp">
            {{idx + 1}}
          </td>
          <th class="table-blank">
            <i class="far fa-hand-point-right"></i>
          </th>
          <td class="table-body">
            <label :name='"comment-"+idx' :id='"comment-"+idx'>
              {{comment}}
            </label>
          </td>
        </tr>
        </tbody>
      </table>
    </div>
    <div>
      <div class="beginner col-11">
        <label class="col-12">
          <input id="new-comment" class="swal2-input"/>
        </label>
      </div>
      <div class="follow col-1">
        <button id="enter" type="submit" class="swal2-actions btn btn-dark">Enter</button>
      </div>
    </div>
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
  .table-sharp {
    width: 100px;
  }
  .table-blank {
    width: 40px;
    text-align: right;
  }
  .table-body {
    width: 600px;
    text-align: left;
  }
  .table-last {
    width: 10px;
  }
  .beginner {
    float: left;
  }
  .follow {
    float: left;
    display:inline-block;
  }
</style>
