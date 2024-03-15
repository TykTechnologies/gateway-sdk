apim-sdk:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    	--generator-name go \
        --input-spec /local/open.yaml \
        --output /local/pkg/apim/ \
        --skip-overwrite \
        --git-host github.com \
        --git-user-id TykTechnologies \
        --git-repo-id gateway-sdk/pkg/apim \
        --package-name apim \
        --api-name-suffix API \
        --global-property skipFormModel=true \
        --global-property apis,apiTests=false,apiDocs=false \
        --global-property models,modelTests=false,modelDocs=false \
        --global-property supportingFiles \
        --additional-properties generateInterfaces=true

codegen-sdk:
	 docker run --rm -v "${PWD}:/local" swaggerapi/swagger-codegen-cli-v3:3.0.22 generate \
        -i /local/swagger.yml \
        --additional-properties=packageName=gateway \
        -l go \
        -o /local/gateway
rm-gateway:
	bash rem.sh

gen:
	oapi-codegen -generate types,client -package spec swagger.yml > ./oapi/gen.go


validate-swagger:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v6.2.0 validate \
		-i /local/swagger.yml
gateway-sdks:
	 docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v6.2.0 generate \
         --type-mappings APIDefinitionCORSModelModel=APIDefinitionCORSModel \
         --type-mappings AuthModelModel=AuthModel \
         --type-mappings AuthProviderMetaModelModel=AuthProviderMetaModel \
         --type-mappings APIDefinitionBasicAuthModelModel=APIDefinitionBasicAuthModel \
         --type-mappings APIDefinitionDefinitionModelModel=APIDefinitionDefinitionModel \
         --type-mappings EventHandlerMetaConfigModelModel=EventHandlerMetaConfigModel \
         --type-mappings SignatureConfigModelModel=SignatureConfig \
         --type-mappings CacheOptionsModelModel=CacheOptionsModel \
         --type-mappings MiddlewareSectionModelModel=MiddlewareSectionModel \
         --type-mappings GlobalRateLimitModelModel=GlobalRateLimitModel \
         --type-mappings NotificationsManagerModelModel=NotificationsManager \
         --type-mappings APIDefinitionOauthMetaModelModel=APIDefinitionOauthMetaModel \
         --type-mappings OpenIDOptionsModelModel=OpenIDOptionsModel \
         --type-mappings APIDefinitionProxyModelModel=APIDefinitionProxyModel \
         --type-mappings ResponseProcessorModelModel=ResponseProcessorModel \
         --type-mappings SessionProviderMetaModelModel=SessionProviderMetaModel\
         --git-user-id TykTechnologies --git-repo-id gateway-sdk \
         --package-name apim \
         --api-name-suffix API \
          --model-name-suffix Model \
         --additional-properties=isGoSubmodule=false,generateInterfaces=true,hideGenerationTimestamp=false,outputAsLibrary=true \
        -i /local/open.yaml \
        -g go \
        -o /local/apim



