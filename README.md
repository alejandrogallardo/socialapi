```
go mod init github.com/nombre/repositorio 
go get -u github.com/go-chi/chi/v5
go install github.com/air-verse/air@latest
air init

go get github.com/joho/godotenv

https://threedots.tech/post/repository-pattern-in-go/
	
go get github.com/lib/pq

migrate create -seq -ext sql -dir ./cmd/migrate/migrations create_users

https://www.dataquest.io/blog/sql-joins/

docker compose up -d
docker compose down

ADDR=:3000
DB_ADDR=postgres://admin:adminpassword@localhost:5433/postgres?sslmode=disable
DB_MAX_OPEN_CONNS=30
DB_MAX_IDLE_CONNS=30
DB_MAX_IDLE_TIME=15m

run api
air
```