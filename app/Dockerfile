FROM golang:1.21.4-alpine3.18

#ENV CGO_ENABLED 0

RUN mkdir /app
# ワーキングディレクトリの設定
WORKDIR /app

COPY /app /app

RUN go mod tidy && \
    go install github.com/cosmtrek/air@v1.46.0

CMD ["air", "-c", ".air.toml"]
