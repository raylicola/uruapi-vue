<template>
  <v-btn
    color="primary"
  >
    {{text}}
  </v-btn>
</template>

<script>

export default {
  name: 'BaseButton',
  props: {
    text: { type: String },
  },
}

</script>