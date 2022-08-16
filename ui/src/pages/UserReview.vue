<template>
  <small-space />
  <v-row align="center">
    <v-col cols="2">
      <base-avatar
        size="100"
        :src="icon_path"
      />
    </v-col>
    <v-col cols="5">
      <div class="text-h5">
        {{username}}
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
  },
  setup() {
    const route = useRoute()
    const router = useRouter()

    const icon_path = ref('')
    const introduction = ref('')
    const username = ref('')
    const reviews = ref([])

    const user_id = route.params.user_id

    const editProfile = () => {
      router.push('/mypage/profile/edit')
    }

    const getUserProfile = async () => {
      const docRef = doc(db, "users", user_id);
      const docSnap = await getDoc(docRef);
      icon_path.value = docSnap.data().icon_path ? docSnap.data().icon_path : require('@/assets/default_icon.jpg')
      introduction.value = docSnap.data().introduction
      username.value = docSnap.data().username

      const url = '/seller/review/' + user_id
      const {data} = await axios.get(url)
      reviews.value = data.reviews
    }

    const toDetail = (item_id) => {
      router.push('/item/' + item_id)
    }

    const toItemPage = () => {
      router.push('/user/' + user_id)
    }

    onMounted(async () => {
      getUserProfile()
    })

    watch(
      introduction, () => introduction.value,
      icon_path, () => icon_path.value,
      reviews, () => reviews.value,
    )

    return {
      username,
      introduction,
      icon_path,
      reviews,
      editProfile,
      toDetail,
      toItemPage,
    }
  }
}
</script>