# \SchemaAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSchema**](SchemaAPI.md#GetSchema) | **Get** /tyk/schema | Get OAS schema.



## GetSchema

> OASSchemaResponse GetSchema(ctx).OasVersion(oasVersion).Execute()

Get OAS schema.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TykTechnologies/gateway-sdk/pkg/apim"
)

func main() {
	oasVersion := "3.0.3" // string | The OAS version to fetch. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchemaAPI.GetSchema(context.Background()).OasVersion(oasVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchemaAPI.GetSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSchema`: OASSchemaResponse
	fmt.Fprintf(os.Stdout, "Response from `SchemaAPI.GetSchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oasVersion** | **string** | The OAS version to fetch. | 

### Return type

[**OASSchemaResponse**](OASSchemaResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

