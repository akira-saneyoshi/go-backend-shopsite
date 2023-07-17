---
Golang: 1.20.6
Framework: Fiber
OS: Windows11 Pro
Docker for Desktop: 4.21.1
---

### Installation

```bash
go mod init go-shopsite
```

```bash
go get package main
```

```bash
go get github.com/gofiber/fiber/v2
```

### Hello, World!

```bash
go run main.go
```

### Dockerコンテナの起動

```bash
docker-compose up -d
```

### データベースの作成

Dockerのコンテナ初期構築時に以下のエラーが出た場合は、MySQLのコンテナに入って次のコマンドを実行する。（原因解消は後ほど。）

```bash
CREATE DATABASE shop_site;
```

```bash
go-backend-backend-1  | panic: Cloud not connect with the database!
go-backend-backend-1  |
go-backend-backend-1  | goroutine 1 [running]:
go-backend-backend-1  | main.main()
go-backend-backend-1  |     /app/main.go:13 +0x178
go-backend-backend-1  | exit status 2
go-backend-backend-1  |
go-backend-backend-1  | YYYY/mm/dd ~ /app/main.go:10
go-backend-backend-1  | [error] failed to initialize database, got error dial tcp 172.23.0.2:3306: connect: connection refused
go-backend-backend-1  | panic: Cloud not connect with the database!
go-backend-backend-1  |
```
