run:
	go run cmd/*

migrate-up:
	 migrate -path db/migration/ -database "mysql://root:123456@tcp(localhost:3343)/nitflex" -verbose up

migrate-down:
	 migrate -path db/migration/ -database "mysql://root:123456@tcp(localhost:3343)/nitflex" -verbose down

migrate:
	migrate create -ext sql -dir  db/migration/ -seq $(name)


.PHONY: run, migrate_up, migrate_down, new_migration