SCHEMA_YAML = openapi/openapi.yaml
OAPI_CONFIG_YAML = openapi/gen.yaml

BACKEND_GEN_FILE = backend/internal/gateways/controller/gen/openapi.go
FRONTEND_GEN_FILE = frontend/api/api.ts

.PHONY: openapi-gen app swagger-ui format-fe

app: openapi-gen
	docker compose up --build backend frontend

### frontned ###
frontend-server: $(FRONTEND_GEN_FILE)
	docker compose up frontend

format-fe:
	docker compose run frontend pnpm format

### backend ###
backend-server: $(BACKEND_GEN_FILE)
	docker compose up backend

backend-lint:
	docker compose run backend golangci-lint run ./...

### openapi-generator ###
openapi-gen: $(FRONTEND_GEN_FILE) $(BACKEND_GEN_FILE)

$(FRONTEND_GEN_FILE): $(SCHEMA_YAML)
	docker compose run openapi-generator generate -g typescript-axios -i home/openapi.yaml -o home/frontend/api

$(BACKEND_GEN_FILE): $(SCHEMA_YAML) $(OAPI_CONFIG_YAML)
	docker compose run backend oapi-codegen --config /openapi/gen.yaml /openapi/openapi.yaml
	docker compose run backend goimports -w ./

### swagger ###
swagger-ui:
	docker compose up swagger-ui