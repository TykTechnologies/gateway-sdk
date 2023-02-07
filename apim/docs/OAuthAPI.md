# \OAuthAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizeClient**](OAuthAPI.md#AuthorizeClient) | **Post** /tyk/oauth/authorize-client/ | Authorize client
[**CreateOAuthClient**](OAuthAPI.md#CreateOAuthClient) | **Post** /tyk/oauth/clients/create | Create new OAuth client
[**DeleteOAuthClient**](OAuthAPI.md#DeleteOAuthClient) | **Delete** /tyk/oauth/clients/{apiID}/{keyName} | Delete OAuth client
[**GetOAuthClient**](OAuthAPI.md#GetOAuthClient) | **Get** /tyk/oauth/clients/{apiID}/{keyName} | Get OAuth client
[**GetOAuthClientTokens**](OAuthAPI.md#GetOAuthClientTokens) | **Get** /tyk/oauth/clients/{apiID}/{keyName}/tokens | List tokens
[**InvalidateOAuthRefresh**](OAuthAPI.md#InvalidateOAuthRefresh) | **Delete** /tyk/oauth/refresh/{keyName} | Invalidate OAuth refresh token
[**ListOAuthClients**](OAuthAPI.md#ListOAuthClients) | **Get** /tyk/oauth/clients/{apiID} | List oAuth clients
[**RevokeAllTokens**](OAuthAPI.md#RevokeAllTokens) | **Post** /tyk/oauth/revoke_all | revoke all client&#39;s tokens
[**RevokeSingleToken**](OAuthAPI.md#RevokeSingleToken) | **Post** /tyk/oauth/revoke | revoke token
[**UpdateoAuthClient**](OAuthAPI.md#UpdateoAuthClient) | **Put** /tyk/oauth/clients/{apiID} | Update OAuth metadata and Policy ID



## AuthorizeClient

> map[string]interface{} AuthorizeClient(ctx).ResponseType(responseType).ClientId(clientId).RedirectUri(redirectUri).KeyRules(keyRules).Execute()

Authorize client



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
    responseType := "responseType_example" // string | Should be provided by requesting client as part of authorisation request, this should be either `code` or `token` depending on the methods you have specified for the API. (optional)
    clientId := "clientId_example" // string | Should be provided by requesting client as part of authorisation request. The Client ID that is making the request. (optional)
    redirectUri := "redirectUri_example" // string | Should be provided by requesting client as part of authorisation request. Must match with the record stored with Tyk. (optional)
    keyRules := "keyRules_example" // string | A string representation of a Session Object (form-encoded). This should be provided by your application in order to apply any quotas or rules to the key. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuthAPI.AuthorizeClient(context.Background()).ResponseType(responseType).ClientId(clientId).RedirectUri(redirectUri).KeyRules(keyRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.AuthorizeClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthorizeClient`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.AuthorizeClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **responseType** | **string** | Should be provided by requesting client as part of authorisation request, this should be either &#x60;code&#x60; or &#x60;token&#x60; depending on the methods you have specified for the API. | 
 **clientId** | **string** | Should be provided by requesting client as part of authorisation request. The Client ID that is making the request. | 
 **redirectUri** | **string** | Should be provided by requesting client as part of authorisation request. Must match with the record stored with Tyk. | 
 **keyRules** | **string** | A string representation of a Session Object (form-encoded). This should be provided by your application in order to apply any quotas or rules to the key. | 

### Return type

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOAuthClient

> NewClientRequestModelModel CreateOAuthClient(ctx).NewClientRequestModel(newClientRequestModel).Execute()

Create new OAuth client



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
    newClientRequestModel := *openapiclient.NewNewClientRequestModel() // NewClientRequestModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuthAPI.CreateOAuthClient(context.Background()).NewClientRequestModel(newClientRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.CreateOAuthClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOAuthClient`: NewClientRequestModelModel
    fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.CreateOAuthClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOAuthClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newClientRequestModel** | [**NewClientRequestModel**](NewClientRequestModel.md) |  | 

### Return type

[**NewClientRequestModelModel**](NewClientRequestModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOAuthClient

> ApiModifyKeySuccessModelModel DeleteOAuthClient(ctx, apiID, keyName).Execute()

Delete OAuth client



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
    keyName := "keyName_example" // string | The Client ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuthAPI.DeleteOAuthClient(context.Background(), apiID, keyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.DeleteOAuthClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOAuthClient`: ApiModifyKeySuccessModelModel
    fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.DeleteOAuthClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID | 
**keyName** | **string** | The Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOAuthClientRequest struct via the builder pattern


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


## GetOAuthClient

> NewClientRequestModelModel GetOAuthClient(ctx, apiID, keyName).Execute()

Get OAuth client

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
    keyName := "keyName_example" // string | The Client ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuthAPI.GetOAuthClient(context.Background(), apiID, keyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.GetOAuthClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOAuthClient`: NewClientRequestModelModel
    fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.GetOAuthClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID | 
**keyName** | **string** | The Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuthClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NewClientRequestModelModel**](NewClientRequestModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOAuthClientTokens

> []string GetOAuthClientTokens(ctx, apiID, keyName).Execute()

List tokens



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
    keyName := "keyName_example" // string | The Client ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuthAPI.GetOAuthClientTokens(context.Background(), apiID, keyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.GetOAuthClientTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOAuthClientTokens`: []string
    fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.GetOAuthClientTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID | 
**keyName** | **string** | The Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuthClientTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**[]string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvalidateOAuthRefresh

> ApiModifyKeySuccessModelModel InvalidateOAuthRefresh(ctx, keyName).ApiId(apiId).Execute()

Invalidate OAuth refresh token



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
    apiId := "apiId_example" // string | The API id
    keyName := "keyName_example" // string | Refresh token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuthAPI.InvalidateOAuthRefresh(context.Background(), keyName).ApiId(apiId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.InvalidateOAuthRefresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvalidateOAuthRefresh`: ApiModifyKeySuccessModelModel
    fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.InvalidateOAuthRefresh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyName** | **string** | Refresh token | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvalidateOAuthRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiId** | **string** | The API id | 


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


## ListOAuthClients

> []NewClientRequestModelModel ListOAuthClients(ctx, apiID).Execute()

List oAuth clients



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
    resp, r, err := apiClient.OAuthAPI.ListOAuthClients(context.Background(), apiID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.ListOAuthClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOAuthClients`: []NewClientRequestModelModel
    fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.ListOAuthClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOAuthClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]NewClientRequestModelModel**](NewClientRequestModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeAllTokens

> RevokeAllTokens(ctx).ClientId(clientId).ClientSecret(clientSecret).Execute()

revoke all client's tokens



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
    clientId := "clientId_example" // string | id of oauth client (optional)
    clientSecret := "clientSecret_example" // string | OAuth client secret to ensure that its a valid operation (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuthAPI.RevokeAllTokens(context.Background()).ClientId(clientId).ClientSecret(clientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.RevokeAllTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeAllTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | id of oauth client | 
 **clientSecret** | **string** | OAuth client secret to ensure that its a valid operation | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeSingleToken

> RevokeSingleToken(ctx).Token(token).ClientId(clientId).TokenTypeHint(tokenTypeHint).Execute()

revoke token



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
    token := "token_example" // string | token to be revoked (optional)
    clientId := "clientId_example" // string | id of oauth client (optional)
    tokenTypeHint := "tokenTypeHint_example" // string | type of token to be revoked, if sent then the accepted values are access_token and refresh_token. String value and optional, of not provided then it will attempt to remove access and refresh tokens that matchs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuthAPI.RevokeSingleToken(context.Background()).Token(token).ClientId(clientId).TokenTypeHint(tokenTypeHint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.RevokeSingleToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeSingleTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | token to be revoked | 
 **clientId** | **string** | id of oauth client | 
 **tokenTypeHint** | **string** | type of token to be revoked, if sent then the accepted values are access_token and refresh_token. String value and optional, of not provided then it will attempt to remove access and refresh tokens that matchs | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateoAuthClient

> []NewClientRequestModelModel UpdateoAuthClient(ctx, apiID).Execute()

Update OAuth metadata and Policy ID



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
    resp, r, err := apiClient.OAuthAPI.UpdateoAuthClient(context.Background(), apiID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.UpdateoAuthClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateoAuthClient`: []NewClientRequestModelModel
    fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.UpdateoAuthClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateoAuthClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]NewClientRequestModelModel**](NewClientRequestModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

