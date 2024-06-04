db/migrate:
	go run ./cmd/migrate

local-db:
	docker-compose --env-file ./.env -f ./docker-compose.yml down
	docker-compose --env-file ./.env -f ./docker-compose.yml up -d

run:
	air -c .air.toml