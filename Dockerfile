FROM golang:1.17-alpine

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN export CGO_ENABLED=0 \
    && go mod download \
    && go build -o app main/main.go

EXPOSE 9000

CMD ["./app"]

   