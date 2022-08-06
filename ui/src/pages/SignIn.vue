<template>
  <div class="text-h5 py-8">
    ログイン
  </div>
  <base-text-field label="メールアドレス" placeholder="yamada@example.com"/>
  <base-text-field label="パスワード" type="password" />
  <base-button text="ログイン" />
</template>

<script>
import { BaseButton, BaseTextField } from '@/components'

export default {
  name: 'SignIn',
  components: {
    'base-button': BaseButton,
    'base-text-field': BaseTextField
  }
}

</script>