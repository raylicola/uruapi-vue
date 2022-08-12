<template>
  <v-card
    class="mx-4 my-4"
    width="300"
    height="350"
  >
    <v-img
      :src="img"
      height="200px"
      cover
    ></v-img>
    <v-card-text>
      <div class="text-h6 text--primary">
        {{title}}
      </div>
      <div class="mt-1"></div>
      <div class="text--primary">
        {{price}}円
      </div>
    </v-card-text>
    <v-card-actions>
      <v-btn color="orange" @click="toDetail(id)">
        詳細
      </v-btn>
      <v-btn color="orange" @click="toEdit(id)">
        編集
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import { useRouter } from 'vue-router'

export default {
  name: 'WishCard',
  props: {
    title: { type: String },
    price: { type: String },
    img: { type: String },
    id: { type: String }
  },
  setup() {
    const router = useRouter()

    const toDetail = (item_id) => {
      router.push('/item/' + item_id)
    }

    const toEdit = (item_id) => {
      router.push('/mypage/item/edit/' + item_id)
    }

    return {
      toDetail,
      toEdit,
    }
  }
}

</script>
