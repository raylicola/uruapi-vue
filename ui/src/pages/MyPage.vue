<template>
  <small-space />
  <navigation-drawer />
  <v-row align="center">
    <v-col cols="2">
      <base-avatar
        size="100"
        :src="iconPath"
      />
    </v-col>
    <v-col cols="5">
      <div class="text-h5">
        {{userName}}
      </div>
    </v-col>
    <v-col cols="2">
      <base-button
        text="プロフィールを編集"
        @click="editProfile"
      />
    </v-col>
  </v-row>
  <small-space />
  <div style="white-space: pre-line;">
    {{introduction}}
  </div>
</template>

<script>
import {
  SmallSpace,
  NavigationDrawer,
  BaseAvatar,
  BaseButton,
} from '@/components'
import { computed, onMounted, ref, watch } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import { db } from '@/firebase'
import { doc, getDoc } from "firebase/firestore";

export default {
  name: 'MyPage',
  components: {
    'small-space': SmallSpace,
    'navigation-drawer': NavigationDrawer,
    'base-avatar': BaseAvatar,
    'base-button': BaseButton,
  },
  setup() {
    const store = useStore()
    const router = useRouter()

    const userName = computed(() => store.state.userName)
    const userID = computed(() => store.state.userID)
    const iconPath = ref('')
    const introduction = ref('')

    const editProfile = () => {
      router.push('/mypage/profile/edit')
    }

    const getUserProfile = async () => {
      const docRef = doc(db, "users", userID.value);
      const docSnap = await getDoc(docRef);
      iconPath.value = docSnap.data().icon_path ? docSnap.data().icon_path : require('@/assets/default_icon.jpg')
      introduction.value = docSnap.data().introduction
    }

    onMounted(async () => {
      getUserProfile()
    })

    watch(
      introduction, () => introduction.value,
      iconPath, () => iconPath.value,
    )

    return {
      userName,
      introduction,
      iconPath,
      editProfile,
    }
  }
}
</script>