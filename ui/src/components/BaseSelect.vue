<template>
  <v-select
    :items="items"
    :label="label"
    variant="underlined"
  ></v-select>
</template>

<script>

export default {
  name: 'BaseSelect',
  props: {
    label: { type: String },
    items: { type: Array },
  }
}

</script>