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
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router