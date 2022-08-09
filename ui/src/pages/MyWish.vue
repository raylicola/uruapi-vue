<template>
<small-space />
<navigation-drawer />
<v-row>
  <wish-card
    v-for="wish in wishes"
    :key="wish.ID"
    :title="wish.Title"
    :detail="wish.Detail"
    @click="toEdit(wish.ID)"
  />
</v-row>
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
    }

    const toEdit = (wish_id) => {
      router.push('/mypage/wish/edit/' + wish_id)
    }

    onMounted(async () => {
      getMyWish()
    })

    watch(
      wishes, () => wishes.value,
    )

    return {
      wishes,
      toEdit,
    }
  }
}

</script>