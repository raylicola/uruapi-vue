<template>
<small-space />
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
import { ItemCard, SmallSpace } from '@/components'
import axios from 'axios'
import { onMounted, watch, ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';

export default {
  name: 'MyPage',
  components: {
    'item-card': ItemCard,
    'small-space': SmallSpace
  },
  setup(){
    const router = useRouter()
    const store = useStore()


    const items = ref([])
    const user_id = computed(() => store.state.user_id)

    const getMyPost = async () => {
      const url = '/mypost/' + user_id.value
      const {data} = await axios.get(url)
      items.value = data.items
      console.log(data)
    }

    const toDetail = (item_id) => {
      router.push('/item/' + item_id)
    }

    onMounted(async () => {
      getMyPost()
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