import { createStore } from 'vuex'
import createPersistedState from 'vuex-persistedstate'

// 新しいストアインスタンスを作成
const store = createStore({
  state: {
    auth: false,
    user_id: '',
    user_name: '',
  },
  mutations: {
    setAuth: (state, auth) => state.auth = auth,
    setUserID: (state, user_id) => state.user_id = user_id,
    setUserName: (state, user_name) => state.user_name = user_name,
  },
  plugins: [createPersistedState(
    { // ストレージのキーを指定
      key: 'uriapi',
      // ストレージの種類を指定
      storage: window.sessionStorage
    })]
})

export default store