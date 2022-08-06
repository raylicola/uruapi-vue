<template>
  <v-text-field
    v-model="message2"
    :label="label"
    :placeholder="placeholder"
    :type="type"
    variant="solo"
    clearable
    clear-icon="mdi-delete"
    class="w-50"
  ></v-text-field>
</template>

<script>

export default {
  name: 'BaseTextField',
  props: {
    label: { type: String },
    placeholder: { type: String },
    type: { type: String }
  }
}

</script>
