<template>
  <v-card
    class="mx-4 my-4"
    width="300"
    height="300"
  >
    <v-card-text>
      <div class="text-h6 text--primary">
        {{title}}
      </div>
      <div class="my-4"></div>
      <div class="text--primary">
        {{detail}}
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
    detail: { type: String },
    id: { type: String },
  },
  setup() {
    const router = useRouter()

    const toDetail = (wish_id) => {
      router.push('/wish/' + wish_id)
    }

    const toEdit = (wish_id) => {
      router.push('/mypage/wish/edit/' + wish_id)
    }

    return {
      toDetail,
      toEdit,
    }
  }
}

</script>
