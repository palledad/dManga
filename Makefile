SCHEMA_YAML = openapi/openapi.yaml

backend-api-interface: $(SCHEMA_YAML)
	docker compose run openapi-generator generate -g go-gin-server -i home/openapi.yaml -o home/backend

frontend-api-interface: $(SCHEMA_YAML)
	docker compose run openapi-generator generate -g typescript-axios -i home/openapi.yaml -o home/frontend/api
