<template>
  <small-space />
  <navigation-drawer />
  <div v-if="items.length != 0">
    <div class="text-h5">
      購入履歴
    </div>
    <small-space />
    <v-row>
      <item-card
        v-for="item in items"
        :key="item.ID"
        :title="item.Title"
        :price="item.Price"
        :img="item.Img"
        @click="toReviewPage(item.ID)"
      />
    </v-row>
  </div>
  <div v-else>
    <div class="text-h5">
      購入履歴がありません
    </div>
  </div>
</template>

<script>
import { ItemCard, SmallSpace, NavigationDrawer } from '@/components'
import axios from 'axios'
import { onMounted, watch, ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';

export default {
  name: 'MyPurchasedItem',
  components: {
    'item-card': ItemCard,
    'small-space': SmallSpace,
    'navigation-drawer': NavigationDrawer,
  },
  setup(){
    const router = useRouter()
    const store = useStore()


    const items = ref([])
    const userID = computed(() => store.state.userID)

    const getPurchasedItem = async () => {
      const url = '/user/' + userID.value + '/purchased'
      const {data} = await axios.get(url)
      items.value = data.items
    }

    const toReviewPage = (itemID) => {
      router.push('/mypage/review/' + itemID)
    }

    onMounted(async () => {
      getPurchasedItem()
    })

    watch(
      items, () => items.value,
    )

    return {
      items,
      toReviewPage,
    }
  }
}

</script>