# \SchemaApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSchema**](SchemaApi.md#GetSchema) | **Get** /tyk/schema | 



## GetSchema

> OASSchemaResponse GetSchema(ctx).OasVersion(oasVersion).Execute()





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
    oasVersion := "oasVersion_example" // string | The OAS version (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemaApi.GetSchema(context.Background()).OasVersion(oasVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaApi.GetSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchema`: OASSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `SchemaApi.GetSchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oasVersion** | **string** | The OAS version | 

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

