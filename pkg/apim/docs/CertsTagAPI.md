# \CertsTagAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCertsWithIDs**](CertsTagAPI.md#ListCertsWithIDs) | **Get** /tyk/certs/{certID} | Return one certificate or list multiple certificates in the Tyk Gateway given a comma separated list of cert IDs.



## ListCertsWithIDs

> ListCertsWithIDs200Response ListCertsWithIDs(ctx, certID).Execute()

Return one certificate or list multiple certificates in the Tyk Gateway given a comma separated list of cert IDs.



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
	certID := "e6ce2b49-3e31-44de-95a7-12f054724283,5e9d9544a1dcd60001d0ed20a6ab77653d5da938f452bb8cc9b55b0630a6743dabd8dc92bfb025abb09ce035" // string | Comma separated list of certificates to list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertsTagAPI.ListCertsWithIDs(context.Background(), certID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertsTagAPI.ListCertsWithIDs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCertsWithIDs`: ListCertsWithIDs200Response
	fmt.Fprintf(os.Stdout, "Response from `CertsTagAPI.ListCertsWithIDs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certID** | **string** | Comma separated list of certificates to list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCertsWithIDsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListCertsWithIDs200Response**](ListCertsWithIDs200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

