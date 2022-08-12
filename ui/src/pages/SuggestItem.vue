<template>
  <div class="text-h5">
    提案する商品を選択
  </div>
  <small-space />
  <v-row>
    <item-card
      v-for="item in items"
      :key="item.ID"
      :title="item.Title"
      :price="item.Price"
      :img="item.Img"
      @click="select(item)"
    />
  </v-row>
</template>

<script>
import { ItemCard, SmallSpace } from '@/components'
import axios from 'axios'
import { onMounted, watch, ref, computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useStore } from 'vuex';

export default {
  name: 'SuggestItem',
  components: {
    'item-card': ItemCard,
    'small-space': SmallSpace,
  },
  setup(){
    const router = useRouter()
    const route = useRoute()
    const store = useStore()
    const wish_id = route.params.wish_id

    const items = ref([])
    const user_id = computed(() => store.state.user_id)

    const getMyItem = async () => {
      const url = '/user/' + user_id.value + '/item'
      const {data} = await axios.get(url)
      items.value = data.items.filter(item => item.IsSold == 0)
    }

    const select = async(item) => {
      if(confirm(item.Title + 'を提案しますか？')) {
        try {
          const url = '/sale/create'
          const params = new URLSearchParams()
          params.append('wish_id', wish_id)
          params.append('item_id', item.ID)
          await axios.post(url, params)
          router.push('/wish/' + wish_id)
        } catch (e) {
          console.log(e)
        }
      }
    }

    onMounted(async () => {
      getMyItem()
    })

    watch(
      items, () => items.value,
    )

    return {
      items,
      select,
    }
  }
}

</script>