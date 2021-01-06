# yource
A service that allows you to create a blog like slack.

# 環境構築

## フロント開発

フロントにはvue.jsを使用し、SPAとして動かします。

環境はvueについてくるサーバーを利用します。
カレントディレクトリから

```
cd front
npm run dev
```

で起動します。

## バックエンド開発

バックエンドにはcakePHPを使用し、APIサーバーとして稼働させます。

環境は[Diwamoto/vagrant-lamp](https://github.com/Diwamoto/vagrant-lamp)を使用しvagrant上にdockerを立ててlampで運用します。
`./backend/config/app_local.php`のデータベース設定を編集してください。
後は`http://yource.localhost`で動きます。