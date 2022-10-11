EXEC=docker-compose exec api

wire: install/wire
	${EXEC} sh -c "wire wire.go"

oapigen: install/oapigen
	${EXEC} sh -c "sh ./controller/openapi/makeinterface.sh"

install/wire:
	${EXEC} go install github.com/google/wire/cmd/wire

install/oapigen:
	${EXEC} go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0

migrate:
	${EXEC} go run ./infrastructure/postgres/migration/main.go
