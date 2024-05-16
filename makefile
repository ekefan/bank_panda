postgres:
	docker run --name postgresSimpleBank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d  postgres

createdb:
	docker exec postgresSimpleBank createdb --username=root --owner=root simpleBank

dropdb:
	docker exec postgresSimpleBank dropdb simpleBank

migrateup:
	migrate -path ./db/migrations -database "postgresql://root:secret@localhost:5432/simpleBank?sslmode=disable" -verbose up
migrateup1:
	migrate -path ./db/migrations -database "postgresql://root:secret@localhost:5432/simpleBank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path ./db/migrations -database "postgresql://root:secret@localhost:5432/simpleBank?sslmode=disable" -verbose down

migratedown1:
	migrate -path ./db/migrations -database "postgresql://root:secret@localhost:5432/simpleBank?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate

testdb:
	go test -v -cover ./db/sqlc

server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/sqlc/mockStore.go github.com/ekefan/bank_panda/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc testdb server migratedown1 migrateup1