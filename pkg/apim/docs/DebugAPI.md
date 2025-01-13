# \DebugAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DebugApiDefinition**](DebugAPI.md#DebugApiDefinition) | **Post** /tyk/debug | Test an an API definition.



## DebugApiDefinition

> TraceResponse DebugApiDefinition(ctx).TraceRequest(traceRequest).Execute()

Test an an API definition.



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
	traceRequest := *openapiclient.NewTraceRequest() // TraceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DebugAPI.DebugApiDefinition(context.Background()).TraceRequest(traceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DebugAPI.DebugApiDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DebugApiDefinition`: TraceResponse
	fmt.Fprintf(os.Stdout, "Response from `DebugAPI.DebugApiDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDebugApiDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **traceRequest** | [**TraceRequest**](TraceRequest.md) |  | 

### Return type

[**TraceResponse**](TraceResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

