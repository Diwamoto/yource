# yource
[![Build Status](https://travis-ci.com/Diwamoto/yource.svg?branch=main)](https://travis-ci.com/Diwamoto/yource)
[![Coverage Status](https://coveralls.io/repos/github/Diwamoto/yource/badge.svg?branch=)](https://coveralls.io/github/Diwamoto/yource?branch=)

A service that allows you to create a blog like slack.


# 環境構築

## 環境イメージ図
![関係イメージ図](https://github.com/Diwamoto/yource/blob/main/docker/relation.png '関係イメージ図')

## フロント開発

フロントにはvue.jsを使用し、SPAとして動かします。
最初はdocker内にサーバを立てて運用していましたが、自動ビルドのファイルウォッチが遅いためやめました。
`npm install`したのち`npm run serve`で起動します。


`http://localhost:9092` で起動します。

## バックエンド開発

バックエンドにはginを使用し、dbアクセスのAPIサーバーとして動かします。

`https://localhost:9091` で起動します。

# CI

自動テスト環境にはtravisCIを使用しています。
