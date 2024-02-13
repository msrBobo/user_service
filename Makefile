CURRENT_DIR=$(shell pwd)
DB_URL := "postgres://postgres:bobo@localhost:5432/userdb?sslmode=disable"

proto-gen:
	chmod +x ./scripts/genproto.sh
	./scripts/genproto.sh
migrate-up:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" -verbose down

migrate-force:
	migrate -path migrations -database "$(DB_URL)" -verbose force 1

migrate-file:
	migrate create -ext sql -dir migrations/ -seq create_comments_table
