# 逆フリマアプリ
URL : https://uruapi-ddf2d.web.app
```
テストユーザー
メールアドレス：test@example.com
パスワード：password
```

## 概要
既存のフリマアプリとは異なり, 売り手側が買い手側に物を売り込むことが出来るアプリ.<br>
買い手側は, 自分の欲しいものの概要を投稿し, 提案されたアイテムを購入することができる.<br>

## 目的
既存のフリマアプリでは, 商品数が多く, どんなに良い商品でもなかなか見つけてもらえないという課題がある.<br>
このアプリでは, 売り手側が自分の商品を積極的にアピールすることで商品を認知させることができる.<br>

## 機能
* 欲しいものを投稿
* 売りたいものを投稿
* 欲しいものを投稿した相手に, 自分の出品した商品を提案
* 出品物についてのチャット
* 出品者へのレビュー

## 工夫した点
* 分かりやすい命名, 統一感のある処理を書くように心掛けた
* 外部のライブラリを積極的に活用した

## 使用技術
* Go (Gin)
* Vue (Vuetify)
* MySQL
* Docker
* Firebase (Firestore Database, Storage, Authentication)

## 使い方
右上のボタンから, ユーザー登録, ログインができます.<br>
テストユーザーでのログインも可能です.
```
テストユーザー
メールアドレス：test@example.com
パスワード：password
```

<img src="https://user-images.githubusercontent.com/76393580/185054554-6f811911-dbfb-4351-8e4f-06c9cf3635d4.png" width="60%">

---

カードをクリックすると, 詳細画面に遷移します
* 自分が出品している商品の提案
* 提案されている商品の閲覧

<img src="https://user-images.githubusercontent.com/76393580/185055205-58ae1ad4-ec85-4668-b778-f98665dfa1f0.png" width="60%">

---

商品のカードをクリックすると, 商品詳細画面に遷移します
* 商品に対するチャットの閲覧・投稿
* 商品の購入・お気に入り登録

<img src="https://user-images.githubusercontent.com/76393580/185055336-24c7f44c-dd70-4c47-91ae-6a6e9ef5fd70.png" width="60%">

---

ユーザーのアイコンをクリックすると, ユーザーのプロフィール画面に遷移します
* ユーザーの出品している商品の閲覧
* ユーザーへのレビューの閲覧

<img src="https://user-images.githubusercontent.com/76393580/185054022-b9304384-b34d-46cb-aa1a-7f6d86600b84.png" width="48%">　<img src="https://user-images.githubusercontent.com/76393580/185054044-385309d0-1f65-48e4-8fb9-51d7fbe8965b.png" width="48%">

---

右上のボタンから, マイページへ遷移できます
  * 欲しいものや出品したいものを投稿・編集・削除・閲覧
  * 売買履歴の閲覧
  * お気に入り登録した商品の閲覧

<img src="https://user-images.githubusercontent.com/76393580/185056338-4dd312c8-5cac-45ba-9e76-c0497f4d57a7.png" width="60%">

