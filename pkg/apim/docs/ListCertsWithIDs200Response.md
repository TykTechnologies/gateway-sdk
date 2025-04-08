# ListCertsWithIDs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsNames** | Pointer to **[]string** |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**HasPrivate** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsCa** | Pointer to **bool** |  | [optional] 
**Issuer** | Pointer to **map[string]interface{}** |  | [optional] 
**NotAfter** | Pointer to **time.Time** |  | [optional] 
**NotBefore** | Pointer to **time.Time** |  | [optional] 
**Subject** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewListCertsWithIDs200Response

`func NewListCertsWithIDs200Response() *ListCertsWithIDs200Response`

NewListCertsWithIDs200Response instantiates a new ListCertsWithIDs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCertsWithIDs200ResponseWithDefaults

`func NewListCertsWithIDs200ResponseWithDefaults() *ListCertsWithIDs200Response`

NewListCertsWithIDs200ResponseWithDefaults instantiates a new ListCertsWithIDs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsNames

`func (o *ListCertsWithIDs200Response) GetDnsNames() []string`

GetDnsNames returns the DnsNames field if non-nil, zero value otherwise.

### GetDnsNamesOk

`func (o *ListCertsWithIDs200Response) GetDnsNamesOk() (*[]string, bool)`

GetDnsNamesOk returns a tuple with the DnsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNames

`func (o *ListCertsWithIDs200Response) SetDnsNames(v []string)`

SetDnsNames sets DnsNames field to given value.

### HasDnsNames

`func (o *ListCertsWithIDs200Response) HasDnsNames() bool`

HasDnsNames returns a boolean if a field has been set.

### GetFingerprint

`func (o *ListCertsWithIDs200Response) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ListCertsWithIDs200Response) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ListCertsWithIDs200Response) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *ListCertsWithIDs200Response) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetHasPrivate

`func (o *ListCertsWithIDs200Response) GetHasPrivate() bool`

GetHasPrivate returns the HasPrivate field if non-nil, zero value otherwise.

### GetHasPrivateOk

`func (o *ListCertsWithIDs200Response) GetHasPrivateOk() (*bool, bool)`

GetHasPrivateOk returns a tuple with the HasPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrivate

`func (o *ListCertsWithIDs200Response) SetHasPrivate(v bool)`

SetHasPrivate sets HasPrivate field to given value.

### HasHasPrivate

`func (o *ListCertsWithIDs200Response) HasHasPrivate() bool`

HasHasPrivate returns a boolean if a field has been set.

### GetId

`func (o *ListCertsWithIDs200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCertsWithIDs200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCertsWithIDs200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListCertsWithIDs200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsCa

`func (o *ListCertsWithIDs200Response) GetIsCa() bool`

GetIsCa returns the IsCa field if non-nil, zero value otherwise.

### GetIsCaOk

`func (o *ListCertsWithIDs200Response) GetIsCaOk() (*bool, bool)`

GetIsCaOk returns a tuple with the IsCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCa

`func (o *ListCertsWithIDs200Response) SetIsCa(v bool)`

SetIsCa sets IsCa field to given value.

### HasIsCa

`func (o *ListCertsWithIDs200Response) HasIsCa() bool`

HasIsCa returns a boolean if a field has been set.

### GetIssuer

`func (o *ListCertsWithIDs200Response) GetIssuer() map[string]interface{}`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ListCertsWithIDs200Response) GetIssuerOk() (*map[string]interface{}, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ListCertsWithIDs200Response) SetIssuer(v map[string]interface{})`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *ListCertsWithIDs200Response) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetNotAfter

`func (o *ListCertsWithIDs200Response) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *ListCertsWithIDs200Response) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *ListCertsWithIDs200Response) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *ListCertsWithIDs200Response) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetNotBefore

`func (o *ListCertsWithIDs200Response) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *ListCertsWithIDs200Response) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *ListCertsWithIDs200Response) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *ListCertsWithIDs200Response) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetSubject

`func (o *ListCertsWithIDs200Response) GetSubject() map[string]interface{}`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ListCertsWithIDs200Response) GetSubjectOk() (*map[string]interface{}, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ListCertsWithIDs200Response) SetSubject(v map[string]interface{})`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ListCertsWithIDs200Response) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


