<template>
<small-space />
<div >
  <v-row>
    <base-card
      v-for="wish in wishes"
      :key="wish.ID"
      :title="wish.Title"
      :detail="wish.Detail"
      @click="toDetail(wish.ID)"
    />
  </v-row>
</div>
</template>

<script>
import { BaseCard, SmallSpace } from '@/components'
import axios from 'axios'
import { onMounted, watch, ref } from 'vue';
import { useRouter } from 'vue-router';

export default {
  name: 'EditWish',
  components: {
    'base-card': BaseCard,
    'small-space': SmallSpace
  },
  setup(){

    const wishes = ref([])

    const router = useRouter()

    const getWish = async () => {
      const url = 'wish'
      const {data} = await axios.get(url)
      wishes.value = data.wishes
      console.log(wishes)
    }

    const toDetail = (wish_id) => {
      router.push('/wish/' + wish_id)
    }

    onMounted(async () => {
      getWish()
    })

    watch(
      wishes, () => wishes.value,
    )

    return {
      wishes,
      toDetail,
    }
  }
}

</script>