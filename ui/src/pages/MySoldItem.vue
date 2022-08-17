<template>
  <small-space />
  <navigation-drawer />
  <div v-if="items.length != 0">
    <div class="text-h5">
      販売履歴
    </div>
    <small-space />
    <v-row>
      <item-card
        v-for="item in items"
        :key="item.ID"
        :title="item.Title"
        :price="item.Price"
        :img="item.Img"
      />
    </v-row>
  </div>
  <div v-else>
    <div class="text-h5">
      販売履歴がありません
    </div>
  </div>
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
    const userID = computed(() => store.state.userID)

    const getSoldItem = async () => {
      const url = '/user/' + userID.value + '/sold'
      const {data} = await axios.get(url)
      items.value = data.items
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