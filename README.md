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


## slack


### References
* [golang で始める Slack bot 開発 - at kaneshin](http://blog.kaneshin.co/entry/2016/12/03/162653)
* [nlopes/slack](https://github.com/nlopes/slack)