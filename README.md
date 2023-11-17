# go-docker-restapi-sample

golang docker REST API のサンプル

## 技術構成

- go
- gorm
- gorilla
- crypto
- godotenv
- mysql
- jwt

## API

### トップ

|                                          | メソッド | URI     | 権限 |
| :--------------------------------------- | :------- | :------ | :--- |
| 文字列を返すだけの確認用のエンドポイント | GET      | /api/v1 | なし |

### ユーザー認証

|                                          | メソッド | URI     | 権限 |
| :--------------------------------------- | :------- | :------ | :--- |
| ログイン | POST      | /api/v1/signin | なし |
| 会員登録 | POST      | /api/v1/signup | なし |

### Todo

|                                          | メソッド | URI     | 権限 |
| :--------------------------------------- | :------- | :------ | :--- |
| ユーザーに紐づく全Todoデータを取得 | GET      | /api/v1/todo | 認証済 |
| TodoのIDに紐づく単一のTodoデータを取得  | GET      | /api/v1/todo/:id | 認証済 |
| Todo新規作成  | POST      | /api/v1/todo | 認証済 |
| Todo更新  | PUT      | /api/v1/todo/:id | 認証済 |
| Todo削除  | DELETE      | /api/v1/todo/:id | 認証済 |

## 環境構築

### 1. env ファイルを作成

- ルートディレクトリ直下に「.env」ファイルを作成
- 「.env.sample」の記述をコピー

```
touch .env
```

- app ディレクトリに移動し、「.env」ファイルを作成
- 「app/.env.sample」ファイルの記述をコピー

```
cd app
touch .env
```

### 2. docker 起動

- ビルド

```
docker compose build
```

- コンテナ起動

```
docker compose up
```

- go のコンテナにアクセス

```
make backend-ssh
```

### 3. データ用意

- マイグレーションを実行
- go のコンテナ内で、以下コマンドを実行する

```
make db-migrate
```

- テーブルが作成されるので、DB(MySQL)に接続し確認する

  - 接続アプリは「Sequel Ace」がおすすめ
  - https://qiita.com/ucan-lab/items/b1304eee2157dbef7774

- シーディングを実行
- go のコンテナ内で、以下コマンドを実行する

```
make db-seed
```

- データが作成されるので、DB(MySQL)に接続し確認する

### 4. API 起動

- DBにデータを作成したので、再度dockerをリスタートしてAPIを起動させる (初回のみ)

```
docker compose restart
```

- 以下の url に接続し、レスポンスが返ってくる事を確認
  - http://localhost:4000/api/v1

## 開発中のコマンド
※何もdockerでAPIを起動した状態で実行してください。

### テスト
```
make test
```

### 静的解析
```
make lint
```

### goのライブラリ追加
```
go-add-library name="[ライブラリ名]"

// 複数のライブラリを指定する場合は、name="xxx yyy" のように""で囲んで実行すること
```

### DBのデータをリセットする場合
gormはロールバック機能がないため、以下のコマンドでDBごと消去する
```
docker compose down -v
```

その後dockerを起動させて再度マイグレーションをしてテーブルを初期化する

## docker コマンド

```
// ビルド
docker-compose build

// コンテナ起動
docker-compose up

// コンテナ起動(バックグラウンド実行)
docker-compose up -d

// コンテナ停止
docker-compose down

// コンテナ停止&ボリューム削除(DBデータを削除)
docker-compose down -v

// appコンテナへログイン
docker exec -it 20211105_go_rest_server sh

// dbコンテナへログイン
docker exec -it 20211105_go_rest_db /bin/bash

```
