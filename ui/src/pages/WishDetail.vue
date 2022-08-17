<template>
  <back-button />
  <small-space />
  <v-card
    class="mx-4 my-4"
    width="800"
  >
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
        {{wish_title}}
      </div>
      <div class="my-4"></div>
      <div class="text--primary" style="white-space: pre-line;">
        {{wish_detail}}
      </div>
    </v-card-text>
    <base-button
      v-if="auth"
      text="商品を提案する"
      class="ma-3"
      @click="suggestItem"
    />
  </v-card>
  <small-space />
  <div v-if="items.length != 0">
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
  </div>
  <div v-else>
    <div class="text-h5">
      提案されている商品はありません
    </div>
  </div>
</template>

<script>
import {
  SmallSpace,
  ItemCard,
  BaseButton,
  BaseAvatar,
  BackButton,
} from '@/components'
import axios from 'axios'
import { onMounted, watch, ref, computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useStore } from 'vuex';
import { doc, getDoc } from "firebase/firestore";
import { db } from '@/firebase'

export default {
  name: 'WishDetail',
  components: {
    'small-space': SmallSpace,
    'item-card': ItemCard,
    'base-button': BaseButton,
    'base-avatar': BaseAvatar,
    'back-button': BackButton,
  },
  setup(){

    const wish_title = ref('')
    const wish_detail = ref('')
    const wish_user_id = ref('')
    const icon_path = ref('')
    const username = ref('')
    const items = ref([])
    const route = useRoute()
    const router = useRouter()
    const store = useStore()

    const auth = computed(() => store.state.auth)

    const toUserProfile = () => {
      router.push('/user/' + wish_user_id.value)
    }

    const suggestItem = () => {
      router.push('/wish/'+route.params.wish_id+'/suggest')
    }

    const getWishDetail = async () => {
      const url = 'wish/' + route.params.wish_id
      const {data} = await axios.get(url)
      wish_title.value = data.wish.Title
      wish_detail.value = data.wish.Detail
      wish_user_id.value = data.wish.UserID
      items.value = data.items.filter(item => item.PurchaserID == '')

      const docRef = doc(db, "users", wish_user_id.value);
      const docSnap = await getDoc(docRef);
      icon_path.value = docSnap.data().icon_path ? docSnap.data().icon_path : require('@/assets/default_icon.jpg')
      username.value = docSnap.data().username
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
      wish_user_id, () => wish_user_id.value,
      items, () => items.value,
      icon_path, () => icon_path.value,
      username, () => username.value,
    )

    return {
      wish_title,
      wish_detail,
      items,
      auth,
      icon_path,
      username,
      toDetail,
      suggestItem,
      toUserProfile,
    }
  }
}

</script>