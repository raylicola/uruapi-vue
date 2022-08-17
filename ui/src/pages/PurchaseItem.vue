<template>
  <back-button />
  <small-space />
  <div class="text-h5">
    購入内容の確認
  </div>
  <small-space />
  <v-row>
    <v-col cols="1"></v-col>
    <v-col cols="4">
      <item-card
        :title="title"
        :price="price"
        :img="img"
      />
    </v-col>
    <v-col cols="3">
      <small-space />
      <div class="d-flex justify-center">
        <div class="text-h5">支払金額</div>
        <div class="mx-4"></div>
        <div class="text-h5">￥{{price}}</div>
      </div>
      <v-divider></v-divider>
      <small-space />
      <div class="d-flex justify-center">
        <base-button
          text="購入する"
          @click="purchase"
        />
      </div>
    </v-col>
  </v-row>
</template>

<script>
import { onMounted, watch, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import axios from 'axios'
import { BaseButton, ItemCard, SmallSpace, BackButton, } from '@/components';
import { useStore } from 'vuex';

export default {
  name: 'PurchaseItem',
  components: {
    'small-space': SmallSpace,
    'item-card': ItemCard,
    'base-button': BaseButton,
    'back-button': BackButton,
  },
  setup(){
    const title = ref('')
    const img = ref('')
    const price = ref('')
    const seller_id = ref('')
    const route = useRoute()
    const router = useRouter()
    const store = useStore()
    const item_id = route.params.item_id

    const getItemDetail = async () => {
      const url = 'item/' + item_id
      const {data} = await axios.get(url)
      title.value = data.item.Title
      img.value = data.item.Img
      price.value = data.item.Price
      seller_id.value = data.item.UserID
    }

    const purchase = async() => {
      try {
        const url = 'transaction/create'
        const params = new URLSearchParams()
        params.append('purchaser_id', store.state.user_id)
        params.append('seller_id', seller_id.value)
        params.append('item_id', item_id)
        await axios.post(url, params)
        router.push('/mypage')
      } catch (e) {
        console.log(e)
      }
    }

    onMounted(async () => {
      getItemDetail()
    })

    watch(
      title, () => title.value,
      price, () => price.value,
      img, () => img.value,
    )

    return {
      title,
      price,
      img,
      purchase,
    }
  }
}

</script>