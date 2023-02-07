# \CertsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCert**](CertsAPI.md#AddCert) | **Post** /tyk/certs | Add a certificate
[**DeleteCerts**](CertsAPI.md#DeleteCerts) | **Delete** /tyk/certs | Delete Certificate
[**ListCerts**](CertsAPI.md#ListCerts) | **Get** /tyk/certs | List Certificates



## AddCert

> APICertificateStatusMessageModelModel AddCert(ctx).OrgId(orgId).Body(body).Execute()

Add a certificate



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
    orgId := "orgId_example" // string | Organisation ID to list the certificates
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertsAPI.AddCert(context.Background()).OrgId(orgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertsAPI.AddCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCert`: APICertificateStatusMessageModelModel
    fmt.Fprintf(os.Stdout, "Response from `CertsAPI.AddCert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **string** | Organisation ID to list the certificates | 
 **body** | **string** |  | 

### Return type

[**APICertificateStatusMessageModelModel**](APICertificateStatusMessageModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCerts

> ApiStatusMessageModelModel DeleteCerts(ctx).CertID(certID).OrgId(orgId).Execute()

Delete Certificate



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
    certID := "certID_example" // string | Certifiicate ID to be deleted
    orgId := "orgId_example" // string | Organisation ID to list the certificates

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertsAPI.DeleteCerts(context.Background()).CertID(certID).OrgId(orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertsAPI.DeleteCerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCerts`: ApiStatusMessageModelModel
    fmt.Fprintf(os.Stdout, "Response from `CertsAPI.DeleteCerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certID** | **string** | Certifiicate ID to be deleted | 
 **orgId** | **string** | Organisation ID to list the certificates | 

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


## ListCerts

> ListCerts200ResponseModelModel ListCerts(ctx).OrgId(orgId).Mode(mode).CertID(certID).Execute()

List Certificates



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
    orgId := "orgId_example" // string | Organisation ID to list the certificates
    mode := "detailed" // string | Mode to list the certificate details (optional)
    certID := "e6ce2b49-3e31-44de-95a7-12f054724283,234a37ac-28d1-4f12-b936-ffb4211b79f1" // string | Comma separated list of certificates to list (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertsAPI.ListCerts(context.Background()).OrgId(orgId).Mode(mode).CertID(certID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertsAPI.ListCerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCerts`: ListCerts200ResponseModelModel
    fmt.Fprintf(os.Stdout, "Response from `CertsAPI.ListCerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **string** | Organisation ID to list the certificates | 
 **mode** | **string** | Mode to list the certificate details | 
 **certID** | **string** | Comma separated list of certificates to list | 

### Return type

[**ListCerts200ResponseModelModel**](ListCerts200ResponseModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

