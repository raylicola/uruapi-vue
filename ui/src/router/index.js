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
  WishDetail,
  ItemDetail,
  SuggestItem,
  PurchaseItem,
  PostReview,
  PasswordReset,
} from '@/pages';

const routes = [
  {
    path: '/',
    name: 'TopPage',
    component: TopPage
  },
  {
    path: '/wish/:wish_id',
    name: 'WishDetail',
    component: WishDetail
  },
  {
    path: '/item/:item_id',
    name: 'ItemDetail',
    component: ItemDetail
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
    path: '/wish/:wish_id/suggest',
    name: 'SuggestItem',
    component: SuggestItem
  },
  {
    path: '/purchase/:item_id',
    name: 'PurchaseItem',
    component: PurchaseItem
  },
  {
    path: '/mypage/wish/create',
    name: 'CreateWish',
    component: CreateWish
  },
  {
    path: '/mypage/item/create',
    name: 'CreateItem',
    component: CreateItem
  },
  {
    path: '/mypage/wish/edit/:wish_id',
    name: 'EditWish',
    component: EditWish
  },
  {
    path: '/mypage/item/edit/:item_id',
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
  {
    path: '/mypage/review/:item_id',
    name: 'PostReview',
    component: PostReview
  },
  {
    path: '/password_reset',
    name: 'PasswordReset',
    component: PasswordReset
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router