# \OrganisationQuotasApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrgKey**](OrganisationQuotasApi.md#AddOrgKey) | **Post** /tyk/org/keys/{keyID} | Create an organisation key
[**DeleteOrgKey**](OrganisationQuotasApi.md#DeleteOrgKey) | **Delete** /tyk/org/keys/{keyID} | Delete Organisation Key
[**GetOrgKey**](OrganisationQuotasApi.md#GetOrgKey) | **Get** /tyk/org/keys/{keyID} | Get an Organisation Key
[**ListOrgKeys**](OrganisationQuotasApi.md#ListOrgKeys) | **Get** /tyk/org/keys | List Organisation Keys
[**UpdateOrgKey**](OrganisationQuotasApi.md#UpdateOrgKey) | **Put** /tyk/org/keys/{keyID} | Update Organisation Key



## AddOrgKey

> ApiModifyKeySuccess AddOrgKey(ctx, keyID).SessionState(sessionState).Execute()

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
    sessionState := *openapiclient.NewSessionState() // SessionState |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationQuotasApi.AddOrgKey(context.Background(), keyID).SessionState(sessionState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationQuotasApi.AddOrgKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOrgKey`: ApiModifyKeySuccess
    fmt.Fprintf(os.Stdout, "Response from `OrganisationQuotasApi.AddOrgKey`: %v\n", resp)
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

 **sessionState** | [**SessionState**](SessionState.md) |  | 

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


## DeleteOrgKey

> ApiStatusMessage DeleteOrgKey(ctx, keyID).Execute()

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
    resp, r, err := apiClient.OrganisationQuotasApi.DeleteOrgKey(context.Background(), keyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationQuotasApi.DeleteOrgKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrgKey`: ApiStatusMessage
    fmt.Fprintf(os.Stdout, "Response from `OrganisationQuotasApi.DeleteOrgKey`: %v\n", resp)
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

[**ApiStatusMessage**](ApiStatusMessage.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgKey

> SessionState GetOrgKey(ctx, keyID).Execute()

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
    resp, r, err := apiClient.OrganisationQuotasApi.GetOrgKey(context.Background(), keyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationQuotasApi.GetOrgKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgKey`: SessionState
    fmt.Fprintf(os.Stdout, "Response from `OrganisationQuotasApi.GetOrgKey`: %v\n", resp)
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

[**SessionState**](SessionState.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgKeys

> ListOrgKeys200Response ListOrgKeys(ctx).Execute()

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
    resp, r, err := apiClient.OrganisationQuotasApi.ListOrgKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationQuotasApi.ListOrgKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgKeys`: ListOrgKeys200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganisationQuotasApi.ListOrgKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgKeysRequest struct via the builder pattern


### Return type

[**ListOrgKeys200Response**](ListOrgKeys200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgKey

> ApiModifyKeySuccess UpdateOrgKey(ctx, keyID).ResetQuota(resetQuota).SessionState(sessionState).Execute()

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
    sessionState := *openapiclient.NewSessionState() // SessionState |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationQuotasApi.UpdateOrgKey(context.Background(), keyID).ResetQuota(resetQuota).SessionState(sessionState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationQuotasApi.UpdateOrgKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrgKey`: ApiModifyKeySuccess
    fmt.Fprintf(os.Stdout, "Response from `OrganisationQuotasApi.UpdateOrgKey`: %v\n", resp)
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
 **sessionState** | [**SessionState**](SessionState.md) |  | 

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

