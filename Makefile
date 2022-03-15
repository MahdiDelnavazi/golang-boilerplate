postgres:
	docker run --name supnext_pr2_image -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it supnext_pr2_image createdb --username=root --owner=root supnext_pr2

dropdb:
	docker exec -it supnext_pr2_image dropdb supnext_pr2

migrateup:
	migrate -path database/schema -database "postgresql://root:secret@localhost:5432/supnext_pr2?sslmode=disable" -verbose up

migratedown:
	migrate -path database/schema -database "postgresql://root:root@localhost:5432/supnext_pr2?sslmode=disable" -verbose down


redis:
	docker run --name supnex_pr2-redis -p 6379:6379 -d redis

test:
	go test -v -cover ./...

server:
	go run main.go


.PHONY: postgres createdb dropdb migrateup migratedown test server redis