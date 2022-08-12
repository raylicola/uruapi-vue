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
    @click="toEdit(item.ID)"
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
  name: 'MyItem',
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

    const getMyItem = async () => {
      const url = '/user/' + user_id.value + '/item'
      const {data} = await axios.get(url)
      items.value = data.items.filter(item => item.IsSold == 0)
    }

    const toEdit = (item_id) => {
      router.push('/mypage/item/edit/' + item_id)
    }

    onMounted(async () => {
      getMyItem()
    })

    watch(
      items, () => items.value,
    )

    return {
      items,
      toEdit,
    }
  }
}

</script>