# CertsCertificateMeta

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

### NewCertsCertificateMeta

`func NewCertsCertificateMeta() *CertsCertificateMeta`

NewCertsCertificateMeta instantiates a new CertsCertificateMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertsCertificateMetaWithDefaults

`func NewCertsCertificateMetaWithDefaults() *CertsCertificateMeta`

NewCertsCertificateMetaWithDefaults instantiates a new CertsCertificateMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsNames

`func (o *CertsCertificateMeta) GetDnsNames() []string`

GetDnsNames returns the DnsNames field if non-nil, zero value otherwise.

### GetDnsNamesOk

`func (o *CertsCertificateMeta) GetDnsNamesOk() (*[]string, bool)`

GetDnsNamesOk returns a tuple with the DnsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNames

`func (o *CertsCertificateMeta) SetDnsNames(v []string)`

SetDnsNames sets DnsNames field to given value.

### HasDnsNames

`func (o *CertsCertificateMeta) HasDnsNames() bool`

HasDnsNames returns a boolean if a field has been set.

### GetFingerprint

`func (o *CertsCertificateMeta) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *CertsCertificateMeta) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *CertsCertificateMeta) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *CertsCertificateMeta) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetHasPrivate

`func (o *CertsCertificateMeta) GetHasPrivate() bool`

GetHasPrivate returns the HasPrivate field if non-nil, zero value otherwise.

### GetHasPrivateOk

`func (o *CertsCertificateMeta) GetHasPrivateOk() (*bool, bool)`

GetHasPrivateOk returns a tuple with the HasPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrivate

`func (o *CertsCertificateMeta) SetHasPrivate(v bool)`

SetHasPrivate sets HasPrivate field to given value.

### HasHasPrivate

`func (o *CertsCertificateMeta) HasHasPrivate() bool`

HasHasPrivate returns a boolean if a field has been set.

### GetId

`func (o *CertsCertificateMeta) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertsCertificateMeta) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertsCertificateMeta) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CertsCertificateMeta) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsCa

`func (o *CertsCertificateMeta) GetIsCa() bool`

GetIsCa returns the IsCa field if non-nil, zero value otherwise.

### GetIsCaOk

`func (o *CertsCertificateMeta) GetIsCaOk() (*bool, bool)`

GetIsCaOk returns a tuple with the IsCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCa

`func (o *CertsCertificateMeta) SetIsCa(v bool)`

SetIsCa sets IsCa field to given value.

### HasIsCa

`func (o *CertsCertificateMeta) HasIsCa() bool`

HasIsCa returns a boolean if a field has been set.

### GetIssuer

`func (o *CertsCertificateMeta) GetIssuer() map[string]interface{}`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CertsCertificateMeta) GetIssuerOk() (*map[string]interface{}, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CertsCertificateMeta) SetIssuer(v map[string]interface{})`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *CertsCertificateMeta) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetNotAfter

`func (o *CertsCertificateMeta) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CertsCertificateMeta) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CertsCertificateMeta) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *CertsCertificateMeta) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetNotBefore

`func (o *CertsCertificateMeta) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CertsCertificateMeta) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CertsCertificateMeta) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *CertsCertificateMeta) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetSubject

`func (o *CertsCertificateMeta) GetSubject() map[string]interface{}`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CertsCertificateMeta) GetSubjectOk() (*map[string]interface{}, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CertsCertificateMeta) SetSubject(v map[string]interface{})`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CertsCertificateMeta) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


