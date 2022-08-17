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
          :src="iconPath"
          @click="toUserProfile"
        />
        <div class="text-h7 text--primary mx-5">
          {{userName}}
        </div>
      </div>
      <div class="text-h6 text--primary">
        {{wishTitle}}
      </div>
      <div class="my-4"></div>
      <div class="text--primary" style="white-space: pre-line;">
        {{wishDetail}}
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
    const route = useRoute()
    const router = useRouter()
    const store = useStore()
    const wishID = route.params.wishID
    const wishTitle = ref('')
    const wishDetail = ref('')
    const wishUserID = ref('')
    const iconPath = ref('')
    const userName = ref('')
    const items = ref([])

    const auth = computed(() => store.state.auth)

    const toUserProfile = () => {
      router.push('/user/' + wishUserID.value)
    }

    const suggestItem = () => {
      router.push('/wish/' + wishID + '/suggest')
    }

    const getWishDetail = async () => {
      const url = 'wish/' + wishID
      const {data} = await axios.get(url)
      wishTitle.value = data.wish.Title
      wishDetail.value = data.wish.Detail
      wishUserID.value = data.wish.UserID
      items.value = data.items.filter(item => item.PurchaserID == '')

      const docRef = doc(db, "users", wishUserID.value);
      const docSnap = await getDoc(docRef);
      iconPath.value = docSnap.data().icon_path ? docSnap.data().icon_path : require('@/assets/default_icon.jpg')
      userName.value = docSnap.data().username
    }

    const toDetail = (item_id) => {
      router.push('/item/' + item_id)
    }

    onMounted(async () => {
      getWishDetail()
    })

    watch(
      wishTitle, () => wishTitle.value,
      wishDetail, () => wishDetail.value,
      wishUserID, () => wishUserID.value,
      items, () => items.value,
      iconPath, () => iconPath.value,
      userName, () => userName.value,
    )

    return {
      wishTitle,
      wishDetail,
      items,
      auth,
      iconPath,
      userName,
      toDetail,
      suggestItem,
      toUserProfile,
    }
  }
}

</script>