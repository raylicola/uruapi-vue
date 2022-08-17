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
        text="商品を見る"
        @click="toItemPage"
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
  <div v-if="reviews.length != 0">
    <div class="text-body-1">
      出品者へのレビュー
    </div>
    <small-space />
    <review-card
      v-for="review in reviews"
      :key="review.ID"
      :review="review"
    />
    <small-space />
    <v-row>
    </v-row>
  </div>
  <div v-else>
    <div class="text-h5">
      レビューがありません
    </div>
  </div>
</template>

<script>
import {
  SmallSpace,
  BaseAvatar,
  BaseButton,
  ReviewCard,
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
    'base-button': BaseButton,
    'review-card': ReviewCard,
    'back-button': BackButton,
  },
  setup() {
    const route = useRoute()
    const router = useRouter()

    const iconPath = ref('')
    const introduction = ref('')
    const userName = ref('')
    const reviews = ref([])

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

      const url = '/seller/review/' + userID
      const {data} = await axios.get(url)
      reviews.value = data.reviews
      console.log()
    }

    const toDetail = (item_id) => {
      router.push('/item/' + item_id)
    }

    const toItemPage = () => {
      router.push('/user/' + userID)
    }

    onMounted(async () => {
      getUserProfile()
    })

    watch(
      introduction, () => introduction.value,
      iconPath, () => iconPath.value,
      reviews, () => reviews.value,
    )

    return {
      userName,
      introduction,
      iconPath,
      reviews,
      editProfile,
      toDetail,
      toItemPage,
    }
  }
}
</script>