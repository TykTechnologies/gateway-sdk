# \APIsAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApi**](APIsAPI.md#CreateApi) | **Post** /tyk/apis | Creat an API
[**DeleteApi**](APIsAPI.md#DeleteApi) | **Delete** /tyk/apis/{apiID} | Deleting an API definition with ID.
[**GetApi**](APIsAPI.md#GetApi) | **Get** /tyk/apis/{apiID} | Get API definition with it&#39;s ID.
[**ListApiVersions**](APIsAPI.md#ListApiVersions) | **Get** /tyk/apis/{apiID}/versions | Listing versions of an API.
[**ListApis**](APIsAPI.md#ListApis) | **Get** /tyk/apis | Get list of apis
[**UpdateApi**](APIsAPI.md#UpdateApi) | **Put** /tyk/apis/{apiID} | Updating an API definition with its ID.



## CreateApi

> ApiModifyKeySuccess CreateApi(ctx).BaseApiId(baseApiId).BaseApiVersionName(baseApiVersionName).NewVersionName(newVersionName).SetDefault(setDefault).APIDefinition(aPIDefinition).Execute()

Creat an API



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
	baseApiId := "663a4ed9b6be920001b191ae" // string | The base API which the new version will be linked to. (optional)
	baseApiVersionName := "Default" // string | The version name of the base API while creating the first version. This doesn't have to be sent for the next versions but if it is set, it will override base API version name. (optional)
	newVersionName := "v2" // string | The version name of the created version. (optional)
	setDefault := true // bool | If true, the new version is set as default version. (optional)
	aPIDefinition := *openapiclient.NewAPIDefinition() // APIDefinition |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.CreateApi(context.Background()).BaseApiId(baseApiId).BaseApiVersionName(baseApiVersionName).NewVersionName(newVersionName).SetDefault(setDefault).APIDefinition(aPIDefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.CreateApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApi`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.CreateApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseApiId** | **string** | The base API which the new version will be linked to. | 
 **baseApiVersionName** | **string** | The version name of the base API while creating the first version. This doesn&#39;t have to be sent for the next versions but if it is set, it will override base API version name. | 
 **newVersionName** | **string** | The version name of the created version. | 
 **setDefault** | **bool** | If true, the new version is set as default version. | 
 **aPIDefinition** | [**APIDefinition**](APIDefinition.md) |  | 

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


## DeleteApi

> ApiModifyKeySuccess DeleteApi(ctx, apiID).Execute()

Deleting an API definition with ID.



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
	apiID := "1bd5c61b0e694082902cf15ddcc9e6a7" // string | The API ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.DeleteApi(context.Background(), apiID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.DeleteApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApi`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.DeleteApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiRequest struct via the builder pattern


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


## GetApi

> APIDefinition GetApi(ctx, apiID).Execute()

Get API definition with it's ID.



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
	apiID := "keyless" // string | The API ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.GetApi(context.Background(), apiID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.GetApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApi`: APIDefinition
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.GetApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**APIDefinition**](APIDefinition.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiVersions

> VersionMetas ListApiVersions(ctx, apiID).SearchText(searchText).AccessType(accessType).Execute()

Listing versions of an API.



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
	apiID := "keyless" // string | The API ID.
	searchText := "Sample oas" // string | Search for API version name (optional)
	accessType := "internal" // string | Filter for internal or external API versions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.ListApiVersions(context.Background(), apiID).SearchText(searchText).AccessType(accessType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.ListApiVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApiVersions`: VersionMetas
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.ListApiVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApiVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchText** | **string** | Search for API version name | 
 **accessType** | **string** | Filter for internal or external API versions | 

### Return type

[**VersionMetas**](VersionMetas.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApis

> []APIDefinition ListApis(ctx).Execute()

Get list of apis



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
	resp, r, err := apiClient.APIsAPI.ListApis(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.ListApis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApis`: []APIDefinition
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.ListApis`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApisRequest struct via the builder pattern


### Return type

[**[]APIDefinition**](APIDefinition.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApi

> ApiModifyKeySuccess UpdateApi(ctx, apiID).APIDefinition(aPIDefinition).Execute()

Updating an API definition with its ID.



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
	apiID := "1bd5c61b0e694082902cf15ddcc9e6a7" // string | The API ID.
	aPIDefinition := *openapiclient.NewAPIDefinition() // APIDefinition |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.UpdateApi(context.Background(), apiID).APIDefinition(aPIDefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.UpdateApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApi`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.UpdateApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIDefinition** | [**APIDefinition**](APIDefinition.md) |  | 

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

