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

PHONY: createpostgres createdb dropdb stoppostgres runpostgres deletepostgres new_migration migrateup migrateup1 migratedown migratedown1 migrateupDEV sqlc_generate sqlc_init