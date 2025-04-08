# \CacheInvalidationAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvalidateCache**](CacheInvalidationAPI.md#InvalidateCache) | **Delete** /tyk/cache/{apiID} | Invalidate cache.



## InvalidateCache

> ApiStatusMessage InvalidateCache(ctx, apiID).Execute()

Invalidate cache.



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
	apiID := "ae67bb862a3241a49117508e0f9ee839" // string | The API ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CacheInvalidationAPI.InvalidateCache(context.Background(), apiID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CacheInvalidationAPI.InvalidateCache``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvalidateCache`: ApiStatusMessage
	fmt.Fprintf(os.Stdout, "Response from `CacheInvalidationAPI.InvalidateCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvalidateCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiStatusMessage**](ApiStatusMessage.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

