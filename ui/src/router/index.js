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
  MyFavoriteItem,
  MyWish,
  MySoldItem,
  MyPurchasedItem,
  WishDetail,
  ItemDetail,
  SuggestItem,
  PurchaseItem,
  PostReview,
  ResetPassword,
  EditProfile,
  UserDetail,
  UserReview,
} from '@/pages';

const routes = [
  {
    path: '/',
    name: 'TopPage',
    component: TopPage
  },
  {
    path: '/wish/:wishID',
    name: 'WishDetail',
    component: WishDetail
  },
  {
    path: '/item/:itemID',
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
    path: '/wish/:wishID/suggest',
    name: 'SuggestItem',
    component: SuggestItem
  },
  {
    path: '/user/:userID',
    name: 'UserDetail',
    component: UserDetail
  },
  {
    path: '/user/:userID/review',
    name: 'UserReview',
    component: UserReview
  },
  {
    path: '/purchase/:itemID',
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
    path: '/mypage/wish/edit/:wishID',
    name: 'EditWish',
    component: EditWish
  },
  {
    path: '/mypage/item/edit/:itemID',
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
    path: '/mypage/favorite',
    name: 'MyFavoriteItem',
    component: MyFavoriteItem
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
    path: '/mypage/review/:itemID',
    name: 'PostReview',
    component: PostReview
  },
  {
    path: '/mypage/profile/edit',
    name: 'EditProfile',
    component: EditProfile
  },
  {
    path: '/password-reset',
    name: 'ResetPassword',
    component: ResetPassword
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router