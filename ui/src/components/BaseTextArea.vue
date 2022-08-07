<template>
  <v-textarea
    :label="label"
    :placeholder="placeholder"
  ></v-textarea>
</template>

<script>

export default {
  name: 'BaseTextArea',
  props: {
    label: { type: String },
    placeholder: { type: String },
    type: { type: String }
  }
}

</script>