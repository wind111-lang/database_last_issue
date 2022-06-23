# database_last_issue

##### ※ Linux,Macを使用している前提  

## データベース,テーブル,ユーザの作成,作成されたユーザへの権限付与

db ディレクトリに`chatdb.sql`が入っているので，MySQL に root でログインした状態で，  
Macの場合,`source /Users/<your_username>/database_last_issue/db/chatdb.sql`と入力する．  
Linuxの場合,`source /home/<your_username>/database_last_issue/db/chatdb.sql`と入力する．  
そうすると，データベース，テーブル，ユーザ(user)が作成され，作成されたユーザに権限が付与される．  

## 作成されたユーザへのログイン

`mysql -u user -p`と入力し，パスワードを求められたら`hoge`と入力するとログインができる．  
データベース名は`chatdb`なので，`use chatdb`と入力すれば切り替えができる．テーブルは  
`chat_log, members, azureapi` のそれぞれ 3 つが入っており，かつ `select * from azureapi`  
と入力した際に各々のカラムに値が入っていれば準備完了．  

## Goのインストール


## 実行準備

実行する前に，必要なモジュールを取得するために
`go mod tidy`と入力してモジュールのインストールを行う

## 実行

`go run main.go`と入力してローカルサーバを開く
