# CertificateMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**HasPrivate** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to [**PkixName**](PkixName.md) |  | [optional] 
**Subject** | Pointer to [**PkixName**](PkixName.md) |  | [optional] 
**NotBefore** | Pointer to **string** |  | [optional] 
**NotAfter** | Pointer to **string** |  | [optional] 
**DnsNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCertificateMeta

`func NewCertificateMeta() *CertificateMeta`

NewCertificateMeta instantiates a new CertificateMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateMetaWithDefaults

`func NewCertificateMetaWithDefaults() *CertificateMeta`

NewCertificateMetaWithDefaults instantiates a new CertificateMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateMeta) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateMeta) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateMeta) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CertificateMeta) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFingerprint

`func (o *CertificateMeta) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *CertificateMeta) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *CertificateMeta) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *CertificateMeta) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetHasPrivate

`func (o *CertificateMeta) GetHasPrivate() string`

GetHasPrivate returns the HasPrivate field if non-nil, zero value otherwise.

### GetHasPrivateOk

`func (o *CertificateMeta) GetHasPrivateOk() (*string, bool)`

GetHasPrivateOk returns a tuple with the HasPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrivate

`func (o *CertificateMeta) SetHasPrivate(v string)`

SetHasPrivate sets HasPrivate field to given value.

### HasHasPrivate

`func (o *CertificateMeta) HasHasPrivate() bool`

HasHasPrivate returns a boolean if a field has been set.

### GetIssuer

`func (o *CertificateMeta) GetIssuer() PkixName`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CertificateMeta) GetIssuerOk() (*PkixName, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CertificateMeta) SetIssuer(v PkixName)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *CertificateMeta) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSubject

`func (o *CertificateMeta) GetSubject() PkixName`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CertificateMeta) GetSubjectOk() (*PkixName, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CertificateMeta) SetSubject(v PkixName)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CertificateMeta) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetNotBefore

`func (o *CertificateMeta) GetNotBefore() string`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CertificateMeta) GetNotBeforeOk() (*string, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CertificateMeta) SetNotBefore(v string)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *CertificateMeta) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetNotAfter

`func (o *CertificateMeta) GetNotAfter() string`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CertificateMeta) GetNotAfterOk() (*string, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CertificateMeta) SetNotAfter(v string)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *CertificateMeta) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetDnsNames

`func (o *CertificateMeta) GetDnsNames() []string`

GetDnsNames returns the DnsNames field if non-nil, zero value otherwise.

### GetDnsNamesOk

`func (o *CertificateMeta) GetDnsNamesOk() (*[]string, bool)`

GetDnsNamesOk returns a tuple with the DnsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNames

`func (o *CertificateMeta) SetDnsNames(v []string)`

SetDnsNames sets DnsNames field to given value.

### HasDnsNames

`func (o *CertificateMeta) HasDnsNames() bool`

HasDnsNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


