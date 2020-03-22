<template>
  <div id="app">
    <sidebar-menu :menu="menu" :collapsed="true"/>
    <router-view />
  </div>
</template>

<script>
export default {
  data() {
    return {
      menu: [
        {
          header: true,
          title: 'Menu',
        },
        {
          href: '/',
          title: 'Home',
          icon: 'fa fa-home',
          child: [
            {
              href: '/board',
              title: 'Board',
              icon: 'fa fa-list',
            },
          ],
        },
      ],
    };
  },
};
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
