import { createStore } from 'vuex'
import createPersistedState from 'vuex-persistedstate'

// 新しいストアインスタンスを作成
const store = createStore({
  state: {
    auth: false,
    userID: '',
    userName: '',
  },
  mutations: {
    setAuth: (state, auth) => state.auth = auth,
    setUserID: (state, userID) => state.userID = userID,
    setUserName: (state, userName) => state.userName = userName,
  },
  plugins: [createPersistedState(
    { // ストレージのキーを指定
      key: 'uriapi',
      // ストレージの種類を指定
      storage: window.sessionStorage
    })]
})

export default store