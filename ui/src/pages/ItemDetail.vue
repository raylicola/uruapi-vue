<template>
  <small-space />
  <v-card
    class="mx-4 my-4"
    width="1000"
  >
    <v-row no-gutters>
      <v-col cols="4">
        <v-img
          :src="item_path"
          cover
        ></v-img>
      </v-col>
      <v-col cols="8">
        <v-card-text>
          <div class="d-flex mb-2 align-center">
            <base-avatar
              size="40"
              :src="icon_path"
              @click="toUserProfile"
            />
            <div class="text-h7 text--primary mx-5">
              {{username}}
            </div>
          </div>
          <div class="text-h6 text--primary">
            {{item_title}}
          </div>
          <div class="text-h6 text--primary">
            ￥{{item_price}}円
          </div>
          <div>
            <base-button
              v-if="auth"
              text="商品を購入する"
              class="ma-3"
              @click="purchase"
            />
          </div>
          <div class="my-4"></div>
          <div class="text-h7 mb-2 text--primary">
            商品の説明
          </div>
          <div class="text--primary" style="white-space: pre-line;">
            {{item_detail}}
          </div>
        </v-card-text>
      </v-col>
    </v-row>
  </v-card>
  <large-space />
  <v-row no-gutters>
    <v-col cols="5">
      <div class="text-body-1">
        出品者へのレビュー
      </div>
      <small-space />
      <review-card
        v-for="review in reviews"
        :key="review.ID"
        :review="review"
      />
    </v-col>
    <v-divider
        class="mx-4"
        vertical
      ></v-divider>
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
      <div class="text-body-1" v-if="chats.length != 0">
        コメント一覧
      </div>
      <chat-card
        v-for="chat in chats"
        :key="chat.ID"
        :chat="chat"
        :seller_id="seller_id"
      />
    </v-col>
    <v-col cols="1"></v-col>
  </v-row>
</template>

<script>
import {
  BaseAvatar,
  BaseButton,
  BaseTextArea,
  ChatCard,
  LargeSpace,
  ReviewCard,
  SmallSpace,
} from '@/components'
import { onMounted, watch, ref, computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import axios from 'axios'
import { useStore } from 'vuex';
import { doc, getDoc } from "firebase/firestore";
import { db } from '@/firebase'

export default {
  name: 'ItemDetail',
  components: {
    'large-space': LargeSpace,
    'small-space': SmallSpace,
    'base-text-area': BaseTextArea,
    'base-button': BaseButton,
    'chat-card': ChatCard,
    'review-card': ReviewCard,
    'base-avatar': BaseAvatar,
  },
  setup(){
    const store = useStore()
    const item_title = ref('')
    const item_detail = ref('')
    const seller_id = ref('')
    const item_path = ref('')
    const item_price = ref('')
    const icon_path = ref('')
    const username = ref('')
    const reviews = ref([])
    const chats = ref([])
    const chat_text = ref('')
    const route = useRoute()
    const router = useRouter()

    const auth = computed(() => store.state.auth)
    const user_id = computed(() => store.state.user_id)
    const item_id = route.params.item_id

    const toUserProfile = () => {
      router.push('/user/' + seller_id.value)
    }

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
      seller_id.value = data.item.SellerID
      item_path.value = data.item.Img
      item_price.value = data.item.Price
      chats.value = data.chats
      reviews.value = data.reviews

      const docRef = doc(db, "users", seller_id.value);
      const docSnap = await getDoc(docRef);
      icon_path.value = docSnap.data().icon_path ? docSnap.data().icon_path : require('@/assets/default_icon.jpg')
      username.value = docSnap.data().username
    }

    onMounted(async () => {
      getItemDetail()
    })

    watch(
      seller_id, () => seller_id.value,
      item_title, () => item_title.value,
      item_detail, () => item_detail.value,
      icon_path, () => icon_path.value,
      username, () => username.value,
      chats, () => chats.value,
      reviews, () => reviews.value,
      chat_text, () => chat_text.value
    )

    return {
      item_title,
      item_detail,
      item_price,
      item_path,
      icon_path,
      username,
      chats,
      reviews,
      auth,
      chat_text,
      seller_id,
      postChat,
      toUserProfile,
    }
  }
}

</script>