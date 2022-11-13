# 概要
GoによるRESTful API作成を始めるためのガイドプロジェクトです。

# 想定する読者

下記についてはある程度学習済みであることを前提としています。  
もし名前も聞いたことがないということであれば先にこれらを学習しましょう。
- Go以外でWebのAPIを自作したことがある
- GitやGitHubを利用してリポジトリを作成したり、Pull Requestを作成したことがある
- Docker、Docker Composeを利用してMySQLなどのミドルウェアやRESTful APIを構築したことがある

# 前提条件
- Goをインストールしていること
  - 最悪なしでもDockerがあればハンズオンはできます
- Dockerをインストールしていること
- GitHubアカウントを持っていること
- なんからのIDEをインストールしていること
  - Visual Studio CodeやGoLandなど
- curlやPostmanなどのHTTPクライアントを使ったことがある

# ハンズオンをすすめるにあたって

このハンズオンをすすめるにあたってPull Requestの作成およびレビューを求める箇所がございます。  
ただ、ご自身のみで進めても問題ございませんので必要がないと判断されれば特にPull Requestの作成やレビュー依頼はせず進めてもらって結構です。  

# 準備

まずはこのハンズオンのリポジトリをforkしてください。  
forkできましたら、ご自身のローカルPCにcloneしてエディタでプロジェクトを開いてください。

参考: https://docs.github.com/ja/get-started/quickstart/fork-a-repo  

# このプロジェクトについて

## アプリケーションの構成

オニオンアーキテクチャを意識した作りになっています。  
各レイヤーの依存関係は下記を守っています。  
handler --> usecase --> domain <-- infra

### main.go

エントリーポイントです。  
Dockerを使わず、エディタからアプリケーションを起動する場合にはこのファイルから起動できます。  

### handler

HTTPリクエストをハンドリングします。  

### domain

ビジネスロジックに関連する状態や振る舞いを担うEntityやRepositoryなどを記述します。  

### usecase

domain層が公開する関数を組み合わせてユースケースを実現します。  

### infra

データベースなどのStorageや外部APIとのやり取りを行います。  
ただし、このハンズオンではデータベースとのやり取りは実装しておらず、下記のように仮の実装をしています。  

https://github.com/raisetech-for-student/golang-web-api-hands-on/blob/5edba42f463dea02ce1c482e78872d398361902f/infra/dao/book.go#L14-L29  

## 利用しているライブラリやフレームワークについて

### go-chi/chiおよびgo-chi/render

https://github.com/go-chi/chi  
https://github.com/go-chi/render  

HTTPリクエストをハンドリングすることができます。  
Goはnet/httpというHTTPクライアントとサーバーの実装を提供していますが、筆者が個人的にgo-chiに興味があるので採用しています。  

### cosmtrek/air

https://github.com/cosmtrek/air  

Live Reloadを実現するために導入しています。  

### mvdan/gofumpt

https://github.com/mvdan/gofumpt  

Goのstandard libraryの1つである`gofmt`よりも厳密にフォーマットするために導入しています。  

### golangci/golangci-lint

https://github.com/golangci/golangci-lint  

Star数も多く、Go界隈で人気のLinterです。  

# 起動手順

```
% docker compose up
[+] Running 1/1
 ⠿ Container dependency-injection-sample-app-1  Recreated                                                                                                                                                                          0.1s
Attaching to golang-web-api-hands-on
golang-web-api-hands-on  | 
golang-web-api-hands-on  |   __    _   ___  
golang-web-api-hands-on  |  / /\  | | | |_) 
golang-web-api-hands-on  | /_/--\ |_| |_| \_ , built with Go 
golang-web-api-hands-on  | 
golang-web-api-hands-on  | watching .
golang-web-api-hands-on  | watching domain
golang-web-api-hands-on  | watching domain/model
golang-web-api-hands-on  | watching domain/repository
golang-web-api-hands-on  | watching handler
golang-web-api-hands-on  | watching handler/response
golang-web-api-hands-on  | watching infra
golang-web-api-hands-on  | watching infra/dao
golang-web-api-hands-on  | !exclude tmp
golang-web-api-hands-on  | watching usecase
golang-web-api-hands-on  | building...
golang-web-api-hands-on  | go: downloading github.com/go-chi/chi/v5 v5.0.7
golang-web-api-hands-on  | go: downloading github.com/go-chi/render v1.0.2
golang-web-api-hands-on  | go: downloading github.com/ajg/form v1.5.1
golang-web-api-hands-on  | running...
```

下記コマンドを実行してレスポンスが得られることを確認してください。  

```
% curl http://localhost:8080/hello   
{"message":"hello world"}
```

# /api/v1/booksにリクエストしてみましょう

サンプルとして実装したAPIにリクエストを送ってみましょう。  

