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
  />
</v-row>
</template>

<script>
import { ItemCard, SmallSpace, NavigationDrawer } from '@/components'
import axios from 'axios'
import { onMounted, watch, ref, computed } from 'vue';
import { useStore } from 'vuex';

export default {
  name: 'MySoldItem',
  components: {
    'item-card': ItemCard,
    'small-space': SmallSpace,
    'navigation-drawer': NavigationDrawer,
  },
  setup(){
    const store = useStore()

    const items = ref([])
    const user_id = computed(() => store.state.user_id)

    const getSoldItem = async () => {
      const url = '/user/' + user_id.value + '/sold'
      const {data} = await axios.get(url)
      console.log(data)
      items.value = data.Items
    }

    onMounted(async () => {
      getSoldItem()
    })

    watch(
      items, () => items.value,
    )

    return {
      items,
    }
  }
}

</script>