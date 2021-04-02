# yource
[![Build Status](https://travis-ci.com/Diwamoto/yource.svg?branch=main)](https://travis-ci.com/Diwamoto/yource)
[![Coverage Status](https://coveralls.io/repos/github/Diwamoto/yource/badge.svg?branch=)](https://coveralls.io/github/Diwamoto/yource?branch=)

A service that allows you to create a blog like slack.

https://yource.space


# 環境構築

```
cd docker/ssl
./initssl.sh
cd ..
docker-compose up -d
cd ..
cp docker/ssl/keys frontend/keys
npm run serve
```
で起動します。
ローカル環境は全てssl通信に対応しています。

## 環境イメージ図
![関係イメージ図](https://github.com/Diwamoto/yource/blob/main/docker/relation.png '関係イメージ図')

## フロント開発

フロントにはvue.jsを使用し、SPAとして動かします。
最初はdocker内にサーバを立てて運用していましたが、自動ビルドのファイルウォッチが遅いためやめました。
`npm install`したのち`npm run serve`で起動します。


`https://localhost:9092` で起動します。


## バックエンド開発

バックエンドにはginを使用し、dbアクセスのAPIサーバーとして動かします。

`https://localhost:9091` で起動します。


# CI/CD

自動テスト環境にはtravisCI/CircleCIを使用しています。
AWS上にCircleCIを使って自動デプロイしています。
設定ファイルは`.circleCI/config.yml`にあり、sshでawsにつないで
`deployment.sh`を叩くだけの簡単なシステムです。

# デプロイ
今まではAWSにフロント、バックエンド両方おいてありましたが、[vercel](https://vercel.com/)
という素晴らしいフロントエンドデプロイサービスがありますので、フロントはそちらを利用して
バックエンドは変わらずawsを利用する形にしようかと思っています。
