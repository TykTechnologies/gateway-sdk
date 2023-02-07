# ListCerts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certs** | Pointer to [**[]CertificateBasics**](CertificateBasics.md) |  | [optional] 

## Methods

### NewListCerts200Response

`func NewListCerts200Response() *ListCerts200Response`

NewListCerts200Response instantiates a new ListCerts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCerts200ResponseWithDefaults

`func NewListCerts200ResponseWithDefaults() *ListCerts200Response`

NewListCerts200ResponseWithDefaults instantiates a new ListCerts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCerts

`func (o *ListCerts200Response) GetCerts() []CertificateBasics`

GetCerts returns the Certs field if non-nil, zero value otherwise.

### GetCertsOk

`func (o *ListCerts200Response) GetCertsOk() (*[]CertificateBasics, bool)`

GetCertsOk returns a tuple with the Certs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCerts

`func (o *ListCerts200Response) SetCerts(v []CertificateBasics)`

SetCerts sets Certs field to given value.

### HasCerts

`func (o *ListCerts200Response) HasCerts() bool`

HasCerts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


