<template>
  <div class="text-h5">
    欲しいもの
  </div>
  <small-space />
  <base-text-field label="タイトル" v-model="title"/>
  <base-text-area label="説明" v-model="detail" />
  <base-button text="更新" @click="editWish" class="mx-5"/>
  <base-button text="削除" @click="deleteWish" />
</template>

<script>
import {
  BaseButton,
  BaseTextField,
  BaseTextArea,
  SmallSpace,
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
  },
  setup(){
    const title = ref('')
    const detail = ref('')

    const router = useRouter()
    const store = useStore()
    const route = useRoute()
    const user_id = computed(() => store.state.user_id)

    const getWish = async () => {
      const url = 'wish/' + route.params.wish_id
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
        const url = '/wish/edit/' + route.params.wish_id
        const params = new URLSearchParams()
        params.append('title', title.value)
        params.append('detail', detail.value)
        params.append('user_id', user_id.value)
        await axios.put(url, params)
        router.push('/mypage')
      } catch (e) {
        console.log(e)
      }
    }

    const deleteWish = async () => {
      try {
        const url = '/wish/delete/' + route.params.wish_id
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