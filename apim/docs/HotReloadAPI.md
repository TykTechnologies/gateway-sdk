# \HotReloadAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HotReload**](HotReloadAPI.md#HotReload) | **Get** /tyk/reload/ | Hot-reload a single node
[**HotReloadGroup**](HotReloadAPI.md#HotReloadGroup) | **Get** /tyk/reload/group | Hot-reload a Tyk group



## HotReload

> ApiStatusMessageModelModel HotReload(ctx).Block(block).Execute()

Hot-reload a single node



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
    block := true // bool | Block a response until the reload is performed. This can be useful in scripting environments like CI/CD workflows. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HotReloadAPI.HotReload(context.Background()).Block(block).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HotReloadAPI.HotReload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HotReload`: ApiStatusMessageModelModel
    fmt.Fprintf(os.Stdout, "Response from `HotReloadAPI.HotReload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHotReloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **block** | **bool** | Block a response until the reload is performed. This can be useful in scripting environments like CI/CD workflows. | 

### Return type

[**ApiStatusMessageModelModel**](ApiStatusMessageModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HotReloadGroup

> ApiStatusMessageModelModel HotReloadGroup(ctx).Execute()

Hot-reload a Tyk group



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HotReloadAPI.HotReloadGroup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HotReloadAPI.HotReloadGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HotReloadGroup`: ApiStatusMessageModelModel
    fmt.Fprintf(os.Stdout, "Response from `HotReloadAPI.HotReloadGroup`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHotReloadGroupRequest struct via the builder pattern


### Return type

[**ApiStatusMessageModelModel**](ApiStatusMessageModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

