postgres:
	docker run --name simpleBank --network bank-net -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec simpleBank createdb --username=root --owner=root bank_panda

dropdb:
	docker exec simpleBank dropdb bank_panda

migrateup:
	migrate -path ./db/migrations -database "postgresql://root:9hafAZzagLg3l6zN5HG1@simple-bank.cp2oqo4yuqpu.eu-west-2.rds.amazonaws.com:5432/bankpanda" -verbose up
migrateup1:
	migrate -path ./db/migrations -database "postgresql://root:9hafAZzagLg3l6zN5HG1@simple-bank.cp2oqo4yuqpu.eu-west-2.rds.amazonaws.com:5432/bankpanda" -verbose up 1

migratedown:
	migrate -path ./db/migrations -database "postgresql://root:9hafAZzagLg3l6zN5HG1@simple-bank.cp2oqo4yuqpu.eu-west-2.rds.amazonaws.com:5432/bankpanda" -verbose down

migratedown1:
	migrate -path ./db/migrations -database "postgresql://root:9hafAZzagLg3l6zN5HG1@simple-bank.cp2oqo4yuqpu.eu-west-2.rds.amazonaws.com:5432/bankpanda" -verbose down 1
sqlc:
	sqlc generate

testdb:
	go test -v -cover ./db/sqlc

server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/sqlc/mockStore.go github.com/ekefan/bank_panda/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc testdb server migratedown1 migrateup1   