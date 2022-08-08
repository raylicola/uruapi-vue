<template>
<small-space />
<navigation-drawer />
<div >
  <v-row>
    <wish-card
      v-for="wish in wishes"
      :key="wish.ID"
      :title="wish.Title"
      :detail="wish.Detail"
      @click="toDetail(wish.ID)"
    />
  </v-row>
</div>
</template>

<script>
import { WishCard, SmallSpace, NavigationDrawer } from '@/components'
import axios from 'axios'
import { onMounted, watch, ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';

export default {
  name: 'MyWish',
  components: {
    'wish-card': WishCard,
    'small-space': SmallSpace,
    'navigation-drawer': NavigationDrawer,
  },
  setup(){
    const router = useRouter()
    const store = useStore()


    const wishes = ref([])
    const user_id = computed(() => store.state.user_id)

    const getMyWish = async () => {
      const url = '/user/' + user_id.value + '/wish'
      const {data} = await axios.get(url)
      wishes.value = data.wishes
      console.log(data)
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
      wishes,
      toDetail,
    }
  }
}

</script>