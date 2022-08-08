<template>
<small-space />
<navigation-drawer />
<div >
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
</template>

<script>
import { ItemCard, SmallSpace, NavigationDrawer } from '@/components'
import axios from 'axios'
import { onMounted, watch, ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';

export default {
  name: 'MySoldItem',
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

    const getSoldItem = async () => {
      const url = '/user/' + user_id.value + '/sold'
      const {data} = await axios.get(url)
      items.value = data.soldItems
      console.log(data)
    }

    const toDetail = (item_id) => {
      router.push('/item/' + item_id)
    }

    onMounted(async () => {
      getSoldItem()
    })

    watch(
      items, () => items.value,
    )

    return {
      items,
      toDetail,
    }
  }
}

</script>