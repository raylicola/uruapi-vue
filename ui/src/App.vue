<template>
  <v-app>
    <the-header />
    <v-main>
      <v-container>
        <router-view></router-view>
      </v-container>
    </v-main>
    <the-footer />
  </v-app>
</template>

<script>
import {
  TheHeader,
  TheFooter,
} from '@/components'

export default {
  name: 'App',

  components: {
    'the-header': TheHeader,
    'the-footer': TheFooter,
  },
}
</script>
