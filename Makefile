BUILD_PATH = ./build
PRODUCT_NAME = cxsast_exporter
PRODUCT_VERSION = $(shell cat VERSION)
PRODUCT_BUILD = $(shell date +%Y%m%d%H%M%S)
PUBLIC_KEY = "internal/app/encryption/public.key"
LD_FLAGS = -ldflags="-s -w -X github.com/checkmarxDev/ast-sast-export/cmd.productName=$(PRODUCT_NAME) -X github.com/checkmarxDev/ast-sast-export/cmd.productVersion=$(PRODUCT_VERSION) -X github.com/checkmarxDev/ast-sast-export/cmd.productBuild=$(PRODUCT_BUILD)"

SAST_EXPORT_USER = '###########'
SAST_EXPORT_PASS = '###########'

lint:
	go fmt ./...
	golangci-lint run

build: windows_amd64 #windows_386 linux_amd64 linux_386 darwin_amd64

run: windows_amd64 run_windows

debug: windows_amd64 debug_windows

unit_test:
	go test -short $(LD_FLAGS) ./... -coverprofile=coverage.out

clean:
	rm -r $(BUILD_PATH)

windows_amd64: check_public_key
	env GOOS=windows GOARCH=amd64 go build -o $(BUILD_PATH)/windows/amd64/$(PRODUCT_NAME).exe $(LD_FLAGS)

#windows_386: check_public_key
#	env GOOS=windows GOARCH=386 go build -o $(BUILD_PATH)/windows/386/$(PRODUCT_NAME).exe $(LD_FLAGS)

#linux_amd64: check_public_key
#	env GOOS=linux GOARCH=amd64 go build -o $(BUILD_PATH)/linux/amd64/$(PRODUCT_NAME) $(LD_FLAGS)

#linux_386: check_public_key
#	env GOOS=linux GOARCH=386 go build -o $(BUILD_PATH)/linux/386/$(PRODUCT_NAME) $(LD_FLAGS)

#darwin_amd64: check_public_key
#	env GOOS=darwin GOARCH=amd64 go build -o $(BUILD_PATH)/darwin/amd64/$(PRODUCT_NAME) $(LD_FLAGS)

public_key:
	aws kms get-public-key --key-id alias/sast-migration-key --region eu-west-1 | jq -r .PublicKey > $(PUBLIC_KEY)

check_public_key:
	if [ ! -f $(PUBLIC_KEY) ]; then echo "Please run: make public_key"; exit 1; fi

run_windows:
	build/windows/amd64/cxsast_exporter --user $(SAST_EXPORT_USER) --pass $(SAST_EXPORT_PASS) --url http://localhost --export users,results,teams --results-project-active-since 1

debug_windows:
	build/windows/amd64/cxsast_exporter --user $(SAST_EXPORT_USER) --pass $(SAST_EXPORT_PASS) --url http://localhost --export users,results,teams --results-project-active-since 10 --debug

mocks:
	rm -rf test/mocks
	mockgen -package mock_integration_rest -destination test/mocks/integration/rest/mock_client.go github.com/checkmarxDev/ast-sast-export/internal/integration/rest Client
	mockgen -package mock_integration_soap -destination test/mocks/integration/soap/mock_adapter.go github.com/checkmarxDev/ast-sast-export/internal/integration/soap Adapter
	mockgen -package mock_integration_similarity -destination test/mocks/integration/similarity/provider_mock.go github.com/checkmarxDev/ast-sast-export/internal/integration/similarity SimilarityIDProvider
	mockgen -package mock_persistence_ast_query_id -destination test/mocks/persistence/ast_query_id/provider_mock.go github.com/checkmarxDev/ast-sast-export/internal/persistence/ast_query_id QueryIDProvider
	mockgen -package mock_persistence_source -destination test/mocks/persistence/source/mock_provider.go github.com/checkmarxDev/ast-sast-export/internal/persistence/source SourceProvider
	mockgen -package mock_persistence_method_line -destination test/mocks/persistence/method_line/mock_provider.go github.com/checkmarxDev/ast-sast-export/internal/persistence/method_line Provider
	mockgen -package mock_app_export -destination test/mocks/app/export/mock_exporter.go github.com/checkmarxDev/ast-sast-export/internal/app/export Exporter
	mockgen -package mock_app_metadata -destination test/mocks/app/metadata/mock_provider.go github.com/checkmarxDev/ast-sast-export/internal/app/metadata MetadataProvider
