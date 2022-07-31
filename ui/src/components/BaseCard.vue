<template>
  <v-card
    class="mx-auto"
    max-width="300"
  >
    <v-card-text>
      <p class="text-h4 text--primary">
        title
      </p>
      <div class="text--primary">
        detail
      </div>
    </v-card-text>
    <v-card-actions>
      <v-btn
        variant="text"
        color="teal-accent-4"
        @click="reveal = true"
      >
        Learn More
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>

export default {
  name: 'BaseCard',
}

</script>
