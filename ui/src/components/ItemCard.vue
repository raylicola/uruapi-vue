<template>
  <v-card
    class="mx-4 my-4"
    width="300"
    height="300"
  >
    <v-img
      :src="img"
      height="200px"
      cover
    ></v-img>
    <v-card-text>
      <div class="text-h6 text--primary">
        {{abbreviateText(title, 9, '...')}}
      </div>
      <div class="my-1"></div>
      <div class="text--primary">
        {{price}}円
      </div>
    </v-card-text>
  </v-card>
</template>

<script>
import { abbreviateText } from '@/utils'
export default {
  name: 'ItemCard',
  props: {
    title: { type: String },
    price: { type: String },
    img: { type: String },
  },
  setup() {
    return {
      abbreviateText,
    }
  }
}

</script>
