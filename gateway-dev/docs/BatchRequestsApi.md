# \BatchRequestsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Batch**](BatchRequestsApi.md#Batch) | **Post** /{listen_path}/tyk/batch | Run batch request



## Batch

> Batch(ctx, listenPath).Execute()

Run batch request

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
    listenPath := "listenPath_example" // string | API listen path

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchRequestsApi.Batch(context.Background(), listenPath).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchRequestsApi.Batch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listenPath** | **string** | API listen path | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

