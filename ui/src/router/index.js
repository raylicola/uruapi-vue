import { createRouter, createWebHistory } from 'vue-router';
import {
  HelloWorld,
  SignIn,
  SignUp,
  SignUpCompleted
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
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router