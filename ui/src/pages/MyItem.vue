<template>
  <small-space />
  <navigation-drawer />
  <div v-if="items.length != 0">
    <div class="text-h5">
      出品している商品
    </div>
    <small-space />
    <v-row>
    <my-item-card
      v-for="item in items"
      :key="item.ID"
      :title="item.Title"
      :price="item.Price"
      :img="item.Img"
      :id="item.ID"
    />
  </v-row>
  </div>
  <div v-else>
    <div class="text-h5">
      出品している商品がありません
    </div>
  </div>
</template>

<script>
import { MyItemCard, SmallSpace, NavigationDrawer } from '@/components'
import axios from 'axios'
import { onMounted, watch, ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';

export default {
  name: 'MyItem',
  components: {
    'my-item-card': MyItemCard,
    'small-space': SmallSpace,
    'navigation-drawer': NavigationDrawer,
  },
  setup(){
    const router = useRouter()
    const store = useStore()


    const items = ref([])
    const userID = computed(() => store.state.userID)

    const getMyItem = async () => {
      const url = '/user/' + userID.value + '/item'
      const {data} = await axios.get(url)
      items.value = data.items.filter(item => item.PurchaserID == '')
    }

    const toEdit = (itemID) => {
      router.push('/mypage/item/edit/' + itemID)
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