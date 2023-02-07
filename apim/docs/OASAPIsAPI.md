# \OASAPIsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiOAS**](OASAPIsAPI.md#CreateApiOAS) | **Post** /tyk/apis/oas | 
[**DeleteOASApi**](OASAPIsAPI.md#DeleteOASApi) | **Delete** /tyk/apis/oas/{apiID} | 
[**DownloadApiOASPublic**](OASAPIsAPI.md#DownloadApiOASPublic) | **Get** /tyk/apis/oas/{apiID}/export | 
[**DownloadApisOASPublic**](OASAPIsAPI.md#DownloadApisOASPublic) | **Get** /tyk/apis/oas/export | 
[**ImportOAS**](OASAPIsAPI.md#ImportOAS) | **Post** /tyk/apis/oas/import | 
[**ListApiOAS**](OASAPIsAPI.md#ListApiOAS) | **Get** /tyk/apis/oas/{apiID} | 
[**ListApisOAS**](OASAPIsAPI.md#ListApisOAS) | **Get** /tyk/apis/oas | 
[**PatchApiOAS**](OASAPIsAPI.md#PatchApiOAS) | **Patch** /tyk/apis/oas/{apiID} | Patch a single OAS API by ID
[**UpdateApiOAS**](OASAPIsAPI.md#UpdateApiOAS) | **Put** /tyk/apis/oas/{apiID} | 



## CreateApiOAS

> ApiModifyKeySuccessModelModel CreateApiOAS(ctx).SchemaModel(schemaModel).Execute()





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
    schemaModel := *openapiclient.NewSchemaModel("Openapi_example", "TODO", map[string]interface{}(123)) // SchemaModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OASAPIsAPI.CreateApiOAS(context.Background()).SchemaModel(schemaModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.CreateApiOAS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiOAS`: ApiModifyKeySuccessModelModel
    fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.CreateApiOAS`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaModel** | [**SchemaModel**](SchemaModel.md) |  | 

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


## DeleteOASApi

> ApiStatusMessageModelModel DeleteOASApi(ctx, apiID).Execute()





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
    resp, r, err := apiClient.OASAPIsAPI.DeleteOASApi(context.Background(), apiID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.DeleteOASApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOASApi`: ApiStatusMessageModelModel
    fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.DeleteOASApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOASApiRequest struct via the builder pattern


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


## DownloadApiOASPublic

> OASSchemaResponseModelModel DownloadApiOASPublic(ctx, apiID).Mode(mode).Execute()





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
    mode := "public" // string | Mode of OAS export, by default mode could be empty which means to export OAS spec including OAS Tyk extension.  When mode=public, OAS spec excluding Tyk extension is exported (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OASAPIsAPI.DownloadApiOASPublic(context.Background(), apiID).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.DownloadApiOASPublic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadApiOASPublic`: OASSchemaResponseModelModel
    fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.DownloadApiOASPublic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadApiOASPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mode** | **string** | Mode of OAS export, by default mode could be empty which means to export OAS spec including OAS Tyk extension.  When mode&#x3D;public, OAS spec excluding Tyk extension is exported | 

### Return type

[**OASSchemaResponseModelModel**](OASSchemaResponseModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadApisOASPublic

> []OASSchemaResponseModelModel DownloadApisOASPublic(ctx).Mode(mode).Execute()





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
    mode := "public" // string | The mode of OAS export. By default the mode is not set which means the OAS spec is exported including the OAS Tyk extension.   If the mode is set to public, the OAS spec excluding the Tyk extension is exported. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OASAPIsAPI.DownloadApisOASPublic(context.Background()).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.DownloadApisOASPublic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadApisOASPublic`: []OASSchemaResponseModelModel
    fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.DownloadApisOASPublic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadApisOASPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mode** | **string** | The mode of OAS export. By default the mode is not set which means the OAS spec is exported including the OAS Tyk extension.   If the mode is set to public, the OAS spec excluding the Tyk extension is exported. | 

### Return type

[**[]OASSchemaResponseModelModel**](OASSchemaResponseModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportOAS

> ApiModifyKeySuccessModelModel ImportOAS(ctx).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).ApiID(apiID).AllowList(allowList).ValidateRequest(validateRequest).Authentication(authentication).SchemaModel(schemaModel).Execute()





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
    upstreamURL := "upstreamURL_example" // string | Upstream URL for the API (optional)
    listenPath := "listenPath_example" // string | Listen path for the API (optional)
    customDomain := "customDomain_example" // string | Custom domain for the API (optional)
    apiID := "apiID_example" // string | ID of the API (optional)
    allowList := openapiclient.BooleanQueryParam("true") // BooleanQueryParamModel | Enable allowList middleware for all endpoints (optional)
    validateRequest := openapiclient.BooleanQueryParam("true") // BooleanQueryParamModel | Enable validateRequest middleware for all endpoints having a request body with media type application/json (optional)
    authentication := openapiclient.BooleanQueryParam("true") // BooleanQueryParamModel | Enable or disable authentication in your Tyk Gateway as per your OAS document. (optional)
    schemaModel := *openapiclient.NewSchemaModel("Openapi_example", "TODO", map[string]interface{}(123)) // SchemaModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OASAPIsAPI.ImportOAS(context.Background()).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).ApiID(apiID).AllowList(allowList).ValidateRequest(validateRequest).Authentication(authentication).SchemaModel(schemaModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.ImportOAS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportOAS`: ApiModifyKeySuccessModelModel
    fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.ImportOAS`: %v\n", resp)
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
 **apiID** | **string** | ID of the API | 
 **allowList** | [**BooleanQueryParamModel**](BooleanQueryParamModel.md) | Enable allowList middleware for all endpoints | 
 **validateRequest** | [**BooleanQueryParamModel**](BooleanQueryParamModel.md) | Enable validateRequest middleware for all endpoints having a request body with media type application/json | 
 **authentication** | [**BooleanQueryParamModel**](BooleanQueryParamModel.md) | Enable or disable authentication in your Tyk Gateway as per your OAS document. | 
 **schemaModel** | [**SchemaModel**](SchemaModel.md) |  | 

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


## ListApiOAS

> OASSchemaResponseModelModel ListApiOAS(ctx, apiID).Mode(mode).Execute()





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
    mode := "public" // string | Mode of OAS get, by default mode could be empty which means to get OAS spec including OAS Tyk extension.  When mode=public, OAS spec excluding Tyk extension will be returned in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OASAPIsAPI.ListApiOAS(context.Background(), apiID).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.ListApiOAS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiOAS`: OASSchemaResponseModelModel
    fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.ListApiOAS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApiOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mode** | **string** | Mode of OAS get, by default mode could be empty which means to get OAS spec including OAS Tyk extension.  When mode&#x3D;public, OAS spec excluding Tyk extension will be returned in the response | 

### Return type

[**OASSchemaResponseModelModel**](OASSchemaResponseModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApisOAS

> []OASSchemaResponseModelModel ListApisOAS(ctx).Mode(mode).Execute()





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
    mode := "public" // string | Mode of OAS get, by default mode could be empty which means to get OAS spec including OAS Tyk extension.  When mode=public, OAS spec excluding Tyk extension will be returned in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OASAPIsAPI.ListApisOAS(context.Background()).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.ListApisOAS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApisOAS`: []OASSchemaResponseModelModel
    fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.ListApisOAS`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApisOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mode** | **string** | Mode of OAS get, by default mode could be empty which means to get OAS spec including OAS Tyk extension.  When mode&#x3D;public, OAS spec excluding Tyk extension will be returned in the response | 

### Return type

[**[]OASSchemaResponseModelModel**](OASSchemaResponseModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApiOAS

> ApiModifyKeySuccessModelModel PatchApiOAS(ctx, apiID).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).ValidateRequest(validateRequest).AllowList(allowList).Authentication(authentication).SchemaModel(schemaModel).Execute()

Patch a single OAS API by ID



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
    upstreamURL := "upstreamURL_example" // string | Upstream URL for the API (optional)
    listenPath := "listenPath_example" // string | Listen path for the API (optional)
    customDomain := "customDomain_example" // string | Custom domain for the API (optional)
    validateRequest := openapiclient.BooleanQueryParam("true") // BooleanQueryParamModel | Enable validateRequest middleware for all endpoints having a request body with media type application/json (optional)
    allowList := openapiclient.BooleanQueryParam("true") // BooleanQueryParamModel | Enable allowList middleware for all endpoints (optional)
    authentication := openapiclient.BooleanQueryParam("true") // BooleanQueryParamModel | Enable or disable authentication in your Tyk Gateway as per your OAS document. (optional)
    schemaModel := *openapiclient.NewSchemaModel("Openapi_example", "TODO", map[string]interface{}(123)) // SchemaModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OASAPIsAPI.PatchApiOAS(context.Background(), apiID).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).ValidateRequest(validateRequest).AllowList(allowList).Authentication(authentication).SchemaModel(schemaModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.PatchApiOAS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchApiOAS`: ApiModifyKeySuccessModelModel
    fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.PatchApiOAS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upstreamURL** | **string** | Upstream URL for the API | 
 **listenPath** | **string** | Listen path for the API | 
 **customDomain** | **string** | Custom domain for the API | 
 **validateRequest** | [**BooleanQueryParamModel**](BooleanQueryParamModel.md) | Enable validateRequest middleware for all endpoints having a request body with media type application/json | 
 **allowList** | [**BooleanQueryParamModel**](BooleanQueryParamModel.md) | Enable allowList middleware for all endpoints | 
 **authentication** | [**BooleanQueryParamModel**](BooleanQueryParamModel.md) | Enable or disable authentication in your Tyk Gateway as per your OAS document. | 
 **schemaModel** | [**SchemaModel**](SchemaModel.md) |  | 

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


## UpdateApiOAS

> ApiModifyKeySuccessModelModel UpdateApiOAS(ctx, apiID).SchemaModel(schemaModel).Execute()





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
    schemaModel := *openapiclient.NewSchemaModel("Openapi_example", "TODO", map[string]interface{}(123)) // SchemaModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OASAPIsAPI.UpdateApiOAS(context.Background(), apiID).SchemaModel(schemaModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.UpdateApiOAS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiOAS`: ApiModifyKeySuccessModelModel
    fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.UpdateApiOAS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **schemaModel** | [**SchemaModel**](SchemaModel.md) |  | 

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

