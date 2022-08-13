<template>
  <small-space />
  <v-card
    class="mx-4 my-4"
    width="1000"
  >
    <v-row no-gutters>
      <v-col cols="4">
        <v-img
          :src="img"
          cover
        ></v-img>
      </v-col>
      <v-col cols="8">
        <v-card-text>
          <div class="text-h6 text--primary">
            {{title}}
          </div>
          <div>
            <base-button
              v-if="auth"
              text="商品を購入する"
              class="ma-3"
              @click="purchase"
            />
          </div>
          <div class="my-4"></div>
          <div class="text-h7 mb-2 text--primary">
            商品の説明
          </div>
          <div class="text--primary" style="white-space: pre-line;">
            {{detail}}
          </div>
        </v-card-text>
      </v-col>
    </v-row>
  </v-card>
</template>

<script>
import { computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import {
  BaseButton,
  SmallSpace,
} from '.'
import { useStore } from 'vuex'

export default {
  name: 'ItemDetailCard',
  components: {
    'base-button': BaseButton,
    'small-space': SmallSpace,
  },
  props: {
    img: { type: String },
    title: { type: String },
    detail: { type: String },
  },
  setup() {
    const router = useRouter()
    const route = useRoute()
    const store = useStore()
    const auth = computed(() => store.state.auth)

    const purchase = () => {
      router.push('/purchase/' + route.params.item_id)
    }
    return {
      purchase,
      auth,
    }
  }
}

</script>
