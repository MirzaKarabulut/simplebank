postgres:
	docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

resetdb: dropdb createdb migrateup

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

run:
	go run main.go

mock:
	mockgen -package mockdb  -destination db/mock/store.go  github.com/MirzaKarabulut/simplebank/db/sqlc Store

.PHONY: createdb dropdb postgres migrateup migratedown migrateup1 migratedown1 resetdb sqlc test server mock