<template>
  <v-card
    class="mx-4 my-4"
    width="800"
  >
    <v-card-text>
      <div class="text-h6 text--primary">
        {{title}}
      </div>
      <div class="my-4"></div>
      <div class="text--primary">
        {{detail}}
      </div>
    </v-card-text>
    <base-button
      text="商品を提案する"
      class="ma-3"
      @click="suggest"
      />
  </v-card>
</template>

<script>
import { useRouter, useRoute } from 'vue-router'
import { BaseButton } from '.'

export default {
  name: 'WishDetailCard',
  components: {
    'base-button': BaseButton,
  },
  props: {
    title: { type: String },
    detail: { type: String },
  },
  setup() {
    const router = useRouter()
    const route = useRoute()
    const suggest = () => {
      router.push('/wish/'+route.params.wish_id+'/suggest')
    }
    return {
      suggest,
    }
  }
}

</script>
