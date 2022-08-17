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

<img src="https://user-images.githubusercontent.com/76393580/185062712-80bc47da-5090-4c29-a412-18650dab6a30.png" width="60%">

---

商品のカードをクリックすると, 商品詳細画面に遷移します
* 商品に対するチャットの閲覧・投稿
* 商品の購入・お気に入り登録

<img src="https://user-images.githubusercontent.com/76393580/185062666-a3157280-3665-449f-9d90-c6dd2b79c902.png" width="60%">

---

ユーザーのアイコンをクリックすると, ユーザーのプロフィール画面に遷移します
* ユーザーの出品している商品の閲覧
* ユーザーへのレビューの閲覧

<img src="https://user-images.githubusercontent.com/76393580/185062537-82bab6f0-d55d-4968-be98-16e1791d0407.png" width="48%">　<img src="https://user-images.githubusercontent.com/76393580/185062548-785d2c7d-9ee7-4f93-8481-b1b17a65629f.png" width="48%">

---

右上のボタンから, マイページへ遷移できます
  * 欲しいものや出品したいものを投稿・編集・削除・閲覧
  * 売買履歴の閲覧
  * お気に入り登録した商品の閲覧

<img src="https://user-images.githubusercontent.com/76393580/185062463-2ac14eb1-46fd-499c-b5f8-6f08dfd3f304.png" width="60%">
