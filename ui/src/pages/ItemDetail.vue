<template>
  <back-button />
  <small-space />
  <v-card
    class="mx-4 my-4"
    width="1000"
  >
    <v-row no-gutters>
      <v-col cols="4">
        <v-img
          :src="itemPath"
          cover
        ></v-img>
      </v-col>
      <v-col cols="8">
        <v-card-text>
          <div class="d-flex mb-2 align-center">
            <base-avatar
              size="40"
              :src="iconPath"
              @click="toUserProfile"
            />
            <div class="text-h7 text--primary mx-5">
              {{userName}}
            </div>
          </div>
          <div class="text-h6 text--primary">
            {{itemTitle}}
          </div>
          <div class="text-h6 text--primary">
            ￥{{itemPrice}}円
          </div>
          <div v-if="isSold">
            <base-button
              v-if="auth"
              disabled
              text="売り切れました"
              class="ma-3"
            />
          </div>
          <div v-else>
            <base-button
              v-if="auth"
              text="商品を購入する"
              class="ma-3"
              @click="purchaseItem"
            />
            <base-button
              v-if="auth"
              text="お気に入り登録"
              class="ma-3"
              @click="favoriteItem"
            />
          </div>
          <div class="my-4"></div>
          <div class="text-h7 mb-2 text--primary">
            商品の説明
          </div>
          <div class="text--primary" style="white-space: pre-line;">
            {{itemDetail}}
          </div>
        </v-card-text>
      </v-col>
    </v-row>
  </v-card>
  <large-space />
  <v-row>
    <v-col
      cols="12"
      md="6"
    >
      <v-form @submit.prevent v-if="auth">
    <base-text-area label="商品へのコメント" v-model="chatText"/>
    <base-button
      text="コメントを送信する"
      variant="outlined"
      @click="postChat"
    />
    </v-form>
    <small-space />
    <div class="text-body-1" v-if="chats.length != 0">
      コメント一覧
    </div>
    <chat-card
      v-for="chat in chats"
      :key="chat.ID"
      :chat="chat"
      :sellerID="sellerID"
    />
    </v-col>
  </v-row>
</template>

<script>
import {
  BaseAvatar,
  BaseButton,
  BaseTextArea,
  ChatCard,
  LargeSpace,
  SmallSpace,
  BackButton,
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
    'base-avatar': BaseAvatar,
    'back-button': BackButton,
  },
  setup(){
    const store = useStore()
    const isSold = ref(false)
    const itemTitle = ref('')
    const itemDetail = ref('')
    const sellerID = ref('')
    const itemPath = ref('')
    const itemPrice = ref('')
    const iconPath = ref('')
    const userName = ref('')
    const reviews = ref([])
    const chats = ref([])
    const chatText = ref('')
    const route = useRoute()
    const router = useRouter()

    const auth = computed(() => store.state.auth)
    const userID = computed(() => store.state.userID)
    const itemID = route.params.itemID

    const toUserProfile = () => {
      router.push('/user/' + sellerID.value)
    }

    const postChat = async() => {
      try {
        const url = '/chat/create'
        const params = new URLSearchParams()
        params.append('item_id', itemID)
        params.append('content', chatText.value)
        params.append('user_id', userID.value)
        await axios.post(url, params)
        chatText.value = ""
        getItemDetail()
      } catch (e) {
        console.log(e)
      }
    }

    const purchaseItem = () => {
      router.push('/purchase/' + itemID)
    }

    const favoriteItem = async() => {
      const url = 'favorite/create'
      const params = new URLSearchParams()
      params.append('user_id', userID.value)
      params.append('item_id', itemID)
      await axios.post(url, params)
      window.confirm('お気に入りの商品に追加しました')
    }

    const getItemDetail = async () => {
      const url = 'item/' + itemID
      const {data} = await axios.get(url)
      itemTitle.value = data.item.Title
      itemDetail.value = data.item.Detail
      sellerID.value = data.item.SellerID
      itemPath.value = data.item.Img
      itemPrice.value = data.item.Price
      isSold.value = data.item.PurchaserID != ""
      console.log(isSold.value)
      chats.value = data.chats
      reviews.value = data.reviews

      const docRef = doc(db, "users", sellerID.value);
      const docSnap = await getDoc(docRef);
      iconPath.value = docSnap.data().icon_path ? docSnap.data().icon_path : require('@/assets/default_icon.jpg')
      userName.value = docSnap.data().username
    }

    onMounted(async () => {
      getItemDetail()
    })

    watch(
      sellerID, () => sellerID.value,
      itemTitle, () => itemTitle.value,
      itemDetail, () => itemDetail.value,
      iconPath, () => iconPath.value,
      userName, () => userName.value,
      chats, () => chats.value,
      reviews, () => reviews.value,
      chatText, () => chatText.value,
      isSold, () => isSold.value
    )

    return {
      itemTitle,
      itemDetail,
      itemPrice,
      itemPath,
      iconPath,
      isSold,
      userName,
      chats,
      reviews,
      auth,
      chatText,
      sellerID,
      postChat,
      toUserProfile,
      purchaseItem,
      favoriteItem,
    }
  }
}

</script>