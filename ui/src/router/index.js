import { createRouter, createWebHistory } from 'vue-router';
import {
  HelloWorld,
  SignIn,
  SignUp,
  SignUpCompleted,
  LogOut,
  CreateWish,
} from '@/pages';

const routes = [
  {
    path: '/',
    name: 'HelloWorld',
    component: HelloWorld
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
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router