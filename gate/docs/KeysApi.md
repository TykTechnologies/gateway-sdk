# \KeysApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddKey**](KeysApi.md#AddKey) | **Post** /tyk/keys | Create a key
[**CreateCustomKey**](KeysApi.md#CreateCustomKey) | **Post** /tyk/keys/{keyID} | Create Custom Key / Import Key
[**DeleteKey**](KeysApi.md#DeleteKey) | **Delete** /tyk/keys/{keyID} | Delete Key
[**GetKey**](KeysApi.md#GetKey) | **Get** /tyk/keys/{keyID} | Get a Key
[**ListKeys**](KeysApi.md#ListKeys) | **Get** /tyk/keys | List Keys
[**UpdateKey**](KeysApi.md#UpdateKey) | **Put** /tyk/keys/{keyID} | Update Key



## AddKey

> ApiModifyKeySuccess AddKey(ctx).SessionState(sessionState).Execute()

Create a key



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
    sessionState := *openapiclient.NewSessionState() // SessionState |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeysApi.AddKey(context.Background()).SessionState(sessionState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeysApi.AddKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddKey`: ApiModifyKeySuccess
    fmt.Fprintf(os.Stdout, "Response from `KeysApi.AddKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddKeyRequest struct via the builder pattern


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


## CreateCustomKey

> ApiModifyKeySuccess CreateCustomKey(ctx, keyID).SuppressReset(suppressReset).SessionState(sessionState).Execute()

Create Custom Key / Import Key



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
    suppressReset := "suppressReset_example" // string | Adding the suppress_reset parameter and setting it to 1, will cause Tyk not to reset the quota limit that is in the current live quota manager. By default Tyk will reset the quota in the live quota manager (initialising it) when adding a key. Adding the `suppress_reset` flag to the URL parameters will avoid this behaviour. (optional)
    sessionState := *openapiclient.NewSessionState() // SessionState |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeysApi.CreateCustomKey(context.Background(), keyID).SuppressReset(suppressReset).SessionState(sessionState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeysApi.CreateCustomKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomKey`: ApiModifyKeySuccess
    fmt.Fprintf(os.Stdout, "Response from `KeysApi.CreateCustomKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyID** | **string** | The Key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suppressReset** | **string** | Adding the suppress_reset parameter and setting it to 1, will cause Tyk not to reset the quota limit that is in the current live quota manager. By default Tyk will reset the quota in the live quota manager (initialising it) when adding a key. Adding the &#x60;suppress_reset&#x60; flag to the URL parameters will avoid this behaviour. | 
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


## DeleteKey

> ApiStatusMessage DeleteKey(ctx, keyID).Execute()

Delete Key



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
    resp, r, err := apiClient.KeysApi.DeleteKey(context.Background(), keyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeysApi.DeleteKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteKey`: ApiStatusMessage
    fmt.Fprintf(os.Stdout, "Response from `KeysApi.DeleteKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyID** | **string** | The Key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyRequest struct via the builder pattern


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


## GetKey

> SessionState GetKey(ctx, keyID).Execute()

Get a Key



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
    resp, r, err := apiClient.KeysApi.GetKey(context.Background(), keyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeysApi.GetKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKey`: SessionState
    fmt.Fprintf(os.Stdout, "Response from `KeysApi.GetKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyID** | **string** | The Key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyRequest struct via the builder pattern


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


## ListKeys

> ApiAllKeys ListKeys(ctx).Execute()

List Keys



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
    resp, r, err := apiClient.KeysApi.ListKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeysApi.ListKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKeys`: ApiAllKeys
    fmt.Fprintf(os.Stdout, "Response from `KeysApi.ListKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListKeysRequest struct via the builder pattern


### Return type

[**ApiAllKeys**](ApiAllKeys.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKey

> ApiModifyKeySuccess UpdateKey(ctx, keyID).SuppressReset(suppressReset).SessionState(sessionState).Execute()

Update Key



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
    suppressReset := "suppressReset_example" // string | Adding the suppress_reset parameter and setting it to 1, will cause Tyk not to reset the quota limit that is in the current live quota manager. By default Tyk will reset the quota in the live quota manager (initialising it) when adding a key. Adding the `suppress_reset` flag to the URL parameters will avoid this behaviour. (optional)
    sessionState := *openapiclient.NewSessionState() // SessionState |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeysApi.UpdateKey(context.Background(), keyID).SuppressReset(suppressReset).SessionState(sessionState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeysApi.UpdateKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKey`: ApiModifyKeySuccess
    fmt.Fprintf(os.Stdout, "Response from `KeysApi.UpdateKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyID** | **string** | The Key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suppressReset** | **string** | Adding the suppress_reset parameter and setting it to 1, will cause Tyk not to reset the quota limit that is in the current live quota manager. By default Tyk will reset the quota in the live quota manager (initialising it) when adding a key. Adding the &#x60;suppress_reset&#x60; flag to the URL parameters will avoid this behaviour. | 
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

