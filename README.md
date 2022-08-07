# 逆フリマアプリ
URL : https://uruapi-ddf2d.web.app/  
API URL : https://uruapi-api.herokuapp.com/

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
* 取引成立時などはメールを送信

## 工夫した点
* 取引成立時などにメールで通知する機能を実装した
* 分かりやすい命名, 統一感のある処理を書くように心掛けた
* 外部のライブラリを積極的に活用した

## 使用技術
* Go (Gin)
* Vue (Vuetify)
* MySQL
* Docker
* SendGrid
* Firebase (Firestore Database, Storage, Authentication)
