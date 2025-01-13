# \CertsAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCert**](CertsAPI.md#AddCert) | **Post** /tyk/certs | Add a certificate.
[**DeleteCerts**](CertsAPI.md#DeleteCerts) | **Delete** /tyk/certs/{certID} | Delete certificate.
[**ListCerts**](CertsAPI.md#ListCerts) | **Get** /tyk/certs | List certificates.



## AddCert

> APICertificateStatusMessage AddCert(ctx).OrgId(orgId).Body(body).Execute()

Add a certificate.



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
	orgId := "5e9d9544a1dcd60001d0ed20" // string | Organisation ID to add the certificate to. (optional)
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertsAPI.AddCert(context.Background()).OrgId(orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertsAPI.AddCert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddCert`: APICertificateStatusMessage
	fmt.Fprintf(os.Stdout, "Response from `CertsAPI.AddCert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **string** | Organisation ID to add the certificate to. | 
 **body** | **string** |  | 

### Return type

[**APICertificateStatusMessage**](APICertificateStatusMessage.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCerts

> ApiStatusMessage DeleteCerts(ctx, certID).OrgId(orgId).Execute()

Delete certificate.



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
	certID := "5e9d9544a1dcd60001d0ed20a6ab77653d5da938f452bb8cc9b55b0630a6743dabd8dc92bfb025abb09ce035" // string | Certificate ID to be deleted.
	orgId := "5e9d9544a1dcd60001d0ed20" // string | Organisation ID to delete the certificates from. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertsAPI.DeleteCerts(context.Background(), certID).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertsAPI.DeleteCerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCerts`: ApiStatusMessage
	fmt.Fprintf(os.Stdout, "Response from `CertsAPI.DeleteCerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certID** | **string** | Certificate ID to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgId** | **string** | Organisation ID to delete the certificates from. | 

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


## ListCerts

> ListCerts200Response ListCerts(ctx).OrgId(orgId).Mode(mode).Execute()

List certificates.



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
	orgId := "5e9d9544a1dcd60001d0ed20" // string | Organisation ID to list the certificates. (optional)
	mode := "detailed" // string | Mode to list the certificate details. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertsAPI.ListCerts(context.Background()).OrgId(orgId).Mode(mode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertsAPI.ListCerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCerts`: ListCerts200Response
	fmt.Fprintf(os.Stdout, "Response from `CertsAPI.ListCerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **string** | Organisation ID to list the certificates. | 
 **mode** | **string** | Mode to list the certificate details. | 

### Return type

[**ListCerts200Response**](ListCerts200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

