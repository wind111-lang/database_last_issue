# database_last_issue

| 使用技術              | バージョン | 
| --------------------- | ---------- | 
| Go                    | 1.18.3     | 
| Gin(Web Framework)    | 1.8.1      | 
| Gorm(SQL Framework) | 1.9.16     | 

##### ※ Linux,Mac を使用している前提

## テーブルの概要

### members テーブル

- id bigint not null auto_increment
- username varchar(128)
- password varchar(128)
- birthday char(10)

### chat_log テーブル

- id bigint not null auto_increment
- username varchar(128)
- text text

## データベース,テーブル,ユーザの作成,作成されたユーザへの権限付与

db ディレクトリに`chatdb.sql`が入っているので，MySQL に root でログインした状態で，  
Mac の場合,`source /Users/<your_username>/database_last_issue/db/chatdb.sql;`と入力する．  
Linux の場合,`source /home/<your_username>/database_last_issue/db/chatdb.sql;`と入力する．  
そうすると，データベース，テーブル，ユーザ(user)が作成され，作成されたユーザに権限が付与される．

## 作成されたユーザへのログイン

`mysql -u user -p`と入力し，パスワードを求められたら`hoge`と入力するとログインができる．  
データベース名は`chatdb`なので，`use chatdb`と入力すれば切り替えができる．テーブルは`chat_log, members` が入っており，  
`show columns from chat_log;`と`show columns from members;`を入力し, 各々のカラムが正しければ準備完了．

## Go のインストール

- Mac の場合  
  `brew install go`でインストールを行う．  
  インストール完了後，`.zshrc`に，

  ```
  export GOPATH=$(go env GOPATH)
  export PATH=$PATH:$GOPATH:bin
  ```

  と入力して PATH を通し，  
  `zsh -l`と入力してプロファイルの再読み込みを行う．  
  (macOS 13 Ventura Developer Beta2 で動作確認済み)

- Linux の場合
  `sudo wget https://dl.google.com/go/go1.18.3.linux-amd64.tar.gz`と入力して Go をダウンロードする．  
  `sudo tar -C /usr/local -xzf go1.18.3.linux-amd64.tar.gz`と入力して tar を展開する．  
  `.bash_profile`に，  
  `export PATH=$PATH:/usr/local/go/bin`  
  と入力して PATH を通し，
  `source .bash_profile`と入力してプロファイルの再読み込みを行う．　　
  （AlmaLinux8 で動作確認済み)

## 実行準備

実行する前に，必要なモジュールを取得するために
`go mod tidy`と入力してモジュールのインストールを行う

## 実行

`go run main.go`と入力してローカルサーバを開く
