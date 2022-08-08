<template>
<v-card>
  <v-navigation-drawer >
    <v-list density="compact" nav>
      <v-list-item title="欲しいもの" value="wish" @click="wish" />
      <v-list-item title="出品した商品" value="item" @click="item" />
      <v-list-item title="購入した商品" value="purchased" @click="purchasedItem" />
      <v-list-item title="売上履歴" value="sold" @click="soldItem" />
    </v-list>
  </v-navigation-drawer>
</v-card>

</template>

<script>
import { useRouter } from 'vue-router'

export default {
  name: 'NavigationDrawer',
  setup() {
    const router = useRouter()

    const wish = () => {
      router.push('/mypage/wish')
    }
    const item = () => {
      router.push('/mypage/item')
    }
    const soldItem = () => {
      router.push('/mypage/sold')
    }
    const purchasedItem = () => {
      router.push('/mypage/purchased')
    }

    return {
      wish,
      item,
      soldItem,
      purchasedItem,
    }
  }
}

</script>