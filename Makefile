apim-sdk:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v7.10.0 generate \
         --generator-name go \
        --input-spec /local/swagger.yml \
        --output /local/pkg/apim/ \
        --skip-overwrite \
        --git-host github.com \
        --git-user-id TykTechnologies \
         --git-repo-id gateway-sdk/pkg/apim \
         --package-name apim \
         --api-name-suffix API \
         --minimal-update \
        --global-property skipFormModel=true \
        --global-property skipFormModel=true \
        --global-property apis,apiTests=false,apiDocs=true \
        --global-property models,modelTests=true,modelDocs=true \
        --global-property supportingFiles \
        -c /local/config.json\
        --name-mappings _id=MID

	sudo rm -rf pkg/apim/model_server_variable.go
	sudo gofmt -s -w .
	go mod tidy

validate-swagger:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v7.10.0 validate \
		-i /local/swagger.yml





