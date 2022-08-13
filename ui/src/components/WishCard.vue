<template>
  <v-card
    class="mx-4 my-4"
    width="300"
    height="300"
  >
    <v-card-text>
      <div class="text-h6 text--primary">
        {{abbreviatedTitle}}
      </div>
      <div class="my-4"></div>
      <div class="text--primary">
        {{abbreviatedDetail}}
      </div>
    </v-card-text>
  </v-card>
</template>

<script>
import { abbreviateText } from '@/utils'
import { ref } from '@vue/reactivity'

export default {
  name: 'WishCard',
  props: {
    title: { type: String },
    detail: { type: String },
  },
  setup(props) {
    const abbreviatedTitle = ref(abbreviateText(props.title, 9, '...'))
    const abbreviatedDetail = ref(abbreviateText(props.detail, 111, '...'))

    return {
      abbreviatedTitle,
      abbreviatedDetail
    }
  }
}

</script>
