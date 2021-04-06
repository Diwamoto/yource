# yource
[![Build Status](https://travis-ci.com/Diwamoto/yource.svg?branch=main)](https://travis-ci.com/Diwamoto/yource)
[![Coverage Status](https://coveralls.io/repos/github/Diwamoto/yource/badge.svg?branch=)](https://coveralls.io/github/Diwamoto/yource?branch=)

A service that allows you to create a blog like slack.

https://yource.space


# 環境構築

```
cd docker/ssl
./initssl.sh
cp docker/ssl/keys frontend/keys
cd ../..
./docker_start.sh
```
で起動します。
全てssl通信に対応しています。

|                   | ローカル環境                         | 本番環境                     | 
| ----------------- | ------------------------------------ | ---------------------------- | 
| フロントエンド    | yource.localhost, *.yource.localhost | [yource.space](yource.space), *.yource.space | 
| バックエンド(api) | api.yource.localhost                 | api.yource.space             | 
| websocket         | ws.yource.localhost/socket           | ws.yource.space/socket       | 

## 環境イメージ図
![関係イメージ図](https://github.com/Diwamoto/yource/blob/main/docker/relation.png '関係イメージ図')

## フロント開発

フロントにはvue.jsを使用し、SPAとして動かします。
最初はdocker内にサーバを立てて運用していましたが、自動ビルドのファイルウォッチが遅いためやめました。
`yarn`したのち`yarn dev`で起動します。


## バックエンド開発

バックエンドにはginを使用し、dbアクセスのAPIサーバーとして動かします。


# CI/CD

自動テスト環境にはtravisCI/CircleCIを使用しています。
AWS上にCircleCIを使って自動デプロイしています。
設定ファイルは`.circleCI/config.yml`にあり、sshでawsにつないで
`deployment.sh`を叩くだけの簡単なシステムです。

# 本番情報
フロントは[vercel](https://vercel.com/)、バックエンドとwebソケットサーバーはAWSに共存させています。
バックエンドへのAPIアクセスはCORSで絞って制限をかけています。
バックエンドはgoのサーバーをライブリロードで起動して、
nginxでリバースプロキシとSSL通信を行っています。
