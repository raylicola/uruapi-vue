<template>
  <v-btn
    color="primary"
    @click="signin"
  >
    Normal Button
  </v-btn>
</template>

<script>
import { useStore } from 'vuex'

export default {
  name: 'BaseButton',
  setup() {
    const store = useStore()

    const signin = () => {
      store.commit('setAuth', true)
    }
    return {
      signin,
    }
  }
}

</script>