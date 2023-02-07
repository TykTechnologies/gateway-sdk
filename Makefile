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



gateway-sdks:
	 docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v6.2.0 generate \
         --git-user-id TykTechnologies --git-repo-id gateway-sdk \
         --package-name gate \
         --api-name-suffix API \
          --model-name-suffix Model \
         --additional-properties=isGoSubmodule=false,generateInterfaces=true,hideGenerationTimestamp=false,outputAsLibrary=true \
        -i /local/swagger.yml \
        -g go \
        -o /local/gate


