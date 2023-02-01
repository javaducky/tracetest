PROJECT_ROOT=${PWD}
help: Makefile ## show list of commands
	@echo "Choose a command run:"
	@echo ""
	@awk 'BEGIN {FS = ":.*?## "} /[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-40s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

install-goreleaser:
ifeq (, $(shell goreleaser --version | grep pro))
	@printf "\033[0;31m Tracetest requires goreleaser pro installed (licence not necessary for local builds)\033[0m\n\n"
	@printf "\033[0;33m See https://goreleaser.com/install/ \033[0m\n\n"
	@exit 1
endif

build-web:
	cd web; npm install
	cd web; npm run build

build-docker: build-go build-web
	VERSION=latest \
		goreleaser release --clean --skip-announce --snapshot -f .goreleaser.dev.yaml

build-go: install-goreleaser
	goreleaser build --single-target --clean --snapshot
	@find ./dist -name 'tracetest*' -exec cp {} ./dist 2>/dev/null \;

generate: generate-server generate-cli generate-web

generate-web: ## generates OpenAPI types for WebUI
	cd web; npm run types:generate

OPENAPI_GENERATOR_VER=v5.4.0
OPENAPI_GENERATOR_IMAGE=openapitools/openapi-generator-cli:$(OPENAPI_GENERATOR_VER)
OPENAPI_GENERATOR_CLI=docker run --rm -u ${shell id -u}  -v "$(PROJECT_ROOT):/local" -w "/local" ${OPENAPI_GENERATOR_IMAGE}
OPENAPI_TARGET_DIR=openapi/

generate-cli:
	$(eval BASE := ./cli)
	mkdir -p $(BASE)/tmp
	rm -rf $(BASE)/$(OPENAPI_TARGET_DIR)
	mkdir -p $(BASE)/$(OPENAPI_TARGET_DIR)

	$(OPENAPI_GENERATOR_CLI) generate \
		-i api/openapi.yaml \
		-g go \
		-o $(BASE)/tmp \
		--generate-alias-as-model
	cp $(BASE)/tmp/*.go $(BASE)/$(OPENAPI_TARGET_DIR)
	chmod 644 $(BASE)/$(OPENAPI_TARGET_DIR)/*.go
	rm -rf $(BASE)/tmp

	cd $(BASE); go fmt ./...

generate-server: ## generates OpenAPI types for server
	$(eval BASE := ./server)
	mkdir -p $(BASE)/tmp
	rm -rf $(BASE)/$(OPENAPI_TARGET_DIR)
	mkdir -p $(BASE)/$(OPENAPI_TARGET_DIR)

	$(OPENAPI_GENERATOR_CLI) generate \
		-i api/openapi.yaml \
		-g go-server \
		-o $(BASE)/tmp \
		--generate-alias-as-model
	cp $(BASE)/tmp/go/*.go $(BASE)/$(OPENAPI_TARGET_DIR)
	chmod 644 $(BASE)/$(OPENAPI_TARGET_DIR)/*.go
	rm -f $(BASE)/$(OPENAPI_TARGET_DIR)/api_api_service.go
	rm -rf $(BASE)/tmp

	cd $(BASE); go fmt ./...

serve-docs: ## serve documentation for Tracetest
	docker build -t tracetest-docs -f docs-Dockerfile .
	docker run --network host tracetest-docs
	sleep 1
	open http://localhost:8000
