<template>
  <v-card
    class="mx-4 my-4"
    width="400"
    variant="outlined"
  >
    <v-card-text>
      <div class="text--primary d-flex">
        <p
          v-for="n in review.Rate"
          :key="n"
        >★</p>
      </div>
      <div class="pb-2"></div>
      <div class="text--primary">
        {{review.Detail}}
      </div>
    </v-card-text>
  </v-card>
</template>

<script>

export default {
  name: 'ChatCard',
  props: {
    review: { type: Object },
  },
}

</script>
