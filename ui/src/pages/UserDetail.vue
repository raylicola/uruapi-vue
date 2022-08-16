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
  </v-row>
  <small-space />
  <div>
    {{introduction}}
  </div>
  <small-space />
  <v-divider></v-divider>
  <small-space />
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
</template>

<script>
import {
  SmallSpace,
  BaseAvatar,
  ItemCard,
} from '@/components'
import { computed, onMounted, ref, watch } from 'vue'
import { useStore } from 'vuex'
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
  },
  setup() {
    const store = useStore()
    const route = useRoute()
    const router = useRouter()

    const username = computed(() => store.state.user_name)
    const icon_path = ref('')
    const introduction = ref('')
    const items = ref([])

    const editProfile = () => {
      router.push('/mypage/profile/edit')
    }

    const getUserProfile = async () => {
      const docRef = doc(db, "users", route.params.user_id);
      const docSnap = await getDoc(docRef);
      icon_path.value = docSnap.data().icon_path ? docSnap.data().icon_path : require('@/assets/default_icon.jpg')
      introduction.value = docSnap.data().introduction

      const url = 'seller/' + route.params.user_id
      const {data} = await axios.get(url)
      items.value = data.items.filter(item => item.PurchaserID == '')
    }

    const toDetail = (item_id) => {
      router.push('/item/' + item_id)
    }

    onMounted(async () => {
      getUserProfile()
    })

    watch(
      introduction, () => introduction.value,
      icon_path, () => icon_path.value,
      items, () => items.value,
    )

    return {
      username,
      introduction,
      icon_path,
      items,
      editProfile,
      toDetail,
    }
  }
}
</script>