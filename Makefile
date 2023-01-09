codegen-sdk:
	 docker run --rm -v "${PWD}:/local" swaggerapi/swagger-codegen-cli-v3 generate \
        -i /local/swagger.yml \
        --additional-properties=packageName=gateway \
        -l go \
        -o /local/gateway

gen:
	oapi-codegen -generate types,client -package spec swagger.yaml > ./oapi/gen.go



gateway-sdks:
	 docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v6.2.0 generate \
         --git-user-id TykTechnologies --git-repo-id gateway-sdk \
         --package-name gate \
         --additional-properties=isGoSubmodule=false,hideGenerationTimestamp=false,outputAsLibrary=true \
        -i /local/swagger.yml \
        -g go \
        -o /local/gateway-dev


