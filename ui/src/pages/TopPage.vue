<template>
<v-form @submit.prevent>
  <base-text-field
  @keyup.enter="searchWish"
  label="キーワードで検索"
  v-model="keyword"
/>
</v-form>
<small-space />
<div >
  <v-row>
    <wish-card
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
import { WishCard, SmallSpace, BaseTextField } from '@/components'
import axios from 'axios'
import { onMounted, watch, ref } from 'vue';
import { useRouter } from 'vue-router';

export default {
  name: 'TopPage',
  components: {
    'wish-card': WishCard,
    'small-space': SmallSpace,
    'base-text-field': BaseTextField,
  },
  setup(){

    const wishes = ref([])
    const keyword = ref('')

    const router = useRouter()

    const getWish = async () => {
      const url = 'wish'
      const {data} = await axios.get(url)
      wishes.value = data.wishes
    }

    const toDetail = (wish_id) => {
      router.push('/wish/' + wish_id)
    }

    const searchWish = async () => {
      const url = '/wish/search?keyword=' + keyword.value
      const {data} = await axios.get(url)
      wishes.value = data.wishes
    }

    onMounted(async () => {
      getWish()
    })

    watch(
      wishes, () => wishes.value,
    )

    return {
      wishes,
      keyword,
      toDetail,
      searchWish
    }
  }
}

</script>