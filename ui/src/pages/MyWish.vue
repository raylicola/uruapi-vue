<template>
  <small-space />
  <navigation-drawer />
  <div v-if="wishes.length != 0">
    <div class="text-h5">
      欲しいものリスト
    </div>
    <small-space />
    <v-row>
      <my-wish-card
        v-for="wish in wishes"
        :key="wish.ID"
        :title="wish.Title"
        :detail="wish.Detail"
        :id="wish.ID"
      />
    </v-row>
  </div>
  <div v-else>
    <div class="text-h5">
      欲しいものが投稿されていません
    </div>
  </div>
</template>

<script>
import { MyWishCard, SmallSpace, NavigationDrawer } from '@/components'
import axios from 'axios'
import { onMounted, watch, ref, computed } from 'vue';
import { useStore } from 'vuex';

export default {
  name: 'MyWish',
  components: {
    'my-wish-card': MyWishCard,
    'small-space': SmallSpace,
    'navigation-drawer': NavigationDrawer,
  },
  setup(){
    const store = useStore()
    const wishes = ref([])
    const userID = computed(() => store.state.userID)

    const getMyWish = async () => {
      const url = '/user/' + userID.value + '/wish'
      const {data} = await axios.get(url)
      wishes.value = data.wishes
    }

    onMounted(async () => {
      getMyWish()
    })

    watch(
      wishes, () => wishes.value,
    )

    return {
      wishes,
    }
  }
}

</script>