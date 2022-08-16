<template>
  <v-file-input
    :label="label"
    filled
    prepend-icon="mdi-camera"
  ></v-file-input>
</template>

<script>

export default {
  name: 'BaseFileInput',
  props: {
    label: { type: String },
  },
}

</script>