# \OAuthAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOAuthClient**](OAuthAPI.md#CreateOAuthClient) | **Post** /tyk/oauth/clients/create | Create new OAuth client
[**DeleteOAuthClient**](OAuthAPI.md#DeleteOAuthClient) | **Delete** /tyk/oauth/clients/{apiID}/{keyName} | Delete OAuth client
[**GetApisForOauthApp**](OAuthAPI.md#GetApisForOauthApp) | **Get** /tyk/oauth/clients/apis/{appID} | Get API IDs for APIS that use the specified client_id(appID) for OAuth
[**GetOAuthClient**](OAuthAPI.md#GetOAuthClient) | **Get** /tyk/oauth/clients/{apiID}/{keyName} | Get OAuth client
[**GetOAuthClientTokens**](OAuthAPI.md#GetOAuthClientTokens) | **Get** /tyk/oauth/clients/{apiID}/{keyName}/tokens | List tokens for a provided API ID and OAuth-client ID
[**InvalidateOAuthRefresh**](OAuthAPI.md#InvalidateOAuthRefresh) | **Delete** /tyk/oauth/refresh/{keyName} | Invalidate OAuth refresh token
[**ListOAuthClients**](OAuthAPI.md#ListOAuthClients) | **Get** /tyk/oauth/clients/{apiID} | List oAuth clients
[**PurgeLapsedOAuthTokens**](OAuthAPI.md#PurgeLapsedOAuthTokens) | **Delete** /tyk/oauth/tokens | Purge lapsed OAuth tokens
[**RevokeAllTokens**](OAuthAPI.md#RevokeAllTokens) | **Post** /tyk/oauth/revoke_all | Revoke all client&#39;s tokens
[**RevokeSingleToken**](OAuthAPI.md#RevokeSingleToken) | **Post** /tyk/oauth/revoke | revoke token
[**RotateOauthClient**](OAuthAPI.md#RotateOauthClient) | **Put** /tyk/oauth/clients/{apiID}/{keyName}/rotate | Rotate the oath client secret
[**UpdateOAuthClient**](OAuthAPI.md#UpdateOAuthClient) | **Put** /tyk/oauth/clients/{apiID}/{keyName} | Update OAuth metadata,redirecturi,description and Policy ID



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
	openapiclient "github.com/TykTechnologies/gateway-sdk/pkg/apim"
)

func main() {
	newClientRequest := *openapiclient.NewNewClientRequest() // NewClientRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuthAPI.CreateOAuthClient(context.Background()).NewClientRequest(newClientRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.CreateOAuthClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOAuthClient`: NewClientRequest
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.CreateOAuthClient`: %v\n", resp)
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
	openapiclient "github.com/TykTechnologies/gateway-sdk/pkg/apim"
)

func main() {
	apiID := "b84fe1a04e5648927971c0557971565c" // string | The API id
	keyName := "2a06b398c17f46908de3dffcb71ef87df" // string | The Client ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuthAPI.DeleteOAuthClient(context.Background(), apiID, keyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.DeleteOAuthClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOAuthClient`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.DeleteOAuthClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API id | 
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


## GetApisForOauthApp

> []string GetApisForOauthApp(ctx, appID).OrgID(orgID).Execute()

Get API IDs for APIS that use the specified client_id(appID) for OAuth



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TykTechnologies/gateway-sdk/pkg/apim"
)

func main() {
	appID := "2a06b398c17f46908de3dffcb71ef87df" // string | The Client ID
	orgID := "orgID_example" // string | The Org Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuthAPI.GetApisForOauthApp(context.Background(), appID).OrgID(orgID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.GetApisForOauthApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApisForOauthApp`: []string
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.GetApisForOauthApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appID** | **string** | The Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApisForOauthAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgID** | **string** | The Org Id | 

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
	openapiclient "github.com/TykTechnologies/gateway-sdk/pkg/apim"
)

func main() {
	apiID := "b84fe1a04e5648927971c0557971565c" // string | The API id
	keyName := "2a06b398c17f46908de3dffcb71ef87df" // string | The Client ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuthAPI.GetOAuthClient(context.Background(), apiID, keyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.GetOAuthClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOAuthClient`: NewClientRequest
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.GetOAuthClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API id | 
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

> GetOAuthClientTokens200Response GetOAuthClientTokens(ctx, apiID, keyName).Page(page).Execute()

List tokens for a provided API ID and OAuth-client ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TykTechnologies/gateway-sdk/pkg/apim"
)

func main() {
	apiID := "b84fe1a04e5648927971c0557971565c" // string | The API id
	keyName := "2a06b398c17f46908de3dffcb71ef87df" // string | The Client ID
	page := int32(1) // int32 | Use page query parameter to say which page number you want returned. (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuthAPI.GetOAuthClientTokens(context.Background(), apiID, keyName).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.GetOAuthClientTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOAuthClientTokens`: GetOAuthClientTokens200Response
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.GetOAuthClientTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API id | 
**keyName** | **string** | The Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuthClientTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Use page query parameter to say which page number you want returned. | [default to 1]

### Return type

[**GetOAuthClientTokens200Response**](GetOAuthClientTokens200Response.md)

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
	openapiclient "github.com/TykTechnologies/gateway-sdk/pkg/apim"
)

func main() {
	keyName := "2a06b398c17f46908de3dffcb71ef87df" // string | The Client ID
	apiId := "b84fe1a04e5648927971c0557971565c" // string | The API id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuthAPI.InvalidateOAuthRefresh(context.Background(), keyName).ApiId(apiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.InvalidateOAuthRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvalidateOAuthRefresh`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.InvalidateOAuthRefresh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyName** | **string** | The Client ID | 

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
	openapiclient "github.com/TykTechnologies/gateway-sdk/pkg/apim"
)

func main() {
	apiID := "1bd5c61b0e694082902cf15ddcc9e6a7" // string | The API ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuthAPI.ListOAuthClients(context.Background(), apiID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.ListOAuthClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOAuthClients`: []NewClientRequest
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

[**[]NewClientRequest**](NewClientRequest.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurgeLapsedOAuthTokens

> ApiStatusMessage PurgeLapsedOAuthTokens(ctx).Scope(scope).Execute()

Purge lapsed OAuth tokens



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TykTechnologies/gateway-sdk/pkg/apim"
)

func main() {
	scope := "lapsed" // string | purge lapsed tokens

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuthAPI.PurgeLapsedOAuthTokens(context.Background()).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.PurgeLapsedOAuthTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurgeLapsedOAuthTokens`: ApiStatusMessage
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.PurgeLapsedOAuthTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPurgeLapsedOAuthTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** | purge lapsed tokens | 

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


## RevokeAllTokens

> ApiStatusMessage RevokeAllTokens(ctx).ClientId(clientId).ClientSecret(clientSecret).OrgId(orgId).Execute()

Revoke all client's tokens



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TykTechnologies/gateway-sdk/pkg/apim"
)

func main() {
	clientId := "clientId_example" // string | id of oauth client
	clientSecret := "clientSecret_example" // string | OAuth client secret to ensure that its a valid operation
	orgId := "orgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuthAPI.RevokeAllTokens(context.Background()).ClientId(clientId).ClientSecret(clientSecret).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.RevokeAllTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevokeAllTokens`: ApiStatusMessage
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.RevokeAllTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeAllTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | id of oauth client | 
 **clientSecret** | **string** | OAuth client secret to ensure that its a valid operation | 
 **orgId** | **string** |  | 

### Return type

[**ApiStatusMessage**](ApiStatusMessage.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeSingleToken

> ApiStatusMessage RevokeSingleToken(ctx).ClientId(clientId).Token(token).OrgId(orgId).TokenTypeHint(tokenTypeHint).Execute()

revoke token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TykTechnologies/gateway-sdk/pkg/apim"
)

func main() {
	clientId := "clientId_example" // string | id of oauth client
	token := "token_example" // string | token to be revoked
	orgId := "orgId_example" // string |  (optional)
	tokenTypeHint := "tokenTypeHint_example" // string | type of token to be revoked, if sent then the accepted values are access_token and refresh_token. String value and optional, of not provided then it will attempt to remove access and refresh tokens that matches (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuthAPI.RevokeSingleToken(context.Background()).ClientId(clientId).Token(token).OrgId(orgId).TokenTypeHint(tokenTypeHint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.RevokeSingleToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevokeSingleToken`: ApiStatusMessage
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.RevokeSingleToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeSingleTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | id of oauth client | 
 **token** | **string** | token to be revoked | 
 **orgId** | **string** |  | 
 **tokenTypeHint** | **string** | type of token to be revoked, if sent then the accepted values are access_token and refresh_token. String value and optional, of not provided then it will attempt to remove access and refresh tokens that matches | 

### Return type

[**ApiStatusMessage**](ApiStatusMessage.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateOauthClient

> NewClientRequest RotateOauthClient(ctx, apiID, keyName).Execute()

Rotate the oath client secret



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TykTechnologies/gateway-sdk/pkg/apim"
)

func main() {
	apiID := "b84fe1a04e5648927971c0557971565c" // string | The API id
	keyName := "2a06b398c17f46908de3dffcb71ef87df" // string | The Client ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuthAPI.RotateOauthClient(context.Background(), apiID, keyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.RotateOauthClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateOauthClient`: NewClientRequest
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.RotateOauthClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API id | 
**keyName** | **string** | The Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateOauthClientRequest struct via the builder pattern


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


## UpdateOAuthClient

> NewClientRequest UpdateOAuthClient(ctx, apiID, keyName).NewClientRequest(newClientRequest).Execute()

Update OAuth metadata,redirecturi,description and Policy ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TykTechnologies/gateway-sdk/pkg/apim"
)

func main() {
	apiID := "b84fe1a04e5648927971c0557971565c" // string | The API id
	keyName := "2a06b398c17f46908de3dffcb71ef87df" // string | The Client ID
	newClientRequest := *openapiclient.NewNewClientRequest() // NewClientRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuthAPI.UpdateOAuthClient(context.Background(), apiID, keyName).NewClientRequest(newClientRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.UpdateOAuthClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOAuthClient`: NewClientRequest
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.UpdateOAuthClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API id | 
**keyName** | **string** | The Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOAuthClientRequest struct via the builder pattern


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

