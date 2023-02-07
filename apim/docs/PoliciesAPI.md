# \PoliciesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPolicy**](PoliciesAPI.md#AddPolicy) | **Post** /tyk/policies | Create a Policy
[**DeletePolicy**](PoliciesAPI.md#DeletePolicy) | **Delete** /tyk/policies/{polID} | Delete a Policy
[**GetPolicy**](PoliciesAPI.md#GetPolicy) | **Get** /tyk/policies/{polID} | Get a Policy
[**ListPolicies**](PoliciesAPI.md#ListPolicies) | **Get** /tyk/policies | List Policies
[**UpdatePolicy**](PoliciesAPI.md#UpdatePolicy) | **Put** /tyk/policies/{polID} | Update a Policy



## AddPolicy

> ApiModifyKeySuccessModelModel AddPolicy(ctx).PolicyModel(policyModel).Execute()

Create a Policy



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
    policyModel := *openapiclient.NewPolicyModel() // PolicyModel | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesAPI.AddPolicy(context.Background()).PolicyModel(policyModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.AddPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPolicy`: ApiModifyKeySuccessModelModel
    fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.AddPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyModel** | [**PolicyModel**](PolicyModel.md) |  | 

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


## DeletePolicy

> ApiModifyKeySuccessModelModel DeletePolicy(ctx, polID).Execute()

Delete a Policy



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
    polID := "polID_example" // string | The policy ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesAPI.DeletePolicy(context.Background(), polID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.DeletePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePolicy`: ApiModifyKeySuccessModelModel
    fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.DeletePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**polID** | **string** | The policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiModifyKeySuccessModelModel**](ApiModifyKeySuccessModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicy

> PolicyModelModel GetPolicy(ctx, polID).Execute()

Get a Policy



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
    polID := "polID_example" // string | The policy ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesAPI.GetPolicy(context.Background(), polID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.GetPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicy`: PolicyModelModel
    fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.GetPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**polID** | **string** | The policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyModelModel**](PolicyModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicies

> []PolicyModelModel ListPolicies(ctx).Execute()

List Policies



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
    resp, r, err := apiClient.PoliciesAPI.ListPolicies(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.ListPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPolicies`: []PolicyModelModel
    fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.ListPolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPoliciesRequest struct via the builder pattern


### Return type

[**[]PolicyModelModel**](PolicyModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicy

> ApiModifyKeySuccessModelModel UpdatePolicy(ctx, polID).PolicyModel(policyModel).Execute()

Update a Policy



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
    polID := "polID_example" // string | The policy ID
    policyModel := *openapiclient.NewPolicyModel() // PolicyModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesAPI.UpdatePolicy(context.Background(), polID).PolicyModel(policyModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.UpdatePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePolicy`: ApiModifyKeySuccessModelModel
    fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.UpdatePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**polID** | **string** | The policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyModel** | [**PolicyModel**](PolicyModel.md) |  | 

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

