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
右上のボタンから, ユーザー登録, ログインができます
* テストユーザーでのログインも可能
* パスワードリセットページがNot Foundになるバグがあります
```
テストユーザー
メールアドレス：test@example.com
パスワード：password
```

<img src="https://user-images.githubusercontent.com/76393580/184478492-306f1964-6643-4cc7-bfcb-2bf7a4968ae4.png" width="70%">

---

カードをクリックすると, 詳細画面に遷移します
* 自分が出品している商品の提案
* 提案されている商品の閲覧

<img src="https://user-images.githubusercontent.com/76393580/184854467-a8d5e12d-405f-4dd5-8694-b464bbf8724e.png" width="70%">

---

商品のカードをクリックすると, 商品詳細画面に遷移します
* 商品に対するチャットの閲覧・投稿

<img src="https://user-images.githubusercontent.com/76393580/184858642-c0896b11-21cd-4fae-bd1c-26b6de579ecd.png" width="70%">

---

右上のボタンから, マイページへ遷移できます
  * 欲しいものや出品したいものを投稿・編集・削除・閲覧
  * 売買履歴の閲覧

<img src="https://user-images.githubusercontent.com/76393580/184855185-cfecfe93-5548-4daf-9183-4b4807554839.png" width="70%">

---

ユーザーのアイコンをクリックすると, ユーザーのプロフィール画面に遷移します
* ユーザーの出品している商品の閲覧
* ユーザーへのレビューの閲覧

---

<img src="https://user-images.githubusercontent.com/76393580/184854492-4c2844df-a0fc-4264-98a6-63cd16ab6c59.png" width="70%">
