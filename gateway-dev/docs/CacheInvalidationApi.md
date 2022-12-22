# \CacheInvalidationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvalidateCache**](CacheInvalidationApi.md#InvalidateCache) | **Delete** /tyk/cache/{apiID} | Invalidate cache



## InvalidateCache

> ApiStatusMessage InvalidateCache(ctx, apiID).Execute()

Invalidate cache



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    apiID := "apiID_example" // string | The API ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CacheInvalidationApi.InvalidateCache(context.Background(), apiID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CacheInvalidationApi.InvalidateCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvalidateCache`: ApiStatusMessage
    fmt.Fprintf(os.Stdout, "Response from `CacheInvalidationApi.InvalidateCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID | 

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

