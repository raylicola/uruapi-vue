<template>
  <small-space />
  <navigation-drawer />
  <div>{{username}}</div>

</template>

<script>
import { SmallSpace, NavigationDrawer } from '@/components'
import { computed } from 'vue'
import { useStore } from 'vuex'

export default {
  name: 'MyPage',
  components: {
    'small-space': SmallSpace,
    'navigation-drawer': NavigationDrawer,
  },
  setup() {
    const store = useStore()
    const username = computed(() => store.state.user_name)

    return {
      username
    }
  }
}
</script>