sdk:
	 docker run --rm -v "${PWD}:/local" swaggerapi/swagger-codegen-cli-v3 generate \
        -i /local/swagger.yml \
        --additional-properties=packageName=gateway \
        -l go \
        -o /local/gateway

gen:
	oapi-codegen -generate types,client -package spec swagger.yaml > ./oapi/gen.go





