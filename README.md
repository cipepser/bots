# bots

各サービスのAPIを触ったサンプル集

## 全体的なルール

* 各サービスのAPIアクセス実装は、個別ディレクトリを作成し、pkgとして実装する
* 各サービスのトークンは`token/<サービス名>_<用途>.json`に記載し、pkgディレクトリに配置しない
* `token/*`を`.gitignore`に記載し、トークンの漏洩を防ぐ
* トークンはjsonフォーマットで記載し、各サービス内にjsonファイルごとに`Token<用途>`型を定義する

## LINE

### How to Use

```go

msg := "hello"

err := line.SendMessage(msg)
if err != nil {
  panic(err)
}

```

### トークン

* トークルーム用のトークンは、`<用途>`を`トークルーム名(を英数にしたもの)`を使う


### References
* [LINE Notify](https://notify-bot.line.me/ja/)
* [LINE Notify API Document](https://notify-bot.line.me/doc/ja/)


## facebook

### How to Use

```go

URL = "https://graph.facebook.com/v2.11/<user name>/feed"

f, err := facebook.GetFeed(URL)
if err != nil {
  panic(err)
}

```

### References
* [curlでFacebook API叩くまでがわりとめんどかったのでメモ - DRYな備忘録](http://otiai10.hatenablog.com/entry/2014/11/26/152404)
* [アクセス許可のリファレンス - Facebookログイン](https://developers.facebook.com/docs/facebook-login/permissions/)


## slack
[nlopes/slack](https://github.com/nlopes/slack)とパッケージ名が衝突したので、自作packageの名前は`myslack`とする。

[golang で始める Slack bot 開発 - at kaneshin](http://blog.kaneshin.co/entry/2016/12/03/162653)にアクセストークンの取得方法とライブラリの使い方が書いてあるので参照。自分が使いそうなものをまとめると以下。

### How to Use

#### 準備

```go



```

#### メッセージを送る

```go



```
#### メッセージを受け取る

```go



```



### トークン
* API tokenはbotごとに割り当てられるため、`<用途>`を`bot名`としたものを使う


### References
* [golang で始める Slack bot 開発 - at kaneshin](http://blog.kaneshin.co/entry/2016/12/03/162653)
* [nlopes/slack](https://github.com/nlopes/slack)