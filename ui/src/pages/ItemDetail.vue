<template>
  <item-detail-card
    :img="item_img"
    :title="item_title"
    :detail="item_detail"
    :price="item_price"
  />
  <large-space />
  <v-row no-gutters>
    <v-col cols="5">
      <div class="text-body-1">
        出品者へのレビュー
      </div>
      <small-space />
      <review-list :reviews="reviews" />
    </v-col>
    <v-col cols="1"></v-col>
    <v-col cols="5">
      <v-form @submit.prevent v-if="auth">
        <base-text-area label="商品へのコメント" />
        <base-button
          text="コメントを送信する"
          variant="outlined"
          @click="postChat"
        />
      </v-form>
      <large-space />
      <comment-list />
    </v-col>
    <v-col cols="1"></v-col>
  </v-row>
</template>

<script>
import {
  BaseButton,
  BaseTextArea,
  CommentList,
  ItemDetailCard,
  LargeSpace,
  ReviewList,
  SmallSpace,
} from '@/components'
import { onMounted, watch, ref, computed } from 'vue';
import { useRoute } from 'vue-router';
import axios from 'axios'
import { useStore } from 'vuex';

export default {
  name: 'ItemDetail',
  components: {
    'item-detail-card': ItemDetailCard,
    'large-space': LargeSpace,
    'small-space': SmallSpace,
    'review-list': ReviewList,
    'comment-list': CommentList,
    'base-text-area': BaseTextArea,
    'base-button': BaseButton,
  },
  setup(){
    const store = useStore()
    const item_title = ref('')
    const item_detail = ref('')
    const item_img = ref('')
    const item_price = ref('')
    const reviews = ref([])
    const chats = ref([])
    const route = useRoute()

    const auth = computed(() => store.state.auth)

    const postChat = () => {
      alert("post")
    }

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
      chats,
      reviews,
      postChat,
      auth,
    }
  }
}

</script>