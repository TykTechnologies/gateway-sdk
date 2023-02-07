# \OrganisationQuotasAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrgKey**](OrganisationQuotasAPI.md#AddOrgKey) | **Post** /tyk/org/keys/{keyID} | Create an organisation key
[**DeleteOrgKey**](OrganisationQuotasAPI.md#DeleteOrgKey) | **Delete** /tyk/org/keys/{keyID} | Delete Organisation Key
[**GetOrgKey**](OrganisationQuotasAPI.md#GetOrgKey) | **Get** /tyk/org/keys/{keyID} | Get an Organisation Key
[**ListOrgKeys**](OrganisationQuotasAPI.md#ListOrgKeys) | **Get** /tyk/org/keys | List Organisation Keys
[**UpdateOrgKey**](OrganisationQuotasAPI.md#UpdateOrgKey) | **Put** /tyk/org/keys/{keyID} | Update Organisation Key



## AddOrgKey

> ApiModifyKeySuccessModelModel AddOrgKey(ctx, keyID).SessionStateModel(sessionStateModel).Execute()

Create an organisation key



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
    keyID := "keyID_example" // string | The Key ID
    sessionStateModel := *openapiclient.NewSessionStateModel() // SessionStateModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationQuotasAPI.AddOrgKey(context.Background(), keyID).SessionStateModel(sessionStateModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationQuotasAPI.AddOrgKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOrgKey`: ApiModifyKeySuccessModelModel
    fmt.Fprintf(os.Stdout, "Response from `OrganisationQuotasAPI.AddOrgKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyID** | **string** | The Key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOrgKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sessionStateModel** | [**SessionStateModel**](SessionStateModel.md) |  | 

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


## DeleteOrgKey

> ApiStatusMessageModelModel DeleteOrgKey(ctx, keyID).Execute()

Delete Organisation Key



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
    keyID := "keyID_example" // string | The Key ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationQuotasAPI.DeleteOrgKey(context.Background(), keyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationQuotasAPI.DeleteOrgKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrgKey`: ApiStatusMessageModelModel
    fmt.Fprintf(os.Stdout, "Response from `OrganisationQuotasAPI.DeleteOrgKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyID** | **string** | The Key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgKeyRequest struct via the builder pattern


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


## GetOrgKey

> SessionStateModelModel GetOrgKey(ctx, keyID).Execute()

Get an Organisation Key



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
    keyID := "keyID_example" // string | The Key ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationQuotasAPI.GetOrgKey(context.Background(), keyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationQuotasAPI.GetOrgKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgKey`: SessionStateModelModel
    fmt.Fprintf(os.Stdout, "Response from `OrganisationQuotasAPI.GetOrgKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyID** | **string** | The Key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SessionStateModelModel**](SessionStateModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgKeys

> ListOrgKeys200ResponseModelModel ListOrgKeys(ctx).Execute()

List Organisation Keys



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
    resp, r, err := apiClient.OrganisationQuotasAPI.ListOrgKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationQuotasAPI.ListOrgKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgKeys`: ListOrgKeys200ResponseModelModel
    fmt.Fprintf(os.Stdout, "Response from `OrganisationQuotasAPI.ListOrgKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgKeysRequest struct via the builder pattern


### Return type

[**ListOrgKeys200ResponseModelModel**](ListOrgKeys200ResponseModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgKey

> ApiModifyKeySuccessModelModel UpdateOrgKey(ctx, keyID).ResetQuota(resetQuota).SessionStateModel(sessionStateModel).Execute()

Update Organisation Key



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
    keyID := "keyID_example" // string | The Key ID
    resetQuota := "resetQuota_example" // string | Adding the `reset_quota` parameter and setting it to 1, will cause Tyk reset the organisations quota in the live quota manager, it is recommended to use this mechanism to reset organisation-level access if a monthly subscription is in place. (optional)
    sessionStateModel := *openapiclient.NewSessionStateModel() // SessionStateModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationQuotasAPI.UpdateOrgKey(context.Background(), keyID).ResetQuota(resetQuota).SessionStateModel(sessionStateModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationQuotasAPI.UpdateOrgKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrgKey`: ApiModifyKeySuccessModelModel
    fmt.Fprintf(os.Stdout, "Response from `OrganisationQuotasAPI.UpdateOrgKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyID** | **string** | The Key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resetQuota** | **string** | Adding the &#x60;reset_quota&#x60; parameter and setting it to 1, will cause Tyk reset the organisations quota in the live quota manager, it is recommended to use this mechanism to reset organisation-level access if a monthly subscription is in place. | 
 **sessionStateModel** | [**SessionStateModel**](SessionStateModel.md) |  | 

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

