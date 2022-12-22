# \OASAPIsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiOAS**](OASAPIsApi.md#CreateApiOAS) | **Post** /tyk/apis/oas | 
[**DeleteOASApi**](OASAPIsApi.md#DeleteOASApi) | **Delete** /tyk/apis/oas/{apiID} | 
[**DownloadApiOASPublic**](OASAPIsApi.md#DownloadApiOASPublic) | **Get** /tyk/apis/oas/{apiID}/export | 
[**DownloadApisOASPublic**](OASAPIsApi.md#DownloadApisOASPublic) | **Get** /tyk/apis/oas/export | 
[**ImportOAS**](OASAPIsApi.md#ImportOAS) | **Post** /tyk/apis/oas/import | 
[**ListApiOAS**](OASAPIsApi.md#ListApiOAS) | **Get** /tyk/apis/oas/{apiID} | 
[**ListApisOAS**](OASAPIsApi.md#ListApisOAS) | **Get** /tyk/apis/oas | 
[**PatchApiOAS**](OASAPIsApi.md#PatchApiOAS) | **Patch** /tyk/apis/oas/{apiID} | Patch a single OAS API by ID
[**UpdateApiOAS**](OASAPIsApi.md#UpdateApiOAS) | **Put** /tyk/apis/oas/{apiID} | 



## CreateApiOAS

> ApiModifyKeySuccess CreateApiOAS(ctx).Schema(schema).Execute()





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
    schema := *openapiclient.NewSchema("Openapi_example", *openapiclient.NewInfo("Title_example", "Version_example"), map[string]interface{}(123)) // Schema |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OASAPIsApi.CreateApiOAS(context.Background()).Schema(schema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsApi.CreateApiOAS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiOAS`: ApiModifyKeySuccess
    fmt.Fprintf(os.Stdout, "Response from `OASAPIsApi.CreateApiOAS`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schema** | [**Schema**](Schema.md) |  | 

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

> ApiStatusMessage DeleteOASApi(ctx, apiID).Execute()





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
    resp, r, err := apiClient.OASAPIsApi.DeleteOASApi(context.Background(), apiID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsApi.DeleteOASApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOASApi`: ApiStatusMessage
    fmt.Fprintf(os.Stdout, "Response from `OASAPIsApi.DeleteOASApi`: %v\n", resp)
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

[**ApiStatusMessage**](ApiStatusMessage.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadApiOASPublic

> OASSchemaResponse DownloadApiOASPublic(ctx, apiID).Mode(mode).Execute()





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
    resp, r, err := apiClient.OASAPIsApi.DownloadApiOASPublic(context.Background(), apiID).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsApi.DownloadApiOASPublic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadApiOASPublic`: OASSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `OASAPIsApi.DownloadApiOASPublic`: %v\n", resp)
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

[**OASSchemaResponse**](OASSchemaResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadApisOASPublic

> []OASSchemaResponse DownloadApisOASPublic(ctx).Mode(mode).Execute()





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
    resp, r, err := apiClient.OASAPIsApi.DownloadApisOASPublic(context.Background()).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsApi.DownloadApisOASPublic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadApisOASPublic`: []OASSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `OASAPIsApi.DownloadApisOASPublic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadApisOASPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mode** | **string** | The mode of OAS export. By default the mode is not set which means the OAS spec is exported including the OAS Tyk extension.   If the mode is set to public, the OAS spec excluding the Tyk extension is exported. | 

### Return type

[**[]OASSchemaResponse**](OASSchemaResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportOAS

> ApiModifyKeySuccess ImportOAS(ctx).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).ApiID(apiID).AllowList(allowList).ValidateRequest(validateRequest).Authentication(authentication).Schema(schema).Execute()





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
    allowList := openapiclient.BooleanQueryParam("true") // BooleanQueryParam | Enable allowList middleware for all endpoints (optional)
    validateRequest := openapiclient.BooleanQueryParam("true") // BooleanQueryParam | Enable validateRequest middleware for all endpoints having a request body with media type application/json (optional)
    authentication := openapiclient.BooleanQueryParam("true") // BooleanQueryParam | Enable or disable authentication in your Tyk Gateway as per your OAS document. (optional)
    schema := *openapiclient.NewSchema("Openapi_example", *openapiclient.NewInfo("Title_example", "Version_example"), map[string]interface{}(123)) // Schema |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OASAPIsApi.ImportOAS(context.Background()).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).ApiID(apiID).AllowList(allowList).ValidateRequest(validateRequest).Authentication(authentication).Schema(schema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsApi.ImportOAS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportOAS`: ApiModifyKeySuccess
    fmt.Fprintf(os.Stdout, "Response from `OASAPIsApi.ImportOAS`: %v\n", resp)
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
 **allowList** | [**BooleanQueryParam**](BooleanQueryParam.md) | Enable allowList middleware for all endpoints | 
 **validateRequest** | [**BooleanQueryParam**](BooleanQueryParam.md) | Enable validateRequest middleware for all endpoints having a request body with media type application/json | 
 **authentication** | [**BooleanQueryParam**](BooleanQueryParam.md) | Enable or disable authentication in your Tyk Gateway as per your OAS document. | 
 **schema** | [**Schema**](Schema.md) |  | 

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


## ListApiOAS

> OASSchemaResponse ListApiOAS(ctx, apiID).Mode(mode).Execute()





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
    resp, r, err := apiClient.OASAPIsApi.ListApiOAS(context.Background(), apiID).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsApi.ListApiOAS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiOAS`: OASSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `OASAPIsApi.ListApiOAS`: %v\n", resp)
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

[**OASSchemaResponse**](OASSchemaResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApisOAS

> []OASSchemaResponse ListApisOAS(ctx).Mode(mode).Execute()





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
    resp, r, err := apiClient.OASAPIsApi.ListApisOAS(context.Background()).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsApi.ListApisOAS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApisOAS`: []OASSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `OASAPIsApi.ListApisOAS`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApisOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mode** | **string** | Mode of OAS get, by default mode could be empty which means to get OAS spec including OAS Tyk extension.  When mode&#x3D;public, OAS spec excluding Tyk extension will be returned in the response | 

### Return type

[**[]OASSchemaResponse**](OASSchemaResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApiOAS

> ApiModifyKeySuccess PatchApiOAS(ctx, apiID).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).ValidateRequest(validateRequest).AllowList(allowList).Authentication(authentication).Schema(schema).Execute()

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
    validateRequest := openapiclient.BooleanQueryParam("true") // BooleanQueryParam | Enable validateRequest middleware for all endpoints having a request body with media type application/json (optional)
    allowList := openapiclient.BooleanQueryParam("true") // BooleanQueryParam | Enable allowList middleware for all endpoints (optional)
    authentication := openapiclient.BooleanQueryParam("true") // BooleanQueryParam | Enable or disable authentication in your Tyk Gateway as per your OAS document. (optional)
    schema := *openapiclient.NewSchema("Openapi_example", *openapiclient.NewInfo("Title_example", "Version_example"), map[string]interface{}(123)) // Schema |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OASAPIsApi.PatchApiOAS(context.Background(), apiID).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).ValidateRequest(validateRequest).AllowList(allowList).Authentication(authentication).Schema(schema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsApi.PatchApiOAS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchApiOAS`: ApiModifyKeySuccess
    fmt.Fprintf(os.Stdout, "Response from `OASAPIsApi.PatchApiOAS`: %v\n", resp)
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
 **validateRequest** | [**BooleanQueryParam**](BooleanQueryParam.md) | Enable validateRequest middleware for all endpoints having a request body with media type application/json | 
 **allowList** | [**BooleanQueryParam**](BooleanQueryParam.md) | Enable allowList middleware for all endpoints | 
 **authentication** | [**BooleanQueryParam**](BooleanQueryParam.md) | Enable or disable authentication in your Tyk Gateway as per your OAS document. | 
 **schema** | [**Schema**](Schema.md) |  | 

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

> ApiModifyKeySuccess UpdateApiOAS(ctx, apiID).Schema(schema).Execute()





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
    schema := *openapiclient.NewSchema("Openapi_example", *openapiclient.NewInfo("Title_example", "Version_example"), map[string]interface{}(123)) // Schema |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OASAPIsApi.UpdateApiOAS(context.Background(), apiID).Schema(schema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsApi.UpdateApiOAS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiOAS`: ApiModifyKeySuccess
    fmt.Fprintf(os.Stdout, "Response from `OASAPIsApi.UpdateApiOAS`: %v\n", resp)
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

 **schema** | [**Schema**](Schema.md) |  | 

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

