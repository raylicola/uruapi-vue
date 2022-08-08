<template>
  <v-app-bar
    color="primary"
    density="compact"
  >

    <v-app-bar-title>UruApi</v-app-bar-title>

    <v-menu offset-y>
      <template v-slot:activator="{ props }">
        <v-btn icon="mdi-dots-vertical" v-bind="props"></v-btn>
      </template>
      <v-list>
        <v-list-item
          v-for="(item, index) in (auth? auth_menu:not_auth_menu)"
          :key="index"
          :value="index"
        >
          <v-list-item-content @click="item.clicked">
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-menu>
  </v-app-bar>
</template>

<script>
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'


export default {
  name: 'TheHeader',
  setup() {
    const router = useRouter()
    const store = useStore()
    const auth = computed(() => store.state.auth)
    const auth_menu = [
      {
        title: 'ホーム',
        clicked: () => {
          router.push('/')
        }
      },
      {
        title: 'マイページ',
        clicked: () => {
          router.push('/mypage')
        }
      },
      {
        title: 'ログアウト',
        clicked: () => {
          store.commit('setAuth', false)
          store.commit('setUserID', '')
          store.commit('setUserName', '')
          router.push('/logout')
        }
      }
    ]

    const not_auth_menu = [
      {
        title: 'ホーム',
        clicked: () => {
          router.push('/')
        }
      },
      {
        title: 'ログイン',
        clicked: () => {
          router.push('/signin')
        }
      },
      {
        title: '新規会員登録',
        clicked: () => {
          router.push('/signup')
        }
      }
    ]

    return {
      auth_menu,
      not_auth_menu,
      auth
    }
  }
}

</script>