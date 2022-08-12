<template>
  <small-space />
  <item-card
    :img="img"
    :title="title"
    :price="price"
  />
  <large-space />
  <v-form @submit.prevent>
    <base-text-area
      label="出品者へのレビューを書く"
      v-model="detail"
    />
    <base-select
      label="5段階評価"
      :items="rates"
      v-model="rate"
    />
    <base-button
      text="レビューを送信する"
      variant="outlined"
      @click="postReview"
    />
  </v-form>
</template>

<script>
import { computed, onMounted, ref, watch } from 'vue'
import axios from 'axios'
import { useRoute, useRouter } from 'vue-router'
import { BaseButton, BaseSelect, BaseTextArea, ItemCard, LargeSpace, SmallSpace } from '@/components'
import { useStore } from 'vuex'

export default {
  name: 'PostReview',
  components: {
    'small-space': SmallSpace,
    'large-space': LargeSpace,
    'item-card': ItemCard,
    'base-text-area': BaseTextArea,
    'base-button': BaseButton,
    'base-select': BaseSelect,
  },
  setup() {
    const route = useRoute()
    const router = useRouter()
    const store= useStore()

    const user_id = computed(() => store.state.user_id)
    const detail = ref('')
    const rate = ref(3)
    const img = ref('')
    const title = ref('')
    const price = ref('')
    const seller_id = ref('')
    const rates = [1, 2, 3, 4, 5]

    const getWishDetail = async () => {
      const url = 'item/' + route.params.item_id
      const {data} = await axios.get(url)
      img.value = data.item.Img
      title.value = data.item.Title
      price.value = data.item.Price
      seller_id.value = data.item.SellerID
    }

    const postReview = async() => {
      try {
        const url = '/review/create'
        const params = new URLSearchParams()
        params.append('detail',  detail.value)
        params.append('rate', rate.value)
        params.append('reviewer_id', user_id.value)
        params.append('reviewee_id', seller_id.value)
        await axios.post(url, params)
        router.push('/mypage/purchased')
      } catch (e) {
        console.log(e)
      }
    }

    onMounted(async () => {
      getWishDetail()
    })

    watch(
      detail, () => detail.value,
      rate, () => rate.value,
      img, () => img.value,
      title, () => title.value,
      price, () => price.value,
      seller_id, () => seller_id.value,
      rate, () => rate.value,
    )

    return {
      detail,
      rate,
      img,
      title,
      price,
      rates,
      postReview,
    }
  }
}

</script>