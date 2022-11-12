include .env
.SHELLFLAGS += -e # discontinue whenever a line fails

POSTGRES_CMD=psql postgresql://${POSTGRES_USERNAME}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}

all: sqlc drop-all create mock api

api:
	cd cmd/api && go clean && go build

drop-all:
	@$(POSTGRES_CMD) -f sql/drop-all.sql

create:
	@$(POSTGRES_CMD) -f sql/schema.sql

mock:
	@$(POSTGRES_CMD) -f sql/mock.sql

sqlc:
	sqlc generate