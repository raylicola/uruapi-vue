<template>
<small-space />
<navigation-drawer />
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
    const user_id = computed(() => store.state.user_id)

    const getPurchasedItem = async () => {
      const url = '/user/' + user_id.value + '/purchased'
      const {data} = await axios.get(url)
      items.value = data.Items
      console.log(items.value)
    }

    const toReviewPage = (item_id) => {
      router.push('/mypage/review/' + item_id)
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