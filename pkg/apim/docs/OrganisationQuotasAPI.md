# \OrganisationQuotasAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrgKey**](OrganisationQuotasAPI.md#AddOrgKey) | **Post** /tyk/org/keys/{keyID} | Create an organisation key
[**DeleteOrgKey**](OrganisationQuotasAPI.md#DeleteOrgKey) | **Delete** /tyk/org/keys/{keyID} | Delete Key
[**GetOrgKey**](OrganisationQuotasAPI.md#GetOrgKey) | **Get** /tyk/org/keys/{keyID} | Get an Organisation Key
[**ListOrgKeys**](OrganisationQuotasAPI.md#ListOrgKeys) | **Get** /tyk/org/keys | List Organisation Keys
[**UpdateOrgKey**](OrganisationQuotasAPI.md#UpdateOrgKey) | **Put** /tyk/org/keys/{keyID} | Update Organisation Key



## AddOrgKey

> ApiModifyKeySuccess AddOrgKey(ctx, keyID).ResetQuota(resetQuota).SessionState(sessionState).Execute()

Create an organisation key



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
	keyID := "e389ae00a2b145feaf28d6cc11f0f86d" // string | The Key ID
	resetQuota := "1" // string | Adding the reset_quota parameter and setting it to 1, will cause Tyk reset the organisations quota in the live quota manager, it is recommended to use this mechanism to reset organisation-level access if a monthly subscription is in place. (optional)
	sessionState := *openapiclient.NewSessionState() // SessionState |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganisationQuotasAPI.AddOrgKey(context.Background(), keyID).ResetQuota(resetQuota).SessionState(sessionState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganisationQuotasAPI.AddOrgKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddOrgKey`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `OrganisationQuotasAPI.AddOrgKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyID** | **string** | The Key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOrgKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resetQuota** | **string** | Adding the reset_quota parameter and setting it to 1, will cause Tyk reset the organisations quota in the live quota manager, it is recommended to use this mechanism to reset organisation-level access if a monthly subscription is in place. | 
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


## DeleteOrgKey

> ApiModifyKeySuccess DeleteOrgKey(ctx, keyID).Execute()

Delete Key



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
	keyID := "e389ae00a2b145feaf28d6cc11f0f86d" // string | The Key ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganisationQuotasAPI.DeleteOrgKey(context.Background(), keyID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganisationQuotasAPI.DeleteOrgKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrgKey`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `OrganisationQuotasAPI.DeleteOrgKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyID** | **string** | The Key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgKeyRequest struct via the builder pattern


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


## GetOrgKey

> GetOrgKey200Response GetOrgKey(ctx, keyID).OrgID(orgID).Execute()

Get an Organisation Key



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
	keyID := "e389ae00a2b145feaf28d6cc11f0f86d" // string | The Key ID
	orgID := "664a14650619d40001f1f00f" // string | The Org ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganisationQuotasAPI.GetOrgKey(context.Background(), keyID).OrgID(orgID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganisationQuotasAPI.GetOrgKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgKey`: GetOrgKey200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganisationQuotasAPI.GetOrgKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyID** | **string** | The Key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgID** | **string** | The Org ID | 

### Return type

[**GetOrgKey200Response**](GetOrgKey200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgKeys

> ApiAllKeys ListOrgKeys(ctx).Filter(filter).Execute()

List Organisation Keys



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
	filter := "default*" // string | Retrieves all keys starting with the specified filter(filter is a prefix - e.g. default* or default will return all keys starting with default  like defaultbd,defaulttwo etc).We don't use filter for hashed keys (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganisationQuotasAPI.ListOrgKeys(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganisationQuotasAPI.ListOrgKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgKeys`: ApiAllKeys
	fmt.Fprintf(os.Stdout, "Response from `OrganisationQuotasAPI.ListOrgKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrgKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Retrieves all keys starting with the specified filter(filter is a prefix - e.g. default* or default will return all keys starting with default  like defaultbd,defaulttwo etc).We don&#39;t use filter for hashed keys | 

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


## UpdateOrgKey

> ApiModifyKeySuccess UpdateOrgKey(ctx, keyID).ResetQuota(resetQuota).SessionState(sessionState).Execute()

Update Organisation Key



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
	keyID := "e389ae00a2b145feaf28d6cc11f0f86d" // string | The Key ID
	resetQuota := "1" // string | Adding the reset_quota parameter and setting it to 1, will cause Tyk reset the organisations quota in the live quota manager, it is recommended to use this mechanism to reset organisation-level access if a monthly subscription is in place. (optional)
	sessionState := *openapiclient.NewSessionState() // SessionState |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganisationQuotasAPI.UpdateOrgKey(context.Background(), keyID).ResetQuota(resetQuota).SessionState(sessionState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganisationQuotasAPI.UpdateOrgKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgKey`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `OrganisationQuotasAPI.UpdateOrgKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyID** | **string** | The Key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resetQuota** | **string** | Adding the reset_quota parameter and setting it to 1, will cause Tyk reset the organisations quota in the live quota manager, it is recommended to use this mechanism to reset organisation-level access if a monthly subscription is in place. | 
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

