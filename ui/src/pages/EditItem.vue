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
        label="画像を変更"
      />
    </v-col>
    <v-col cols="1" />
    <v-col cols="4">
      <image-preview :src="itemPath" />
    </v-col>
  </v-row>
  <small-space />
  <base-button text="更新" @click="editItem" class="mx-5"/>
  <base-button text="削除" @click="deleteItem" />
</template>

<script>
import {
  BaseButton,
  BaseTextField,
  BaseTextArea,
  BaseFileInput,
  ImagePreview,
  SmallSpace,
  BackButton,
  } from '@/components'
import { ref } from '@vue/reactivity'
import { useRoute, useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { computed, onMounted, watch } from 'vue'
import axios from 'axios'
import UUID from 'uuidjs'
import { storage } from '@/firebase'
import { getDownloadURL, ref as fbRef, uploadBytes, deleteObject } from "firebase/storage";

export default {
  name: 'EditItem',
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
    const router = useRouter()
    const route = useRoute()
    const store = useStore()

    const title = ref('')
    const detail = ref('')
    const price = ref('')
    const itemPath = ref('')
    const itemFileName = ref('')
    const prevItemFileName = ref('')
    const userID = computed(() => store.state.userID)
    const itemID = route.params.itemID

    const getItem = async () => {
      const url = 'item/' + itemID
      const {data} = await axios.get(url)
      title.value = data.item.Title
      detail.value = data.item.Detail
      price.value = data.item.Price
      itemPath.value = data.item.Img
      itemFileName.value = data.item.FileName
    }

    const uploadImage = (e) => {
      const file = e.target.files[0]
      const uuid = UUID.generate();
      const extention = file.name.split('.').pop();
      const fileName = uuid + '.' + extention;
      const storageRef = fbRef(storage, 'images/items/'+fileName);

      uploadBytes(storageRef, file).then(() => {
        getDownloadURL(fbRef(storage, 'images/items/'+fileName))
        .then((url) => {
          const newImage = {id: fileName, path: url};
          itemPath.value = newImage.path
          itemFileName.value = fileName
        })
      });

      if(prevItemFileName.value != '') {
        const desertRef = fbRef(storage, 'images/items/'+ prevItemFileName.value);
        deleteObject(desertRef).then(() => {
        }).catch((error) => {
          console.log(error)
        });
      } else {
        const desertRef = fbRef(storage, 'images/items/'+ itemFileName.value);
        deleteObject(desertRef).then(() => {
        }).catch((error) => {
          console.log(error)
        });
      }
      prevItemFileName.value = fileName
    }

    const editItem = async() => {
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
        const url = '/item/edit/' + itemID
        const params = new URLSearchParams()
        params.append('title', title.value)
        params.append('price', price.value)
        params.append('detail', detail.value)
        params.append('seller_id', userID.value)
        params.append('img', itemPath.value)
        params.append('file_name', itemFileName.value)
        await axios.put(url, params)
        router.push('/mypage/item')
      } catch (e) {
        console.log(e)
      }
    }

    const deleteItem = async () => {
      try {
        const url = '/item/delete/' + itemID
        await axios.delete(url)
        router.push('/mypage/item')
      } catch (e) {
        console.log(e)
      }
      const desertRef = fbRef(storage, 'images/items/'+ itemFileName.value);
        deleteObject(desertRef).then(() => {
        // File deleted successfully
        }).catch((error) => {
          console.log(error)
        });
    }

    onMounted(async () => {
      getItem()
    })

    watch(
      detail, () => detail.value,
      title, () => title.value,
      price, () => price.value,
      itemPath, () => itemPath.value,
      prevItemFileName, () => prevItemFileName.value
    )

    return {
      title,
      price,
      detail,
      itemPath,
      editItem,
      deleteItem,
      uploadImage,
    }
  }
}

</script>