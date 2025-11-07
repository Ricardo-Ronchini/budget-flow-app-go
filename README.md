# budget-flow-app-go

Handler API: Echo
Postgres: banco de dados
Cache: Redis

#Run
## Create Database
- docker compose -f ./scripts/docker-compose.yml -p budget-app up -d

## Migrations
- go run ./cmd/migrations

## API
- go run ./cmd/api
