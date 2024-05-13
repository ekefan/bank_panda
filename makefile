postgres:
	docker run --name postgresSimpleBank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d  postgres

createdb:
	docker exec postgresSimpleBank createdb --username=root --owner=root simpleBank

dropdb:
	docker exec postgresSimpleBank dropdb simpleBank

migrateup:
	migrate -path ./db/migrations -database "postgresql://root:secret@localhost:5432/simpleBank?sslmode=disable" -verbose up

migratedown:
	migrate -path ./db/migrations -database "postgresql://root:secret@localhost:5432/simpleBank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

testdb:
	go test -v -cover ./db/sqlc


.PHONY: postgres createdb dropdb migrateup migratedown sqlc testdbgi