`-i`オプションをつけてHTTPステータスコードが200であることも確認しましょう。  

```
% curl -i http://localhost:8080/api/v1/books/1
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Tue, 22 Nov 2022 16:09:23 GMT
Content-Length: 57

{"id":"1","name":"The Lord of the Rings","price":"1600"}
```

`/books/1`から`/books/4`までは書籍情報を返却します。  

`/books/5`でリクエストしましょう。  

HTTPステータスコードが404になり、レスポンスボディとしてリソースが見つからなかったというメッセージが返却されます。  
```
% curl -i http://localhost:8080/api/v1/books/5
HTTP/1.1 404 Not Found
Content-Type: application/json; charset=utf-8
Date: Tue, 22 Nov 2022 16:12:28 GMT
Content-Length: 33

{"message":"resource not found"}
```

# hello worldの実装を変更してみましょう

下記の箇所を修正してみましょう。  

https://github.com/raisetech-for-student/golang-web-api-hands-on/blob/5edba42f463dea02ce1c482e78872d398361902f/main.go#L21-L25  

たとえば、`hello world`を`good morning`に修正してみてください。

コンソールを確認すると、変更が検知されていることがわかります。  

```
golang-web-api-hands-on  | main.go has changed
golang-web-api-hands-on  | building...
golang-web-api-hands-on  | running...
```

再度、リクエストしてレスポンスが変化することを確認してください。  

```
% curl http://localhost:8080/hello
{"message":"good morning"}
```

## handlerの実装をしてみましょう

下記リクエストを実行してレスポンスを確認してください。  
```
% curl http://localhost:8080/message
{"message":"There is always light behind the clouds."}
```

`handler/message.go`を編集して、下記いづれかのメッセージをランダムに返すように修正してみてください。  

- There is always light behind the clouds.
- Change before you have to.
- If you can dream it, you can do it.
- Love the life you live. Live the life you love.

実装にはGoのArraysやSlicesをどう扱うか、ArraysやSlicesからランダムに値を取り出す方法を学ぶ必要があります。  

https://go.dev/tour/moretypes/7  

修正したら、下記コマンドを実行してみてください。  
それぞれフォーマットとLintを実行します。
```
% make fmt
% make lint
```

コードが自動でフォーマットされます。  
Lintエラーが出た場合、エラーメッセージに合わせて修正しましょう。

自身のリポジトリにてPull Requestを作成してレビュー依頼をしてください。  
※このハンズオンリポジトリではなくご自身のリポジトリのmainブランチに対してマージするPull Requestを作成してください。  

## usecaseの実装をしてみましょう

先の手順でhandlerの実装をしました。  

この実装をusecaseとして実装しましょう。  

下記の手順で実装しましょう。  

1. `usecase/message.go`を作成し、下記を記述してください。
```go
package usecase

import (
	"context"
)

type Message interface {
	Get(ctx context.Context) string
}

type messageUseCase struct{}

func NewMessage() Message {
	return &messageUseCase{}
}

func (m *messageUseCase) Get(ctx context.Context) string {
	return ""
}
```
2. Getメソッドの`return ""`を修正してhandlerに記述したメッセージをランダムに返す処理を移してください。  
3. `messageHandler`の記述を修正して、usecaseのMessageを呼び出すように修正してください。  
※記述方法はusecase/book.goを参考にしましょう。  
4. `main.go`の下記箇所の直前に`messageUseCase`を宣言して、`messageHandler`を使える状態にしましょう。
```go
messageHandler := handler.NewMessage()
```
こちらも`bookUseCase`、`bookHandler`を参考にするとよいです。  

下記リクエストを実行してレスポンスを確認してください。
```
% curl http://localhost:8080/message
{"message":"There is always light behind the clouds."}
```

フォーマットとLintを実行します。  
```
% make fmt
% make lint
```

Pull Requestを作成しましょう。  

## repositoryとinfra(dao)の実装をしてみましょう

最後にusecaseに移動したメッセージ取得処理をdaoに記述しましょう。  
本来はdaoではデータベースから値を取得しますが、このハンズオンではdao内でいくつかの固定のメッセージからランダムに1つのメッセージを返すように実装すればよいです。  

つまりhandler -> usecaseに移動した処理をusecase -> daoに移すだけです。  

気をつけるポイントをいくつか記載しておきます。  
すこしヒントが少ないかもしれませんが頑張って実装してみてください。  

- repository/message.goを作成し、`Message`という名前でinterfaceを定義する
- repositoryに定義した`Message`の実装はinfra/daoにmessage.goを作成しそこに記述する
- messageUseCaseからはrepositoryに定義した`Message`を利用する
- main.goを修正してmessageUseCaseを使える状態にする

完成したらフォーマットとLintを実行して、Pull Requestを作成してください。
