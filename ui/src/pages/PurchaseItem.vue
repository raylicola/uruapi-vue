<template>
<div class="text-h5">
  購入画面
</div>
</template>

<script>
import { onMounted, watch, ref } from 'vue';
import { useRoute } from 'vue-router';
import axios from 'axios'

export default {
  name: 'PurchaseItem',

  setup(){
    const item_title = ref('')
    const item_detail = ref('')
    const item_img = ref('')
    const item_price = ref('')
    const reviews = ref([])
    const chats = ref([])
    const route = useRoute()

    const getItemDetail = async () => {
      const url = 'item/' + route.params.item_id
      const {data} = await axios.get(url)
      item_title.value = data.item.Title
      item_detail.value = data.item.Detail
      item_img.value = data.item.Img
      item_price.value = data.item.Price
      chats.value = data.chats
      reviews.value = data.reviews
    }

    onMounted(async () => {
      getItemDetail()
    })

    watch(
      item_title, () => item_title.value,
      item_detail, () => item_detail.value,
      chats, () => chats.value,
      reviews, () => reviews.value
    )

    return {
      item_title,
      item_detail,
      item_price,
      item_img,
    }
  }
}

</script>