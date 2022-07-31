<template>
  <v-app>
    <TheHeader />
    <v-main>
      <v-container fluid>
      <router-view></router-view>
    </v-container>
    </v-main>
    <TheFooter />
  </v-app>
</template>

<script>
import {TheHeader, TheFooter} from '@/components'

export default {
  name: 'App',

  components: {
    TheHeader,
    TheFooter
  },
}
</script>
