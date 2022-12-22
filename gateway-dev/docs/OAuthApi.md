# \OAuthApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizeClient**](OAuthApi.md#AuthorizeClient) | **Post** /tyk/oauth/authorize-client/ | Authorize client
[**CreateOAuthClient**](OAuthApi.md#CreateOAuthClient) | **Post** /tyk/oauth/clients/create | Create new OAuth client
[**DeleteOAuthClient**](OAuthApi.md#DeleteOAuthClient) | **Delete** /tyk/oauth/clients/{apiID}/{keyName} | Delete OAuth client
[**GetOAuthClient**](OAuthApi.md#GetOAuthClient) | **Get** /tyk/oauth/clients/{apiID}/{keyName} | Get OAuth client
[**GetOAuthClientTokens**](OAuthApi.md#GetOAuthClientTokens) | **Get** /tyk/oauth/clients/{apiID}/{keyName}/tokens | List tokens
[**InvalidateOAuthRefresh**](OAuthApi.md#InvalidateOAuthRefresh) | **Delete** /tyk/oauth/refresh/{keyName} | Invalidate OAuth refresh token
[**ListOAuthClients**](OAuthApi.md#ListOAuthClients) | **Get** /tyk/oauth/clients/{apiID} | List oAuth clients
[**RevokeAllTokens**](OAuthApi.md#RevokeAllTokens) | **Post** /tyk/oauth/revoke_all | revoke all client&#39;s tokens
[**RevokeSingleToken**](OAuthApi.md#RevokeSingleToken) | **Post** /tyk/oauth/revoke | revoke token
[**UpdateoAuthClient**](OAuthApi.md#UpdateoAuthClient) | **Put** /tyk/oauth/clients/{apiID} | Update OAuth metadata and Policy ID



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
    resp, r, err := apiClient.OAuthApi.AuthorizeClient(context.Background()).ResponseType(responseType).ClientId(clientId).RedirectUri(redirectUri).KeyRules(keyRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthApi.AuthorizeClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthorizeClient`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OAuthApi.AuthorizeClient`: %v\n", resp)
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

> NewClientRequest CreateOAuthClient(ctx).NewClientRequest(newClientRequest).Execute()

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
    newClientRequest := *openapiclient.NewNewClientRequest() // NewClientRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuthApi.CreateOAuthClient(context.Background()).NewClientRequest(newClientRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthApi.CreateOAuthClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOAuthClient`: NewClientRequest
    fmt.Fprintf(os.Stdout, "Response from `OAuthApi.CreateOAuthClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOAuthClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newClientRequest** | [**NewClientRequest**](NewClientRequest.md) |  | 

### Return type

[**NewClientRequest**](NewClientRequest.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOAuthClient

> ApiModifyKeySuccess DeleteOAuthClient(ctx, apiID, keyName).Execute()

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
    resp, r, err := apiClient.OAuthApi.DeleteOAuthClient(context.Background(), apiID, keyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthApi.DeleteOAuthClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOAuthClient`: ApiModifyKeySuccess
    fmt.Fprintf(os.Stdout, "Response from `OAuthApi.DeleteOAuthClient`: %v\n", resp)
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

[**ApiModifyKeySuccess**](ApiModifyKeySuccess.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOAuthClient

> NewClientRequest GetOAuthClient(ctx, apiID, keyName).Execute()

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
    resp, r, err := apiClient.OAuthApi.GetOAuthClient(context.Background(), apiID, keyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthApi.GetOAuthClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOAuthClient`: NewClientRequest
    fmt.Fprintf(os.Stdout, "Response from `OAuthApi.GetOAuthClient`: %v\n", resp)
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

[**NewClientRequest**](NewClientRequest.md)

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
    resp, r, err := apiClient.OAuthApi.GetOAuthClientTokens(context.Background(), apiID, keyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthApi.GetOAuthClientTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOAuthClientTokens`: []string
    fmt.Fprintf(os.Stdout, "Response from `OAuthApi.GetOAuthClientTokens`: %v\n", resp)
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

> ApiModifyKeySuccess InvalidateOAuthRefresh(ctx, keyName).ApiId(apiId).Execute()

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
    resp, r, err := apiClient.OAuthApi.InvalidateOAuthRefresh(context.Background(), keyName).ApiId(apiId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthApi.InvalidateOAuthRefresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvalidateOAuthRefresh`: ApiModifyKeySuccess
    fmt.Fprintf(os.Stdout, "Response from `OAuthApi.InvalidateOAuthRefresh`: %v\n", resp)
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

[**ApiModifyKeySuccess**](ApiModifyKeySuccess.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOAuthClients

> []NewClientRequest ListOAuthClients(ctx, apiID).Execute()

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
    resp, r, err := apiClient.OAuthApi.ListOAuthClients(context.Background(), apiID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthApi.ListOAuthClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOAuthClients`: []NewClientRequest
    fmt.Fprintf(os.Stdout, "Response from `OAuthApi.ListOAuthClients`: %v\n", resp)
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

[**[]NewClientRequest**](NewClientRequest.md)

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
    resp, r, err := apiClient.OAuthApi.RevokeAllTokens(context.Background()).ClientId(clientId).ClientSecret(clientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthApi.RevokeAllTokens``: %v\n", err)
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
    resp, r, err := apiClient.OAuthApi.RevokeSingleToken(context.Background()).Token(token).ClientId(clientId).TokenTypeHint(tokenTypeHint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthApi.RevokeSingleToken``: %v\n", err)
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

> []NewClientRequest UpdateoAuthClient(ctx, apiID).Execute()

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
    resp, r, err := apiClient.OAuthApi.UpdateoAuthClient(context.Background(), apiID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthApi.UpdateoAuthClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateoAuthClient`: []NewClientRequest
    fmt.Fprintf(os.Stdout, "Response from `OAuthApi.UpdateoAuthClient`: %v\n", resp)
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

[**[]NewClientRequest**](NewClientRequest.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

