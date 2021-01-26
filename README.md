# yource
[![Build Status](https://travis-ci.com/Diwamoto/yource.svg?branch=main)](https://travis-ci.com/Diwamoto/yource)
A service that allows you to create a blog like slack.


# 環境構築

## 環境イメージ図
![関係イメージ図](https://github.com/Diwamoto/yource/blob/main/docker/relation.png '関係イメージ図')

## フロント開発

フロントにはvue.jsを使用し、SPAとして動かします。

`http://localhost:8082` で起動します。

## バックエンド開発

バックエンドにはginを使用し、dbアクセスのAPIサーバーとして動かします。

`http://localhost:8081` で起動します。

# CI

自動テスト環境にはtravisCIを使用しています。
