<template>
  <back-button />
  <div class="text-h5">
    欲しいもの
  </div>
  <small-space />
  <v-row>
    <v-col
      cols="12"
      md="6"
    >
      <base-text-field label="タイトル" v-model="title"/>
      <base-text-area label="説明" v-model="detail" />
    </v-col>
  </v-row>
  <base-button text="投稿" @click="createWish" />
</template>

<script>
import {
  BaseButton,
  BaseTextField,
  BaseTextArea,
  SmallSpace,
  BackButton,
  } from '@/components'
import { ref } from '@vue/reactivity'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex';
import axios from 'axios'
import { computed } from 'vue';

export default {
  name: 'CreateWish',
  components: {
    'base-button': BaseButton,
    'base-text-field': BaseTextField,
    'base-text-area': BaseTextArea,
    'small-space': SmallSpace,
    'back-button': BackButton,
  },
  setup(){
    const title = ref('')
    const detail = ref('')
    const router = useRouter()
    const store = useStore()
    const userID = computed(() => store.state.userID)

    const createWish = async() => {
      if (
        title.value === '' ||
        detail.value === '') {
          alert('必須項目が未入力です')
          return false;
      }

      try {
        const url = '/wish/create'
        const params = new URLSearchParams()
        params.append('title', title.value)
        params.append('detail', detail.value)
        params.append('user_id', userID.value)
        await axios.post(url, params)
        router.push('/mypage/wish')
      } catch (e) {
        console.log(e)
      }
    }

    return {
      title,
      detail,
      createWish,
    }
  }
}

</script>