<template>
  <back-button />
  <div class="text-h5">
    出品
  </div>
  <small-space />
  <v-row>
    <v-col
      cols="12"
      md="6"
    >
      <base-text-field label="タイトル" v-model="title"/>
      <base-text-field label="価格（円）" v-model="price" />
      <base-text-area label="説明" v-model="detail" />
    </v-col>
  </v-row>
  <v-row>
    <v-col cols="3">
      <base-file-input
        @change="uploadImage"
        label="画像を選択"
      />
    </v-col>
    <v-col cols="1" />
    <v-col cols="4">
      <image-preview :src="itemPath" />
    </v-col>
  </v-row>
  <small-space />
  <base-button text="投稿" @click="createItem" />
</template>

<script>
import {
  BaseButton,
  BaseTextField,
  BaseTextArea,
  BaseFileInput,
  ImagePreview,
  SmallSpace,
  BackButton
  } from '@/components'
import { ref } from '@vue/reactivity'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { computed } from 'vue'
import axios from 'axios'
import UUID from 'uuidjs'
import { storage } from '@/firebase'
import { getDownloadURL, ref as fbRef, uploadBytes, deleteObject } from "firebase/storage";

export default {
  name: 'CreateItem',
  components: {
    'base-button': BaseButton,
    'base-text-field': BaseTextField,
    'base-text-area': BaseTextArea,
    'base-file-input': BaseFileInput,
    'image-preview': ImagePreview,
    'small-space': SmallSpace,
    'back-button': BackButton,
  },
  setup(){
    const title = ref('')
    const detail = ref('')
    const price = ref('')
    const itemPath = ref('')
    const itemFileName = ref('')
    const prevFileName = ref('')

    const router = useRouter()
    const store = useStore()
    const userID = computed(() => store.state.userID)

    const uploadImage = (e) => {
      const file = e.target.files[0]
      const uuid = UUID.generate();
      const extention = file.name.split('.').pop();
      const fileName = uuid + '.' + extention;
      itemFileName.value = fileName
      const storageRef = fbRef(storage, 'images/items/'+fileName);
      uploadBytes(storageRef, file).then(() => {
        getDownloadURL(fbRef(storage, 'images/items/'+fileName))
        .then((url) => {
          const newImage = {id: fileName, path: url};
          itemPath.value = newImage.path
        })
      });

      const desertRef = fbRef(storage, 'images/items/'+ prevFileName.value);

      if(prevFileName.value != '') {
        deleteObject(desertRef).then(() => {
        // File deleted successfully
        }).catch((error) => {
          console.log(error)
        });
      }
      prevFileName.value = fileName
    }

    const createItem = async() => {
      if (
        title.value === '' ||
        price.value === '' ||
        detail.value === '') {
          alert('必須項目が未入力です')
          return false
      }

      if(itemPath.value == '') {
        alert('商品の画像を選択してください')
        return false
      }

      try {
        const url = '/item/create'
        const params = new URLSearchParams()
        params.append('title', title.value)
        params.append('price', price.value)
        params.append('detail', detail.value)
        params.append('seller_id', userID.value)
        params.append('img', itemPath.value)
        params.append('file_name', itemFileName.value)
        await axios.post(url, params)
        router.push('/mypage/item')
      } catch (e) {
        console.log(e)
      }
    }

    return {
      title,
      price,
      detail,
      itemPath,
      createItem,
      uploadImage,
    }
  }
}

</script>