include .env

empty:
	echo "empty"

# 開発環境のdocker compose コマンド
dcb-dev:
	docker compose build
dcu-dev:
	docker compose up -d
dcd-dev:
	docker compose down

# コンテナ環境へsshログイン
backend-ssh:
	docker exec -it ${BACKEND_CONTAINER_NAME} sh


# DB関連
## マイグレーション
db-migrate:
	docker exec -it ${BACKEND_CONTAINER_NAME} sh -c "go run db/migrate/migrate.go"
db-seed:
	docker exec -it ${BACKEND_CONTAINER_NAME} sh -c "go run db/seeds/seed.go"


# ローカル開発用
# go library install
## 複数のライブラリを指定する場合は、name="xxx yyy" のように""で囲んで実行すること
go-add-library:
	docker exec -it ${BACKEND_CONTAINER_NAME} sh -c "go get ${name}"