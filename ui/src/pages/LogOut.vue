<template>
  <p>ログアウトしました</p>
  <small-space />
  <v-btn href="/">トップページへ戻る</v-btn>
</template>

<script>
import { SmallSpace } from '@/components'

export default {
  name: 'SignUp',
  components: {
    'small-space': SmallSpace,
  },
}

</script>
    SmallSpace