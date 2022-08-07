<template>
  <v-img
    class="bg-white"
    width="300"
    aspect-ratio="1"
    :src="url"
    cover
  ></v-img>
</template>

<script>

export default {
  name: 'ImagePreview',
  props: {
    url: {
      type: String,
    },
  }
}

</script>