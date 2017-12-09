# bots

各サービスのAPIを触ったサンプル集

## 全体的なルール

* 各サービスのAPIアクセス実装は、個別ディレクトリを作成し、pkgとして実装する
* 各サービスのトークンは`token/<サービス名>.txt`とし、pkgディレクトリに配置しない
* `token/*`を`.gitignore`に記載し、トークンの漏洩を防ぐ

## LINE





### References
* [LINE Notify](https://notify-bot.line.me/ja/)
* [LINE Notify API Document](https://notify-bot.line.me/doc/ja/)
