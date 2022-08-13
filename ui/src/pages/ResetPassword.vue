<template>
  <small-space />
  <div class="text-h5">
    パスワードを変更する
  </div>
  <small-space />
  <v-form @submit.prevent>
    <base-text-field label="メールアドレス" v-model="email"/>
    <base-button
      text="送信する"
      @click="resetPassword"
    />
  </v-form>
</template>

<script>
import { sendResetPasswordEmail } from 'firebase/auth'
import { auth } from '@/firebase'
import { BaseButton, BaseTextField, SmallSpace } from '@/components'
import { ref } from '@vue/reactivity'
import { useRouter } from 'vue-router'

export default {
  name: 'ResetPassword',
  components: {
    'base-text-field': BaseTextField,
    'base-button': BaseButton,
    'small-space': SmallSpace,
  },
  setup() {
    const router = useRouter()
    const email = ref('')
    const resetPassword = () => {
      sendResetPasswordEmail(auth, email.value)
      .then(() => {
        confirm('パスワード再設定用のメールを送信しました。')
        router.push('/signin')
      })
      .catch((error) => {
        confirm('メールアドレスが登録されていません。')
        console.log(error)
      });
    }
    return {
      email,
      resetPassword,
    }
  }
}

</script>