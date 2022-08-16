<template>
  <v-avatar
    :size="size"
  >
    <v-img
      alt="Avatar"
      :src="src"
    />
  </v-avatar>
</template>

<script>
export default {
  name: 'BaseAvatar',
  props: {
    size: { type: Number },
    src: { type: String, default: require('@/assets/default_icon.jpg') }
  }
}

</script>