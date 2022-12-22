# \PoliciesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPolicy**](PoliciesApi.md#AddPolicy) | **Post** /tyk/policies | Create a Policy
[**DeletePolicy**](PoliciesApi.md#DeletePolicy) | **Delete** /tyk/policies/{polID} | Delete a Policy
[**GetPolicy**](PoliciesApi.md#GetPolicy) | **Get** /tyk/policies/{polID} | Get a Policy
[**ListPolicies**](PoliciesApi.md#ListPolicies) | **Get** /tyk/policies | List Policies
[**UpdatePolicy**](PoliciesApi.md#UpdatePolicy) | **Put** /tyk/policies/{polID} | Update a Policy



## AddPolicy

> ApiModifyKeySuccess AddPolicy(ctx).Policy(policy).Execute()

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
    policy := *openapiclient.NewPolicy() // Policy |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesApi.AddPolicy(context.Background()).Policy(policy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.AddPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPolicy`: ApiModifyKeySuccess
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.AddPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policy** | [**Policy**](Policy.md) |  | 

### Return type

[**ApiModifyKeySuccess**](ApiModifyKeySuccess.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicy

> ApiModifyKeySuccess DeletePolicy(ctx, polID).Execute()

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
    resp, r, err := apiClient.PoliciesApi.DeletePolicy(context.Background(), polID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.DeletePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePolicy`: ApiModifyKeySuccess
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.DeletePolicy`: %v\n", resp)
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

[**ApiModifyKeySuccess**](ApiModifyKeySuccess.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicy

> Policy GetPolicy(ctx, polID).Execute()

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
    resp, r, err := apiClient.PoliciesApi.GetPolicy(context.Background(), polID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.GetPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicy`: Policy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.GetPolicy`: %v\n", resp)
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

[**Policy**](Policy.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicies

> []Policy ListPolicies(ctx).Execute()

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
    resp, r, err := apiClient.PoliciesApi.ListPolicies(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.ListPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPolicies`: []Policy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.ListPolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPoliciesRequest struct via the builder pattern


### Return type

[**[]Policy**](Policy.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicy

> ApiModifyKeySuccess UpdatePolicy(ctx, polID).Policy(policy).Execute()

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
    policy := *openapiclient.NewPolicy() // Policy |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesApi.UpdatePolicy(context.Background(), polID).Policy(policy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.UpdatePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePolicy`: ApiModifyKeySuccess
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.UpdatePolicy`: %v\n", resp)
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

 **policy** | [**Policy**](Policy.md) |  | 

### Return type

[**ApiModifyKeySuccess**](ApiModifyKeySuccess.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

