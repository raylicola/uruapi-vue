<template>
  <div class="text-h5">
    新規登録
  </div>
  <small-space />
  <base-text-field label="ユーザーネーム" placeholder="やまだ" v-model="username" />
  <base-text-field label="メールアドレス" placeholder="yamada@example.com" v-model="email"/>
  <base-text-field label="パスワード" type="password" v-model="password" />
  <base-text-field label="パスワード（確認）" type="password" v-model="confirmPassword"/>
  <base-button text="登録" @click="signup" />
</template>

<script>
import { createUserWithEmailAndPassword } from "firebase/auth";
import { doc, setDoc } from "firebase/firestore";
import { db, auth, FirebaseTimestamp } from '@/firebase'
import { BaseButton, BaseTextField, SmallSpace } from '@/components'
import { ref } from '@vue/reactivity'
import { useRouter } from 'vue-router'

export default {
  name: 'SignUp',
  components: {
    'base-button': BaseButton,
    'base-text-field': BaseTextField,
    'small-space': SmallSpace,
  },
  setup(){
    const username = ref('')
    const email = ref('')
    const password = ref('')
    const confirmPassword = ref('')
    const router = useRouter()

    const signup = async() => {
      if (
        username.value === '' ||
        email.value === '' ||
        password.value === '' ||
        confirmPassword.value === '') {
        alert('必須項目が未入力です');
        return false;
      }
      if (password.value !== confirmPassword.value) {
        alert('パスワードが一致しません');
        return false;
      }
      createUserWithEmailAndPassword(auth, email.value, password.value)
      .then((userCredential) => {
        const user = userCredential.user;
        if (user) {
          const uid = user.uid;
          const timestamp = FirebaseTimestamp.now();

          const userInitialData = {
            created_at: timestamp,
            updated_at: timestamp,
            email: email.value,
            username: username.value,
            uid: uid,
            icon_path: '',
            introduction: '',
          }
          setDoc(doc(db, "users", uid), userInitialData)
          .then(() => {
            router.push('/signup-completed')
          });
        }})
      .catch((error) => {
        console.log(error.code, error.message)
      });
    }

    return {
      username,
      email,
      password,
      confirmPassword,
      signup,
    }
  }
}

</script>