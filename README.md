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

`slack_<bot名>.json`内の`api_token`をもとにRTM(Real Time Message) connectionを作成し、goroutine化する。

```go

rtm, err := myslack.NewRTM()
if err != nil {
  panic(err)
}
go rtm.ManageConnection()


```

#### メッセージを送る

送信先の`channel`を設定する必要があり、`channel`の`id`は[channels.list](https://api.slack.com/methods/channels.list/test)で取得できる。

`Generate tokens to test with here`からトークンを作った状態で`Test Method`を押すとスペースのchannel listが得られる。


```go

msg := "This is a sample message"
rtm.SendMessage(rtm.NewOutgoingMessage(msg, channel))

```

#### メッセージを受け取る

`rtm.IncomingEvents`チャンネルを介して、イベントループを回す。
`rtm.IncomingEvents.Data.(type)`でそれぞれのイベントごとに処理を変えられる。  
下記の`ev.Channel`でメッセージが送られてきたchannelへレスポンスを返せる。

```go

for {
  select {
  case msg := <-rtm.IncomingEvents:
    switch ev := msg.Data.(type) {
    case *slack.HelloEvent:
      log.Print("bot start")
    case *slack.MessageEvent:
      log.Printf("Message: %v\n", ev)
      rtm.SendMessage(rtm.NewOutgoingMessage("new message", ev.Channel))
    case *slack.InvalidAuthEvent:
      log.Print("Invalid credentials")
      return
    }
  }
}

```

### トークン
* API tokenはbotごとに割り当てられるため、`<用途>`を`bot名`としたものを使う


## 複数サービスの連携

`chan`を使うと簡単に書ける。以下ではslackからメッセージを受信したら、そのメッセージをLINEに送る例


```go

package main

import (
	"log"

	"github.com/nlopes/slack"

	"./line"
	"./myslack"
)

func main() {
	// メッセージをやり取りするためのchannel
	ch := make(chan string)

	// slack bot起動
	go func() {
		rtm, err := myslack.NewRTM()
		if err != nil {
			panic(err)
		}

		go rtm.ManageConnection()

		for {
			select {
			case msg := <-rtm.IncomingEvents: // slackからのメッセージ受信
				switch ev := msg.Data.(type) {
				case *slack.MessageEvent:
					log.Print(ev)
					ch <- ev.Msg.Text
				case *slack.HelloEvent:
					log.Print("bot start")
				case *slack.InvalidAuthEvent:
					log.Print("Invalid credentials")
				}
			}
		}
	}()

	for {
		select {
		case msg := <-ch: // LINEにメッセージを送る
			log.Print("message recieved")
			err := line.SendMessage(msg)
			if err != nil {
				panic(err)
			}
		}
	}
}

```


## 定期的にポーリングしたいとき

```go

func main() {
	t := time.NewTicker(60 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
      doSomething()
		}
	}
}

```

### References
* [golang で始める Slack bot 開発 - at kaneshin](http://blog.kaneshin.co/entry/2016/12/03/162653)
* [nlopes/slack](https://github.com/nlopes/slack)