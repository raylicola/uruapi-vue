<template>
  <small-space />
  <navigation-drawer />
  <div>{{username}}</div>
  <v-row>
  <wish-card
    v-for="wish in wishes"
    :key="wish.ID"
    :title="wish.Title"
    :detail="wish.Detail"
    @click="toDetail(wish.ID)"
  />
</v-row>
</template>

<script>
import { SmallSpace, NavigationDrawer, WishCard } from '@/components'
import { computed, onMounted, ref, watch } from 'vue'
import { useStore } from 'vuex'
import axios from 'axios'
import { useRouter } from 'vue-router'

export default {
  name: 'MyPage',
  components: {
    'small-space': SmallSpace,
    'navigation-drawer': NavigationDrawer,
    'wish-card': WishCard,
  },
  setup() {
    const store = useStore()
    const router = useRouter()

    const wishes = ref([])
    const user_id = computed(() => store.state.user_id)
    const username = computed(() => store.state.user_name)

    const getMyWish = async () => {
      const url = '/user/' + user_id.value + '/wish'
      const {data} = await axios.get(url)
      wishes.value = data.wishes
      console.log(wishes.value)
    }

    const toDetail = (wish_id) => {
      router.push('/wish/' + wish_id)
    }

    onMounted(async () => {
      getMyWish()
    })

    watch(
      wishes, () => wishes.value,
    )

    return {
      username,
      wishes,
      toDetail,
    }
  }
}
</script>