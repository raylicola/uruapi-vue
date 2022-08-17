<template>
  <small-space />
  <navigation-drawer />
  <div v-if="items.length != 0">
    <div class="text-h5">
      お気に入り登録した商品
    </div>
    <small-space />
    <v-row>
    <my-favorite-item-card
      v-for="item in items"
      :key="item.ID"
      :title="item.Title"
      :price="item.Price"
      :img="item.Img"
      :itemID="item.ItemID"
      :favoriteID="item.ID"
    />
  </v-row>
  </div>
  <div v-else>
    <div class="text-h5">
      お気に入り登録した商品がありません
    </div>
  </div>
</template>

<script>
import {
  SmallSpace,
  NavigationDrawer,
  MyFavoriteItemCard,
} from '@/components'
import axios from 'axios'
import { onMounted, watch, ref, computed } from 'vue';
import { useStore } from 'vuex';
import { useRouter } from 'vue-router';

export default {
  name: 'MyFavoriteItem',
  components: {
    'my-favorite-item-card': MyFavoriteItemCard,
    'small-space': SmallSpace,
    'navigation-drawer': NavigationDrawer,
  },
  setup(){
    const store = useStore()
    const router = useRouter()
    const items = ref([])
    const userID = computed(() => store.state.userID)

    const getMyFavoriteItem = async () => {
      const url = '/favorite/' + userID.value
      const {data} = await axios.get(url)
      items.value = data.favorites
    }

    const toDetail = (item_id) => {
      router.push('/item/' + item_id)
    }

    onMounted(async () => {
      getMyFavoriteItem()
    })

    watch(
      items, () => items.value,
    )

    return {
      items,
      toDetail,
    }
  }
}

</script>