IMAGE_NAME = 'auth-service'
USER = root
PASSWORD = secret
DEV_PASSWORD = RORCqbTq3qpzNGXCQeJh
DEV_HOST = simple-bank.cziqss8y0sci.us-east-2.rds.amazonaws.com
PORT = 5432
DB_NAME = auth_service

createpostgres:
	docker run --name $(IMAGE_NAME)  -e POSTGRES_USER=$(USER) -e POSTGRES_PASSWORD=$(PASSWORD) -p $(PORT):$(PORT) -d postgres

createdb:
	docker exec -it $(IMAGE_NAME) createdb --username=$(USER) --owner=$(USER) $(DB_NAME)

dropdb:
	docker exec -it $(IMAGE_NAME) dropdb $(DB_NAME)

stoppostgres:
	docker stop $(IMAGE_NAME)

runpostgres:
	docker start $(IMAGE_NAME)

deletepostgres:
	docker rm $(IMAGE_NAME)

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

migrateup:
	migrate -path db/migration -database "postgresql://$(USER):$(PASSWORD)@localhost:$(PORT)/$(DB_NAME)?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://$(USER):$(PASSWORD)@localhost:$(PORT)/$(DB_NAME)?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://$(USER):$(PASSWORD)@localhost:$(PORT)/$(DB_NAME)?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://$(USER):$(PASSWORD)@localhost:$(PORT)/$(DB_NAME)?sslmode=disable" -verbose down 1

migrateupDEV:
	migrate -path db/migration -database "postgresql://$(USER):$(DEV_PASSWORD)@$(DEV_HOST):$(PORT)/$(DB_NAME)" -verbose up

sqlc_generate:
	docker run --rm -v $(PWD):/src -w /src kjconroy/sqlc generate

sqlc_init:
	docker run --rm -v $(PWD):/src -w /src kjconroy/sqlc init

db_docs:
	dbdocs build doc/db.dbml

protob: # make sure statik is installed and this is ran as an admin (run in bash)
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
           --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
           --grpc-gateway_out=pb \
           --grpc-gateway_opt paths=source_relative \
           --openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
           proto/*.proto
	statik -src=./doc/swagger -dest=./doc

redis:
	docker run --name redis -p 6379:6379 -d redis:8.0-M02-alpine

PHONY: createpostgres createdb dropdb stoppostgres runpostgres deletepostgres new_migration migrateup migrateup1 migratedown migratedown1 migrateupDEV sqlc_generate sqlc_init db_docs protob redis