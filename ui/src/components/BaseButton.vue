<template>
  <v-btn
    color="primary"
    variant="outlined"
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