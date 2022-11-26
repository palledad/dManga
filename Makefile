SCHEMA_YAML = openapi/openapi.yaml

backend-api-interface: $(SCHEMA_YAML)
	docker compose run openapi-generator generate -g go-server --additional-properties=outputAsLibrary=true -i home/openapi.yaml -o home/backend
	docker compose run backend goimports -w ./

frontend-api-interface: $(SCHEMA_YAML)
	docker compose run openapi-generator generate -g typescript-axios -i home/openapi.yaml -o home/frontend/api
