<template>
  <div class="text-h5 py-8">
    ログイン
  </div>
  <base-text-field label="メールアドレス" placeholder="yamada@example.com" v-model="email"/>
  <base-text-field label="パスワード" type="password" v-model="password" />
  <base-button text="ログイン" @click="signin" />
</template>

<script>
import { signInWithEmailAndPassword  } from "firebase/auth";
import { doc, getDoc } from "firebase/firestore";
import { db, auth } from '@/firebase'
import { BaseButton, BaseTextField } from '@/components'
import { ref } from '@vue/reactivity'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex';

export default {
  name: 'SignIn',
  components: {
    'base-button': BaseButton,
    'base-text-field': BaseTextField
  },
  setup(){
    const email = ref('')
    const password = ref('')
    const router = useRouter()
    const store = useStore()

    const signin = () => {
      if (
        email.value === '' ||
        password.value === '') {
        alert('必須項目が未入力です');
        return false;
      }

      signInWithEmailAndPassword(auth, email.value, password.value)
      .then((userCredential) => {
        const user = userCredential.user;
        if (user) {
          const uid = user.uid
          getDoc(doc(db, "users", uid))
          .then((userCredential) => {
            const data = userCredential.data();
            if (data){
              store.commit('setAuth', true)
              store.commit('setUserID', data.uid)
              store.commit('setUserName', data.username)
            }
          })
          .catch((e) => {
            console.log(e);
            return false;
          });
          router.push('/')
        }
      })
      .catch((error) => {
        console.log(error.code, error.message)
      });
    }

    return {
      email,
      password,
      signin,
    }
  }
}

</script>