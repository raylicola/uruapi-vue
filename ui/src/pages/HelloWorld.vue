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
    <BaseTextField />
    <BaseFileInput />
    <BaseTable />
    <BaseCard />
    <ImagePreview />
    <BasePagination />
  </v-container>

</template>

<script>

import {
  BaseTextField,
  BasePagination,
  BaseTable,
  BaseCard,
  BaseFileInput,
  ImagePreview
} from '@/components'

export default {
  name: 'HelloWorld',

  components: {
    BaseTextField,
    BasePagination,
    BaseTable,
    BaseCard,
    BaseFileInput,
    ImagePreview,
},

}
</script>
