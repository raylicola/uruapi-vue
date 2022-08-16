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
    src: { type: String }
  }
}

</script>