# \KeysAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddKey**](KeysAPI.md#AddKey) | **Post** /tyk/keys | Create a key.
[**CreateCustomKey**](KeysAPI.md#CreateCustomKey) | **Post** /tyk/keys/{keyID} | Create custom key / Import key
[**CreateKey**](KeysAPI.md#CreateKey) | **Post** /tyk/keys/create | Create a key.
[**DeleteKey**](KeysAPI.md#DeleteKey) | **Delete** /tyk/keys/{keyID} | Delete a key.
[**GetKey**](KeysAPI.md#GetKey) | **Get** /tyk/keys/{keyID} | Get a key with ID.
[**ListKeys**](KeysAPI.md#ListKeys) | **Get** /tyk/keys | List keys.
[**SetPoliciesToHashedKey**](KeysAPI.md#SetPoliciesToHashedKey) | **Post** /tyk/keys/policy/{keyID} | Set policies for a hashed key.
[**UpdateKey**](KeysAPI.md#UpdateKey) | **Put** /tyk/keys/{keyID} | Update key.
[**ValidateAKeyDefinition**](KeysAPI.md#ValidateAKeyDefinition) | **Post** /tyk/keys/preview | This will validate a key definition.



## AddKey

> ApiModifyKeySuccess AddKey(ctx).Hashed(hashed).SessionState(sessionState).Execute()

Create a key.



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
	hashed := true // bool | When set to true the key_hash returned will be similar to the un-hashed key name. (optional)
	sessionState := *openapiclient.NewSessionState() // SessionState |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.AddKey(context.Background()).Hashed(hashed).SessionState(sessionState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.AddKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddKey`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.AddKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hashed** | **bool** | When set to true the key_hash returned will be similar to the un-hashed key name. | 
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

> ApiModifyKeySuccess CreateCustomKey(ctx, keyID).SuppressReset(suppressReset).Hashed(hashed).SessionState(sessionState).Execute()

Create custom key / Import key



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
	keyID := "customKey" // string | Name to give the custom key.
	suppressReset := "1" // string | Adding the suppress_reset parameter and setting it to 1, will cause Tyk not to reset the quota limit that is in the current live quota manager. By default Tyk will reset the quota in the live quota manager (initialising it) when adding a key. Adding the `suppress_reset` flag to the URL parameters will avoid this behaviour. (optional)
	hashed := true // bool | When set to true the key_hash returned will be similar to the un-hashed key name. (optional)
	sessionState := *openapiclient.NewSessionState() // SessionState |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.CreateCustomKey(context.Background(), keyID).SuppressReset(suppressReset).Hashed(hashed).SessionState(sessionState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.CreateCustomKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomKey`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.CreateCustomKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyID** | **string** | Name to give the custom key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suppressReset** | **string** | Adding the suppress_reset parameter and setting it to 1, will cause Tyk not to reset the quota limit that is in the current live quota manager. By default Tyk will reset the quota in the live quota manager (initialising it) when adding a key. Adding the &#x60;suppress_reset&#x60; flag to the URL parameters will avoid this behaviour. | 
 **hashed** | **bool** | When set to true the key_hash returned will be similar to the un-hashed key name. | 
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


## CreateKey

> ApiModifyKeySuccess CreateKey(ctx).SessionState(sessionState).Execute()

Create a key.



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
	sessionState := *openapiclient.NewSessionState() // SessionState |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.CreateKey(context.Background()).SessionState(sessionState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.CreateKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateKey`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.CreateKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeyRequest struct via the builder pattern


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


## DeleteKey

> ApiModifyKeySuccess DeleteKey(ctx, keyID).Hashed(hashed).Execute()

Delete a key.



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
	keyID := "5e9d9544a1dcd60001d0ed20e7f75f9e03534825b7aef9df749582e5" // string | The key ID.
	hashed := false // bool | Use the hash of the key as input instead of the full key. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.DeleteKey(context.Background(), keyID).Hashed(hashed).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.DeleteKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteKey`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.DeleteKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyID** | **string** | The key ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hashed** | **bool** | Use the hash of the key as input instead of the full key. | 

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


## GetKey

> SessionState GetKey(ctx, keyID).Hashed(hashed).Execute()

Get a key with ID.



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
	keyID := "5e9d9544a1dcd60001d0ed20e7f75f9e03534825b7aef9df749582e5" // string | The key ID.
	hashed := true // bool | Use the hash of the key as input instead of the full key. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.GetKey(context.Background(), keyID).Hashed(hashed).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.GetKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKey`: SessionState
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.GetKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyID** | **string** | The key ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hashed** | **bool** | Use the hash of the key as input instead of the full key. | 

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

List keys.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.ListKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.ListKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKeys`: ApiAllKeys
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.ListKeys`: %v\n", resp)
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


## SetPoliciesToHashedKey

> ApiModifyKeySuccess SetPoliciesToHashedKey(ctx, keyID).PolicyUpdateObj(policyUpdateObj).Execute()

Set policies for a hashed key.



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
	keyID := "5e9d9544a1dcd60001d0ed207eb558517c3c48fb826c62cc6f6161eb" // string | Name to give the custom key.
	policyUpdateObj := *openapiclient.NewPolicyUpdateObj() // PolicyUpdateObj |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.SetPoliciesToHashedKey(context.Background(), keyID).PolicyUpdateObj(policyUpdateObj).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.SetPoliciesToHashedKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPoliciesToHashedKey`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.SetPoliciesToHashedKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyID** | **string** | Name to give the custom key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPoliciesToHashedKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyUpdateObj** | [**PolicyUpdateObj**](PolicyUpdateObj.md) |  | 

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


## UpdateKey

> ApiModifyKeySuccess UpdateKey(ctx, keyID).SuppressReset(suppressReset).Hashed(hashed).SessionState(sessionState).Execute()

Update key.



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
	keyID := "5e9d9544a1dcd60001d0ed20766d9a6ec6b4403b93a554feefef4708" // string | ID of the key you want to update.
	suppressReset := "1" // string | Adding the suppress_reset parameter and setting it to 1 will cause Tyk not to reset the quota limit that is in the current live quota manager. By default Tyk will reset the quota in the live quota manager (initialising it) when adding a key. Adding the `suppress_reset` flag to the URL parameters will avoid this behaviour. (optional)
	hashed := true // bool | When set to true the key_hash returned will be similar to the un-hashed key name. (optional)
	sessionState := *openapiclient.NewSessionState() // SessionState |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.UpdateKey(context.Background(), keyID).SuppressReset(suppressReset).Hashed(hashed).SessionState(sessionState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.UpdateKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateKey`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.UpdateKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyID** | **string** | ID of the key you want to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suppressReset** | **string** | Adding the suppress_reset parameter and setting it to 1 will cause Tyk not to reset the quota limit that is in the current live quota manager. By default Tyk will reset the quota in the live quota manager (initialising it) when adding a key. Adding the &#x60;suppress_reset&#x60; flag to the URL parameters will avoid this behaviour. | 
 **hashed** | **bool** | When set to true the key_hash returned will be similar to the un-hashed key name. | 
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


## ValidateAKeyDefinition

> SessionState ValidateAKeyDefinition(ctx).SessionState(sessionState).Execute()

This will validate a key definition.



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
	sessionState := *openapiclient.NewSessionState() // SessionState |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.ValidateAKeyDefinition(context.Background()).SessionState(sessionState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.ValidateAKeyDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateAKeyDefinition`: SessionState
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.ValidateAKeyDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAKeyDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionState** | [**SessionState**](SessionState.md) |  | 

### Return type

[**SessionState**](SessionState.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

