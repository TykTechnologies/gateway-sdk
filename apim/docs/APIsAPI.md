# \APIsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApi**](APIsAPI.md#CreateApi) | **Post** /tyk/apis | 
[**DeleteApi**](APIsAPI.md#DeleteApi) | **Delete** /tyk/apis/{apiID} | 
[**GetApi**](APIsAPI.md#GetApi) | **Get** /tyk/apis/{apiID} | 
[**ListApis**](APIsAPI.md#ListApis) | **Get** /tyk/apis | 
[**UpdateApi**](APIsAPI.md#UpdateApi) | **Put** /tyk/apis/{apiID} | 



## CreateApi

> ApiModifyKeySuccessModelModel CreateApi(ctx).APIDefinitionModel(aPIDefinitionModel).Execute()





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
    aPIDefinitionModel := *openapiclient.NewAPIDefinitionModel() // APIDefinitionModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIsAPI.CreateApi(context.Background()).APIDefinitionModel(aPIDefinitionModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.CreateApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApi`: ApiModifyKeySuccessModelModel
    fmt.Fprintf(os.Stdout, "Response from `APIsAPI.CreateApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aPIDefinitionModel** | [**APIDefinitionModel**](APIDefinitionModel.md) |  | 

### Return type

[**ApiModifyKeySuccessModelModel**](ApiModifyKeySuccessModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApi

> ApiStatusMessageModelModel DeleteApi(ctx, apiID).Execute()





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
    resp, r, err := apiClient.APIsAPI.DeleteApi(context.Background(), apiID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.DeleteApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteApi`: ApiStatusMessageModelModel
    fmt.Fprintf(os.Stdout, "Response from `APIsAPI.DeleteApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetApi

> APIDefinitionModelModel GetApi(ctx, apiID).Execute()





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
    resp, r, err := apiClient.APIsAPI.GetApi(context.Background(), apiID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.GetApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApi`: APIDefinitionModelModel
    fmt.Fprintf(os.Stdout, "Response from `APIsAPI.GetApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**APIDefinitionModelModel**](APIDefinitionModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApis

> []APIDefinitionModelModel ListApis(ctx).Execute()





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
    resp, r, err := apiClient.APIsAPI.ListApis(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.ListApis``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApis`: []APIDefinitionModelModel
    fmt.Fprintf(os.Stdout, "Response from `APIsAPI.ListApis`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApisRequest struct via the builder pattern


### Return type

[**[]APIDefinitionModelModel**](APIDefinitionModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApi

> ApiModifyKeySuccessModelModel UpdateApi(ctx, apiID).APIDefinitionModel(aPIDefinitionModel).Execute()





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
    aPIDefinitionModel := *openapiclient.NewAPIDefinitionModel() // APIDefinitionModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIsAPI.UpdateApi(context.Background(), apiID).APIDefinitionModel(aPIDefinitionModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.UpdateApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApi`: ApiModifyKeySuccessModelModel
    fmt.Fprintf(os.Stdout, "Response from `APIsAPI.UpdateApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIDefinitionModel** | [**APIDefinitionModel**](APIDefinitionModel.md) |  | 

### Return type

[**ApiModifyKeySuccessModelModel**](ApiModifyKeySuccessModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

