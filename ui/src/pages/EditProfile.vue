<template>
  <back-button />
  <div class="text-h5">
    プロフィーを編集
  </div>
  <small-space />
  <v-row align="center">
    <v-col cols="2">
      <base-avatar
        size="100"
        :src="iconPath"
      />
    </v-col>
    <v-col cols="2">
      <base-file-input
        @change="uploadImage"
        label="アイコン画像を変更"
      />
    </v-col>
  </v-row>
  <small-space />
  <v-row>
    <v-col
      cols="12"
      md="6"
    >
      <base-text-field label="ユーザーネーム" v-model="userName"/>
      <base-text-area label="自己紹介" v-model="introduction" />
      <small-space />
      <base-button text="更新" @click="updateProfile" class="mx-5"/>
    </v-col>
  </v-row>
</template>

<script>
import {
  BaseButton,
  BaseTextField,
  BaseTextArea,
  BaseFileInput,
  SmallSpace,
  BaseAvatar,
  BackButton,
} from '@/components'
import { ref } from '@vue/reactivity'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { computed, onMounted, watch } from 'vue'
import UUID from 'uuidjs'
import { storage } from '@/firebase'
import { getDownloadURL, ref as fb_ref, uploadBytes, deleteObject } from "firebase/storage";
import { doc, updateDoc, getDoc } from "firebase/firestore";
import { db } from '@/firebase'

export default {
  name: 'EditItem',
  components: {
    'base-button': BaseButton,
    'base-text-field': BaseTextField,
    'base-text-area': BaseTextArea,
    'base-file-input': BaseFileInput,
    'small-space': SmallSpace,
    'base-avatar': BaseAvatar,
    'back-button': BackButton,
  },
  setup(){
    const introduction = ref('')
    const iconPath = ref('')
    const userName = ref('')
    const iconFileName = ref('')
    const prevIconFileName = ref('')

    const router = useRouter()
    const store = useStore()
    const userID = computed(() => store.state.userID)

    const uploadImage = (e) => {
      const file = e.target.files[0]
      const uuid = UUID.generate();
      const extention = file.name.split('.').pop();
      const fileName = uuid + '.' + extention;
      const storageRef = fb_ref(storage, 'images/icons/'+fileName);

      uploadBytes(storageRef, file).then(() => {
        getDownloadURL(fb_ref(storage, 'images/icons/'+fileName))
        .then((url) => {
          const newImage = {id: fileName, path: url};
          iconPath.value = newImage.path
          iconFileName.value = fileName
        })
      });

      if(prevIconFileName.value != '') {
        const desertRef = fb_ref(storage, 'images/icons/'+ prevIconFileName.value);
        deleteObject(desertRef).then(() => {
        // File deleted successfully
        }).catch((error) => {
          console.log(error)
        });
      } else {
        const desertRef = fb_ref(storage, 'images/icons/'+ iconFileName.value);
        deleteObject(desertRef).then(() => {
        // File deleted successfully
        }).catch((error) => {
          console.log(error)
        });
      }
      prevIconFileName.value = fileName
    }

    const updateProfile = async() => {
      await updateDoc(doc(db, "users", userID.value), {
        username: userName.value,
        introduction: introduction.value,
        icon_path: iconPath.value == require('@/assets/default_icon.jpg') ? "" : iconPath.value,
      });
      store.commit('setUserName', userName.value)
      router.push('/mypage')
    }

    const getUserProfile = async () => {
      const docRef = doc(db, "users", userID.value);
      const docSnap = await getDoc(docRef);
      iconPath.value = docSnap.data().icon_path ? docSnap.data().icon_path : require('@/assets/default_icon.jpg')
      introduction.value = docSnap.data().introduction
      userName.value = docSnap.data().username
    }

    onMounted(async () => {
      getUserProfile()
    })

    watch(
      introduction, () => introduction.value,
      userName, () => userName.value,
      iconPath, () => iconPath.value,
      prevIconFileName, () => prevIconFileName.value
    )

    return {
      userName,
      introduction,
      iconPath,
      updateProfile,
      uploadImage,
    }
  }
}

</script>