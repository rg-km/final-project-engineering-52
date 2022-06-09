start: 
	go run main.go

migrate-up:
	migrate -path db/migration -database "sqlite3://db/go_scholarship.db?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migration -database "sqlite3://db/go_scholarship.db?sslmode=disable" -verbose down

seed-all:
	go run main.go seed