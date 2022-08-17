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
  <base-button text="更新" @click="editWish" class="mx-5"/>
  <base-button text="削除" @click="deleteWish" />
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
import { useRoute, useRouter } from 'vue-router'
import { useStore } from 'vuex';
import axios from 'axios'
import { computed, onMounted, watch } from 'vue';

export default {
  name: 'EditWish',
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
    const route = useRoute()
    const userID = computed(() => store.state.userID)
    const wishID = route.params.wishID

    const getWish = async () => {
      const url = 'wish/' + wishID
      const {data} = await axios.get(url)
      title.value = data.wish.Title
      detail.value = data.wish.Detail
    }

    const editWish = async() => {
      if (
        title.value === '' ||
        detail.value === '') {
          alert('必須項目が未入力です')
          return false;
      }

      try {
        const url = '/wish/edit/' + wishID
        const params = new URLSearchParams()
        params.append('title', title.value)
        params.append('detail', detail.value)
        params.append('user_id', userID.value)
        await axios.put(url, params)
        router.push('/mypage/wish')
      } catch (e) {
        console.log(e)
      }
    }

    const deleteWish = async () => {
      try {
        const url = '/wish/delete/' + wishID
        await axios.delete(url)
        router.push('/mypage/wish')
      } catch (e) {
        console.log(e)
      }
    }

    onMounted(async () => {
      getWish()
    })

    watch(
      detail, () => detail.value,
      title, () => title.value,
    )

    return {
      title,
      detail,
      editWish,
      deleteWish,
    }
  }
}

</script>