<template>
  <a @click="router.back()">&lt;</a>
</template>

<script>
import { useRouter } from 'vue-router'

export default {
  name: 'BackButton',
  setup() {
    const router = useRouter()
    return {
      router
    }
  },
}

</script>