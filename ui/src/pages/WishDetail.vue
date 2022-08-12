<template>
  <small-space />
  <wish-detail-card
    :title="wish_title"
    :detail="wish_detail"
  />
  <small-space />
  <div class="text-h5">
    提案されている商品
  </div>
  <small-space />
  <v-row>
    <item-card
      v-for="item in items"
      :key="item.ID"
      :title="item.Title"
      :price="item.Price"
      :img="item.Img"
      @click="toDetail(item.ID)"
    />
  </v-row>
</template>

<script>
import { WishDetailCard, SmallSpace, ItemCard } from '@/components'
import axios from 'axios'
import { onMounted, watch, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

export default {
  name: 'WishDetail',
  components: {
    'wish-detail-card': WishDetailCard,
    'small-space': SmallSpace,
    'item-card': ItemCard,
  },
  setup(){

    const wish_title = ref('')
    const wish_detail = ref('')
    const items = ref([])
    const route = useRoute()
    const router = useRouter()

    const getWishDetail = async () => {
      const url = 'wish/' + route.params.wish_id
      const {data} = await axios.get(url)
      wish_title.value = data.wish.Title
      wish_detail.value = data.wish.Detail
      items.value = data.items
    }

    const toDetail = (item_id) => {
      router.push('/item/' + item_id)
    }

    onMounted(async () => {
      getWishDetail()
    })

    watch(
      wish_title, () => wish_title.value,
      wish_detail, () => wish_detail.value,
      items, () => items.value
    )

    return {
      wish_title,
      wish_detail,
      items,
      toDetail,
    }
  }
}

</script>