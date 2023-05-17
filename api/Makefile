EXEC=docker-compose exec api

wire: install/wire
	${EXEC} sh -c "wire wire.go"

oapigen: install/oapigen
	${EXEC} sh -c "sh ./controller/openapi/makeinterface.sh"

create_migration: install/goose
	${EXEC} sh -c "cd migrations && goose create ${NAME} sql"

migrate: install/goose
	${EXEC} sh -c 'cd migrations && goose postgres "host=db port=5432 user=root password=password dbname=test sslmode=disable" up'

install/wire:
	${EXEC} go install github.com/google/wire/cmd/wire

install/oapigen:
	${EXEC} go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0

install/goose:
	${EXEC} go install github.com/pressly/goose/v3/cmd/goose@latest
