<template>
  <div id="app">
    <section class="container">
      <myheader></myheader>
      <h1>グルメ検索</h1>
      <p v-if="keyword.length > 0">
        {{ keyword }}
      </p>
      <p v-else>
        文字を入力してください
      </p>
      <input type="text" v-model="keyword">
      <button @click="clear()">clear</button>
      <input @click="getShopList" type="button" value="取得">
      <div class="list_container">
        <article v-for="sh in shops" :key="sh.Name" class="shop_list">
          <a :href="sh.Urls.Pc" class="item" target="_blank" rel="noopener noreferrer">
            <div class="item_container">
              <div class="item_logo">
                <img :src="sh.LogoImage" alt="">
              </div>
              <div class="item_name">
                {{sh.Name}}
              </div>
            </div>
          </a>
        </article>
      </div>
    </section>
  </div>
</template>

<script>
import myheader from './components/myheader.vue'

export default {
  name: 'App',
  components: {
    myheader
  },
  data () {
    return {
      keyword: '',
      shops: [],
      // url: ''
    }
  },
  methods: {
    clear () {
      this.keyword = ''
    },
    getShopList () {
      this.axios.get('http://localhost:8080/keyword?keyword='+this.keyword)
      .then(function(response){
        this.shops = response.data
        console.log(response)
      }.bind(this))
      .catch(function(error){
        console.log('取得に失敗しました。',error)
      })
    }
  },
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.list_container {
  display: flex;
  flex-wrap: wrap;
}
.shop_list {
  width: 80%;
  margin: 2rem;
}
.item {
  width: 100%;
  height: 100%;
  display: block;
}
.item_container {
  display: flex;
}

</style>
