# database_last_issue

##### ※ Linux,Macを使用している前提  

## テーブルの概要

### membersテーブル
- id bigint not null auto_increment
- username varchar(128)
- password varchar(128)
- birthday char(10)


### chat_logテーブル
- id bigint not null auto_increment
- username varchar(128)
- text text


### azureapiテーブル
- subkey varchar(128)
- location varchar(16)
- endpoint varchar(128)
- uri varchar(32)


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
- Macの場合  
`brew install go`でインストールを行う．    
インストール完了後，`.zshrc`に，  
```
export GOPATH=$(go env GOPATH)
export PATH=$PATH:$GOPATH:bin
```
と入力してPATHを通し，  
`zsh -l`と入力してプロファイルの再読み込みを行う．  
(macOS Monterey 12.5 Beta5で動作確認済み)    

- Linuxの場合
`sudo wget https://dl.google.com/go/go1.18.3.linux-amd64.tar.gz`と入力してGoをダウンロードする．  
`sudo tar -C /usr/local -xzf go1.18.3.linux-amd64.tar.gz`と入力してtarを展開する．  
`.bash_profile`に，  
`export PATH=$PATH:/usr/local/go/bin`  
と入力してPATHを通し，
`source .bash_profile`と入力してプロファイルの再読み込みを行う．　　
（AlmaLinux8で動作確認済み)  


## 実行準備

実行する前に，必要なモジュールを取得するために
`go mod tidy`と入力してモジュールのインストールを行う

## 実行

`go run main.go`と入力してローカルサーバを開く
