import { createRouter, createWebHistory } from 'vue-router';
import {
  TopPage,
  SignIn,
  SignUp,
  SignUpCompleted,
  LogOut,
  CreateWish,
  CreateItem,
  EditWish,
  EditItem,
  MyPage,
  MyItem,
  MyWish,
  MySoldItem,
  MyPurchasedItem,
} from '@/pages';

const routes = [
  {
    path: '/',
    name: 'TopPage',
    component: TopPage
  },
  {
    path: '/signin',
    name: 'SignIn',
    component: SignIn
  },
  {
    path: '/signup',
    name: 'SignUp',
    component: SignUp
  },
  {
    path: '/signup-completed',
    name: 'SignUpCompleted',
    component: SignUpCompleted
  },
  {
    path: '/logout',
    name: 'LogOut',
    component: LogOut
  },
  {
    path: '/wish/create',
    name: 'CreateWish',
    component: CreateWish
  },
  {
    path: '/item/create',
    name: 'CreateItem',
    component: CreateItem
  },
  {
    path: '/wish/edit/:wish_id',
    name: 'EditWish',
    component: EditWish
  },
  {
    path: '/item/edit/:item_id',
    name: 'EditItem',
    component: EditItem
  },
  {
    path: '/mypage',
    name: 'MyPage',
    component: MyPage
  },
  {
    path: '/mypage/wish',
    name: 'MyWish',
    component: MyWish
  },
  {
    path: '/mypage/item',
    name: 'MyItem',
    component: MyItem
  },
  {
    path: '/mypage/sold',
    name: 'MySoldItem',
    component: MySoldItem
  },
  {
    path: '/mypage/purchased',
    name: 'MyPurchasedItem',
    component: MyPurchasedItem
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router