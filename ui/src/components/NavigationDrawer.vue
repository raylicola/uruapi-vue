<template>
<v-card>
  <v-navigation-drawer >
    <v-list density="compact" nav>
      <v-list-item title="プロフィール" value="createWish" @click="mypage" />
      <v-list-item title="欲しいものを投稿する" value="createWish" @click="createWish" />
      <v-list-item title="商品を出品する" value="createItem" @click="createItem" />
      <v-list-item title="欲しいものリスト" value="wish" @click="wish" />
      <v-list-item title="出品リスト" value="item" @click="item" />
      <v-list-item title="購入した商品" value="purchased" @click="purchasedItem" />
      <v-list-item title="売上履歴" value="sold" @click="soldItem" />
      <v-list-item title="ログアウト" value="logout" @click="logout" />
    </v-list>
  </v-navigation-drawer>
</v-card>

</template>

<script>
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'

export default {
  name: 'NavigationDrawer',
  setup() {
    const store = useStore()
    const router = useRouter()

    const mypage = () => {
      router.push('/mypage')
    }
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
    const createWish = () => {
      router.push('/mypage/wish/create')
    }
    const createItem = () => {
      router.push('/mypage/item/create')
    }
    const logout = () => {
      store.commit('setAuth', false)
      store.commit('setUserID', '')
      store.commit('setUserName', '')
      router.push('/logout')
    }

    return {
      wish,
      item,
      soldItem,
      purchasedItem,
      createItem,
      createWish,
      logout,
      mypage,
    }
  }
}

</script>