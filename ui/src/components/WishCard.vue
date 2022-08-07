<template>
  <v-card
    class="mx-4 my-4"
    width="300"
    height="300"
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
  </v-card>
</template>

<script>

export default {
  name: 'WishCard',
  props: {
    title: { type: String },
    detail: { type: String },
  },
}

</script>
