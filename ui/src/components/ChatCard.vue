<template>
  <v-card
    class="mx-4 my-4"
    width="400"
    variant="outlined"
  >
    <v-card-text>
      <div
        v-if="chat.UserID == seller_id"
        class="
          text--primary
          px-1
          mb-3
          bg-yellow-darken-2
          d-inline-block
        "
      >
        出品者
      </div>
      <div class="text--primary">
        {{chat.Content}}
      </div>
    </v-card-text>
  </v-card>
</template>

<script>

export default {
  name: 'ChatCard',
  props: {
    chat: { type: Object },
    seller_id: { type: String },
  },
}

</script>
