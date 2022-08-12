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
    </v-col>
    <v-col cols="1"></v-col>
    <v-col cols="5">
      <v-form @submit.prevent v-if="auth">
        <base-text-area label="商品へのコメント" v-model="chat_text"/>
        <base-button
          text="コメントを送信する"
          variant="outlined"
          @click="postChat"
        />
      </v-form>
      <large-space />
      <div class="text-body-1">
        コメント一覧
      </div>
      <chat-card
        v-for="chat in chats"
        :key="chat.ID"
        :content="chat.Content"
        :user_id="chat.UserID"
        :seller_id="seller_id"
      />
    </v-col>
    <v-col cols="1"></v-col>
  </v-row>
</template>

<script>
import {
  BaseButton,
  BaseTextArea,
  ChatCard,
  ItemDetailCard,
  LargeSpace,
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
    'base-text-area': BaseTextArea,
    'base-button': BaseButton,
    'chat-card': ChatCard,
  },
  setup(){
    const store = useStore()
    const item_title = ref('')
    const item_detail = ref('')
    const seller_id = ref('')
    const item_img = ref('')
    const item_price = ref('')
    const reviews = ref([])
    const chats = ref([])
    const chat_text = ref('')
    const route = useRoute()

    const auth = computed(() => store.state.auth)
    const user_id = computed(() => store.state.user_id)
    const item_id = route.params.item_id

    const postChat = async() => {
      try {
        const url = '/chat/create'
        const params = new URLSearchParams()
        params.append('item_id', item_id)
        params.append('content', chat_text.value)
        params.append('user_id', user_id.value)
        await axios.post(url, params)
        getItemDetail()
      } catch (e) {
        console.log(e)
      }
    }

    const getItemDetail = async () => {
      const url = 'item/' + item_id
      const {data} = await axios.get(url)
      item_title.value = data.item.Title
      item_detail.value = data.item.Detail
      seller_id.value = data.item.UserID
      item_img.value = data.item.Img
      item_price.value = data.item.Price
      chats.value = data.chats
      reviews.value = data.reviews
    }

    onMounted(async () => {
      getItemDetail()
    })

    watch(
      seller_id, () => seller_id.value,
      item_title, () => item_title.value,
      item_detail, () => item_detail.value,
      chats, () => chats.value,
      reviews, () => reviews.value,
      chat_text, () => chat_text.value
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
      chat_text,
      seller_id,
    }
  }
}

</script>