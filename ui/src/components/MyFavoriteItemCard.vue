<template>
  <v-card
    class="mx-4 my-4"
    width="300"
    height="350"
  >
    <v-img
      :src="img"
      height="200px"
      cover
    ></v-img>
    <v-card-text>
      <div class="text-h6 text--primary">
        {{abbreviatedTitle}}
      </div>
      <div class="mt-1"></div>
      <div class="text--primary">
        {{price}}円
      </div>
    </v-card-text>
    <v-card-actions>
      <v-btn color="orange" @click="toDetail(itemID)">
        詳細
      </v-btn>
      <v-btn color="orange" @click="deleteItem(favoriteID)">
        削除
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import { ref } from '@vue/reactivity'
import { useRouter } from 'vue-router'
import { abbreviateText } from '@/utils'
import axios from 'axios'

export default {
  name: 'MyFavoriteItem',
  props: {
    title: { type: String },
    price: { type: String },
    img: { type: String },
    itemID: { type: String },
    favoriteID: {type: String}
  },
  setup(props) {
    const router = useRouter()
    const abbreviatedTitle = ref(abbreviateText(props.title, 9, '...'))

    const toDetail = (itemID) => {
      router.push('/item/' + itemID)
    }

    const deleteItem = async(favoriteID) => {
      const url = '/favorite/delete/' + favoriteID
      await axios.delete(url)
      window.confirm('お気に入り登録を解除しました')
      router.go()
    }

    return {
      toDetail,
      deleteItem,
      abbreviatedTitle
    }
  }
}

</script>
