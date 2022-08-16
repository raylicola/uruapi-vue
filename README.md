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
* 右上のボタンから, ユーザー登録, ログインを行う
  * テストユーザーでのログインも可能
```
テストユーザー
メールアドレス：test@example.com
パスワード：password
```

<img src="https://user-images.githubusercontent.com/76393580/184478492-306f1964-6643-4cc7-bfcb-2bf7a4968ae4.png" width="70%">

* カードをクリックすると詳細画面に移る
  * 自分が出品している商品の提案
  * 提案されている商品の閲覧

<img src="https://user-images.githubusercontent.com/76393580/184342503-4982ba41-f9ce-4d85-ad18-2ff615b3334f.png" width="70%">

* 商品のカードをクリックすると商品詳細画面に移る
  * 出品者へのレビューの閲覧
  * 商品に対するチャットの閲覧・投稿

<img src="https://user-images.githubusercontent.com/76393580/184342500-431be561-c0e2-4dff-86be-c9f6085987d0.png" width="70%">

* 右上のボタンから, マイページへ遷移
  * 欲しいものや出品したいものを投稿・編集・削除・閲覧
  * 売買履歴の閲覧

<img src="https://user-images.githubusercontent.com/76393580/184343069-f496f68f-dc83-4097-a9be-67c8096219da.png" width="70%">
