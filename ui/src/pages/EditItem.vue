<template>
  <div class="text-h5">
    出品
  </div>
  <small-space />
  <base-text-field label="タイトル" v-model="title"/>
  <base-text-field label="価格（円）" v-model="price" />
  <base-text-area label="説明" v-model="detail" />
  <v-row>
    <v-col cols="3">
      <base-file-input
        @change="uploadImage"
        label="画像を変更"
      />
    </v-col>
    <v-col cols="1" />
    <v-col cols="4">
      <image-preview :src="img" />
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
  SmallSpace
  } from '@/components'
import { ref } from '@vue/reactivity'
import { useRoute, useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { computed, onMounted, watch } from 'vue'
import axios from 'axios'
import UUID from 'uuidjs'
import { storage } from '@/firebase'
import { getDownloadURL, ref as fb_ref, uploadBytes, deleteObject } from "firebase/storage";

export default {
  name: 'EditItem',
  components: {
    'base-button': BaseButton,
    'base-text-field': BaseTextField,
    'base-text-area': BaseTextArea,
    'base-file-input': BaseFileInput,
    'image-preview': ImagePreview,
    'small-space': SmallSpace,
  },
  setup(){
    const title = ref('')
    const detail = ref('')
    const price = ref('')
    const img = ref('')
    const imgName = ref('')
    const prevFileName = ref('')

    const router = useRouter()
    const route = useRoute()
    const store = useStore()
    const user_id = computed(() => store.state.user_id)

    const getItem = async () => {
      const url = 'item/' + route.params.item_id
      const {data} = await axios.get(url)
      title.value = data.item.Title
      detail.value = data.item.Detail
      price.value = data.item.Price
      img.value = data.item.Img
      imgName.value = data.item.FileName
    }

    const uploadImage = (e) => {
      const file = e.target.files[0]
      const uuid = UUID.generate();
      const extention = file.name.split('.').pop();
      const fileName = uuid + '.' + extention;
      const storageRef = fb_ref(storage, 'images/items/'+fileName);

      uploadBytes(storageRef, file).then(() => {
        getDownloadURL(fb_ref(storage, 'images/items/'+fileName))
        .then((url) => {
          const newImage = {id: fileName, path: url};
          img.value = newImage.path
          imgName.value = fileName
        })
      });

      if(prevFileName.value != '') {
        const desertRef = fb_ref(storage, 'images/items/'+ prevFileName.value);
        deleteObject(desertRef).then(() => {
        // File deleted successfully
        }).catch((error) => {
          console.log(error)
        });
      } else {
        const desertRef = fb_ref(storage, 'images/items/'+ imgName.value);
        deleteObject(desertRef).then(() => {
        // File deleted successfully
        }).catch((error) => {
          console.log(error)
        });
      }
      prevFileName.value = fileName
    }

    const editItem = async() => {
      if (
        title.value === '' ||
        price.value === '' ||
        detail.value === '') {
          alert('必須項目が未入力です')
          return false
      }

      if(img.value == '') {
        alert('商品の画像を選択してください')
        return false
      }

      try {
        const url = '/item/edit/' + route.params.item_id
        const params = new URLSearchParams()
        params.append('title', title.value)
        params.append('price', price.value)
        params.append('detail', detail.value)
        params.append('seller_id', user_id.value)
        params.append('img', img.value)
        params.append('file_name', imgName.value)
        await axios.put(url, params)
        router.push('/mypage/item')
      } catch (e) {
        console.log(e)
      }
    }

    const deleteItem = async () => {
      try {
        const url = '/item/delete/' + route.params.item_id
        await axios.delete(url)
        router.push('/mypage/item')
      } catch (e) {
        console.log(e)
      }
      const desertRef = fb_ref(storage, 'images/items/'+ imgName.value);
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
      img, () => img.value,
      prevFileName, () => prevFileName.value
    )

    return {
      title,
      price,
      detail,
      img,
      editItem,
      deleteItem,
      uploadImage,
    }
  }
}

</script>