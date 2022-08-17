<template>
  <back-button />
  <small-space />
  <v-row align="center">
    <v-col cols="2">
      <base-avatar
        size="100"
        :src="iconPath"
      />
    </v-col>
    <v-col cols="5">
      <div class="text-h5">
        {{userName}}
      </div>
    </v-col>
    <v-col cols="2">
      <base-button
        text="レビューを見る"
        @click="toReviewPage"
      />
    </v-col>
  </v-row>
  <small-space />
  <div style="white-space: pre-line;">
    {{introduction}}
  </div>
  <small-space />
  <v-divider></v-divider>
  <small-space />
  <div v-if="items.length != 0">
    <div class="text-h5">
      出品している商品
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
      出品している商品はありません
    </div>
  </div>
</template>

<script>
import {
  SmallSpace,
  BaseAvatar,
  ItemCard,
  BaseButton,
  BackButton,
} from '@/components'
import { onMounted, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { db } from '@/firebase'
import { doc, getDoc } from "firebase/firestore";
import axios from 'axios'

export default {
  name: 'MyPage',
  components: {
    'small-space': SmallSpace,
    'base-avatar': BaseAvatar,
    'item-card': ItemCard,
    'base-button': BaseButton,
    'back-button': BackButton,
  },
  setup() {
    const route = useRoute()
    const router = useRouter()

    const iconPath = ref('')
    const introduction = ref('')
    const userName = ref('')
    const items = ref([])

    const userID = route.params.userID

    const editProfile = () => {
      router.push('/mypage/profile/edit')
    }

    const getUserProfile = async () => {
      const docRef = doc(db, "users", userID);
      const docSnap = await getDoc(docRef);
      iconPath.value = docSnap.data().icon_path ? docSnap.data().icon_path : require('@/assets/default_icon.jpg')
      introduction.value = docSnap.data().introduction
      userName.value = docSnap.data().username

      const url = 'seller/' + userID
      const {data} = await axios.get(url)
      items.value = data.items.filter(item => item.PurchaserID == '')
    }

    const toDetail = (itemID) => {
      router.push('/item/' + itemID)
    }

    const toReviewPage = () => {
      router.push('/user/' + userID + '/review')
    }

    onMounted(async () => {
      getUserProfile()
    })

    watch(
      introduction, () => introduction.value,
      iconPath, () => iconPath.value,
      items, () => items.value,
    )

    return {
      userName,
      introduction,
      iconPath,
      items,
      editProfile,
      toDetail,
      toReviewPage,
    }
  }
}
</script>