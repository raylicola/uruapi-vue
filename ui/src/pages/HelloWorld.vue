<template>

  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <v-img
          :src="require('../assets/logo.svg')"
          class="my-3"
          contain
          height="200"
        />
      </v-col>
    </v-row>
    <p>auth: {{auth}}</p>
    <BaseTextField />
    <BaseFileInput />
    <BaseTable />
    <WishCard />
    <ImagePreview />
    <BasePagination />
  </v-container>

</template>

<script>
import { computed } from 'vue'
import { useStore } from 'vuex'
import {
  BaseTextField,
  BasePagination,
  BaseTable,
  WishCard,
  BaseFileInput,
  ImagePreview
} from '@/components'

export default {
  name: 'HelloWorld',

  components: {
    BaseTextField,
    BasePagination,
    BaseTable,
    WishCard,
    BaseFileInput,
    ImagePreview,
  },
  setup () {
    const store = useStore()
    const auth = computed(() => store.state.auth)
    return {
      auth
    }
  }

}
</script>
