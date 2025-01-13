# \TykOASAPIsAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiOAS**](TykOASAPIsAPI.md#CreateApiOAS) | **Post** /tyk/apis/oas | Create an API with Tyk OAS format.
[**DeleteOASApi**](TykOASAPIsAPI.md#DeleteOASApi) | **Delete** /tyk/apis/oas/{apiID} | Deleting a Tyk OAS API.
[**DownloadApiOASPublic**](TykOASAPIsAPI.md#DownloadApiOASPublic) | **Get** /tyk/apis/oas/{apiID}/export | Download a Tyk OAS format API.
[**DownloadApisOASPublic**](TykOASAPIsAPI.md#DownloadApisOASPublic) | **Get** /tyk/apis/oas/export | Download all Tyk OAS format APIs.
[**GetOASApi**](TykOASAPIsAPI.md#GetOASApi) | **Get** /tyk/apis/oas/{apiID} | Get a Tyk OAS API definition.
[**ImportOAS**](TykOASAPIsAPI.md#ImportOAS) | **Post** /tyk/apis/oas/import | Import an API in Tyk OAS format.
[**ListApisOAS**](TykOASAPIsAPI.md#ListApisOAS) | **Get** /tyk/apis/oas | List all APIs in Tyk OAS API format.
[**ListOASApiVersions**](TykOASAPIsAPI.md#ListOASApiVersions) | **Get** /tyk/apis/oas/{apiID}/versions | Listing versions of a Tyk OAS API.
[**PatchApiOAS**](TykOASAPIsAPI.md#PatchApiOAS) | **Patch** /tyk/apis/oas/{apiID} | Patch API in Tyk OAS format.
[**UpdateApiOAS**](TykOASAPIsAPI.md#UpdateApiOAS) | **Put** /tyk/apis/oas/{apiID} | Update a Tyk OAS API definition.



## CreateApiOAS

> ApiModifyKeySuccess CreateApiOAS(ctx).BaseApiId(baseApiId).BaseApiVersionName(baseApiVersionName).NewVersionName(newVersionName).SetDefault(setDefault).ListApisOAS200ResponseInner(listApisOAS200ResponseInner).Execute()

Create an API with Tyk OAS format.



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
	listApisOAS200ResponseInner := *openapiclient.NewListApisOAS200ResponseInner("Openapi_example", *openapiclient.NewInfo(), map[string]interface{}(123)) // ListApisOAS200ResponseInner |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TykOASAPIsAPI.CreateApiOAS(context.Background()).BaseApiId(baseApiId).BaseApiVersionName(baseApiVersionName).NewVersionName(newVersionName).SetDefault(setDefault).ListApisOAS200ResponseInner(listApisOAS200ResponseInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TykOASAPIsAPI.CreateApiOAS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApiOAS`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `TykOASAPIsAPI.CreateApiOAS`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseApiId** | **string** | The base API which the new version will be linked to. | 
 **baseApiVersionName** | **string** | The version name of the base API while creating the first version. This doesn&#39;t have to be sent for the next versions but if it is set, it will override base API version name. | 
 **newVersionName** | **string** | The version name of the created version. | 
 **setDefault** | **bool** | If true, the new version is set as default version. | 
 **listApisOAS200ResponseInner** | [**ListApisOAS200ResponseInner**](ListApisOAS200ResponseInner.md) |  | 

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


## DeleteOASApi

> ApiModifyKeySuccess DeleteOASApi(ctx, apiID).Execute()

Deleting a Tyk OAS API.



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
	resp, r, err := apiClient.TykOASAPIsAPI.DeleteOASApi(context.Background(), apiID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TykOASAPIsAPI.DeleteOASApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOASApi`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `TykOASAPIsAPI.DeleteOASApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOASApiRequest struct via the builder pattern


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


## DownloadApiOASPublic

> *os.File DownloadApiOASPublic(ctx, apiID).Mode(mode).Execute()

Download a Tyk OAS format API.



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
	apiID := "4c1c0d8fc885401053ddac4e39ef676b" // string | ID of the API you want to fetch.
	mode := "public" // string | By default mode is empty which means it will return the Tyk API OAS spec including the x-tyk-api-gateway part.   When mode=public, the Tyk OAS API spec will exclude the x-tyk-api-gateway part in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TykOASAPIsAPI.DownloadApiOASPublic(context.Background(), apiID).Mode(mode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TykOASAPIsAPI.DownloadApiOASPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadApiOASPublic`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TykOASAPIsAPI.DownloadApiOASPublic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | ID of the API you want to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadApiOASPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mode** | **string** | By default mode is empty which means it will return the Tyk API OAS spec including the x-tyk-api-gateway part.   When mode&#x3D;public, the Tyk OAS API spec will exclude the x-tyk-api-gateway part in the response. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadApisOASPublic

> *os.File DownloadApisOASPublic(ctx).Mode(mode).Execute()

Download all Tyk OAS format APIs.



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
	mode := "public" // string | By default mode is empty which means it will return the Tyk API OAS spec including the x-tyk-api-gateway part.   When mode=public, the Tyk OAS API spec will exclude the x-tyk-api-gateway part in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TykOASAPIsAPI.DownloadApisOASPublic(context.Background()).Mode(mode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TykOASAPIsAPI.DownloadApisOASPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadApisOASPublic`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TykOASAPIsAPI.DownloadApisOASPublic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadApisOASPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mode** | **string** | By default mode is empty which means it will return the Tyk API OAS spec including the x-tyk-api-gateway part.   When mode&#x3D;public, the Tyk OAS API spec will exclude the x-tyk-api-gateway part in the response. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOASApi

> ListApisOAS200ResponseInner GetOASApi(ctx, apiID).Mode(mode).Execute()

Get a Tyk OAS API definition.



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
	apiID := "4c1c0d8fc885401053ddac4e39ef676b" // string | ID of the API you want to fetch
	mode := "public" // string | By default mode is empty which means it will return the Tyk API OAS spec including the x-tyk-api-gateway part.   When mode=public, the Tyk OAS API spec will exclude the x-tyk-api-gateway part in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TykOASAPIsAPI.GetOASApi(context.Background(), apiID).Mode(mode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TykOASAPIsAPI.GetOASApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOASApi`: ListApisOAS200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `TykOASAPIsAPI.GetOASApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | ID of the API you want to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOASApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mode** | **string** | By default mode is empty which means it will return the Tyk API OAS spec including the x-tyk-api-gateway part.   When mode&#x3D;public, the Tyk OAS API spec will exclude the x-tyk-api-gateway part in the response. | 

### Return type

[**ListApisOAS200ResponseInner**](ListApisOAS200ResponseInner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportOAS

> ApiModifyKeySuccess ImportOAS(ctx).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).AllowList(allowList).ValidateRequest(validateRequest).MockResponse(mockResponse).Authentication(authentication).BaseApiId(baseApiId).BaseApiVersionName(baseApiVersionName).NewVersionName(newVersionName).SetDefault(setDefault).Model30(model30).Execute()

Import an API in Tyk OAS format.



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
	upstreamURL := "https://localhost:8080" // string | Upstream URL for the API (optional)
	listenPath := "/user-test/" // string | Listen path for the API (optional)
	customDomain := "tyk.io" // string | Custom domain for the API (optional)
	allowList := true // bool | Enable allowList middleware for all endpoints (optional)
	validateRequest := true // bool | Enable validateRequest middleware for all endpoints having a request body with media type application/json (optional)
	mockResponse := true // bool | Enable mockResponse middleware for all endpoints having responses configured. (optional)
	authentication := true // bool | Enable/disable the authentication mechanism in your Tyk Gateway for your OAS API (optional)
	baseApiId := "663a4ed9b6be920001b191ae" // string | The base API which the new version will be linked to. (optional)
	baseApiVersionName := "Default" // string | The version name of the base API while creating the first version. This doesn't have to be sent for the next versions but if it is set, it will override base API version name. (optional)
	newVersionName := "v2" // string | The version name of the created version. (optional)
	setDefault := true // bool | If true, the new version is set as default version. (optional)
	model30 := *openapiclient.NewModel30("Openapi_example", *openapiclient.NewInfo1("Title_example", "Version_example"), map[string]interface{}(123)) // Model30 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TykOASAPIsAPI.ImportOAS(context.Background()).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).AllowList(allowList).ValidateRequest(validateRequest).MockResponse(mockResponse).Authentication(authentication).BaseApiId(baseApiId).BaseApiVersionName(baseApiVersionName).NewVersionName(newVersionName).SetDefault(setDefault).Model30(model30).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TykOASAPIsAPI.ImportOAS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportOAS`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `TykOASAPIsAPI.ImportOAS`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upstreamURL** | **string** | Upstream URL for the API | 
 **listenPath** | **string** | Listen path for the API | 
 **customDomain** | **string** | Custom domain for the API | 
 **allowList** | **bool** | Enable allowList middleware for all endpoints | 
 **validateRequest** | **bool** | Enable validateRequest middleware for all endpoints having a request body with media type application/json | 
 **mockResponse** | **bool** | Enable mockResponse middleware for all endpoints having responses configured. | 
 **authentication** | **bool** | Enable/disable the authentication mechanism in your Tyk Gateway for your OAS API | 
 **baseApiId** | **string** | The base API which the new version will be linked to. | 
 **baseApiVersionName** | **string** | The version name of the base API while creating the first version. This doesn&#39;t have to be sent for the next versions but if it is set, it will override base API version name. | 
 **newVersionName** | **string** | The version name of the created version. | 
 **setDefault** | **bool** | If true, the new version is set as default version. | 
 **model30** | [**Model30**](Model30.md) |  | 

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


## ListApisOAS

> []ListApisOAS200ResponseInner ListApisOAS(ctx).Mode(mode).Execute()

List all APIs in Tyk OAS API format.



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
	mode := "public" // string | By default mode is empty which means it will return the Tyk API OAS spec including the x-tyk-api-gateway part.   When mode=public, the Tyk OAS API spec will exclude the x-tyk-api-gateway part in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TykOASAPIsAPI.ListApisOAS(context.Background()).Mode(mode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TykOASAPIsAPI.ListApisOAS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApisOAS`: []ListApisOAS200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `TykOASAPIsAPI.ListApisOAS`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApisOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mode** | **string** | By default mode is empty which means it will return the Tyk API OAS spec including the x-tyk-api-gateway part.   When mode&#x3D;public, the Tyk OAS API spec will exclude the x-tyk-api-gateway part in the response. | 

### Return type

[**[]ListApisOAS200ResponseInner**](ListApisOAS200ResponseInner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOASApiVersions

> VersionMetas ListOASApiVersions(ctx, apiID).SearchText(searchText).AccessType(accessType).Execute()

Listing versions of a Tyk OAS API.



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
	apiID := "4c1c0d8fc885401053ddac4e39ef676b" // string | ID of the API you want to fetch.
	searchText := "Sample oas" // string | Search for API version name (optional)
	accessType := "internal" // string | Filter for internal or external API versions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TykOASAPIsAPI.ListOASApiVersions(context.Background(), apiID).SearchText(searchText).AccessType(accessType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TykOASAPIsAPI.ListOASApiVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOASApiVersions`: VersionMetas
	fmt.Fprintf(os.Stdout, "Response from `TykOASAPIsAPI.ListOASApiVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | ID of the API you want to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOASApiVersionsRequest struct via the builder pattern


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


## PatchApiOAS

> ApiModifyKeySuccess PatchApiOAS(ctx, apiID).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).AllowList(allowList).ValidateRequest(validateRequest).MockResponse(mockResponse).Authentication(authentication).Model30(model30).Execute()

Patch API in Tyk OAS format.



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
	apiID := "4c1c0d8fc885401053ddac4e39ef676b" // string | ID of the API you want to fetch.
	upstreamURL := "https://localhost:8080" // string | Upstream URL for the API (optional)
	listenPath := "/user-test/" // string | Listen path for the API (optional)
	customDomain := "tyk.io" // string | Custom domain for the API (optional)
	allowList := true // bool | Enable allowList middleware for all endpoints (optional)
	validateRequest := true // bool | Enable validateRequest middleware for all endpoints having a request body with media type application/json (optional)
	mockResponse := true // bool | Enable mockResponse middleware for all endpoints having responses configured. (optional)
	authentication := true // bool | Enable/disable the authentication mechanism in your Tyk Gateway for your OAS API (optional)
	model30 := *openapiclient.NewModel30("Openapi_example", *openapiclient.NewInfo1("Title_example", "Version_example"), map[string]interface{}(123)) // Model30 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TykOASAPIsAPI.PatchApiOAS(context.Background(), apiID).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).AllowList(allowList).ValidateRequest(validateRequest).MockResponse(mockResponse).Authentication(authentication).Model30(model30).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TykOASAPIsAPI.PatchApiOAS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchApiOAS`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `TykOASAPIsAPI.PatchApiOAS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | ID of the API you want to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upstreamURL** | **string** | Upstream URL for the API | 
 **listenPath** | **string** | Listen path for the API | 
 **customDomain** | **string** | Custom domain for the API | 
 **allowList** | **bool** | Enable allowList middleware for all endpoints | 
 **validateRequest** | **bool** | Enable validateRequest middleware for all endpoints having a request body with media type application/json | 
 **mockResponse** | **bool** | Enable mockResponse middleware for all endpoints having responses configured. | 
 **authentication** | **bool** | Enable/disable the authentication mechanism in your Tyk Gateway for your OAS API | 
 **model30** | [**Model30**](Model30.md) |  | 

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


## UpdateApiOAS

> ApiModifyKeySuccess UpdateApiOAS(ctx, apiID).ListApisOAS200ResponseInner(listApisOAS200ResponseInner).Execute()

Update a Tyk OAS API definition.



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
	apiID := "4c1c0d8fc885401053ddac4e39ef676b" // string | ID of the API you want to fetch
	listApisOAS200ResponseInner := *openapiclient.NewListApisOAS200ResponseInner("Openapi_example", *openapiclient.NewInfo(), map[string]interface{}(123)) // ListApisOAS200ResponseInner |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TykOASAPIsAPI.UpdateApiOAS(context.Background(), apiID).ListApisOAS200ResponseInner(listApisOAS200ResponseInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TykOASAPIsAPI.UpdateApiOAS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApiOAS`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `TykOASAPIsAPI.UpdateApiOAS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | ID of the API you want to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **listApisOAS200ResponseInner** | [**ListApisOAS200ResponseInner**](ListApisOAS200ResponseInner.md) |  | 

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

