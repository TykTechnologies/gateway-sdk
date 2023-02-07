# CertificateMetaModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**HasPrivate** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to [**PkixNameModelModel**](PkixNameModel.md) |  | [optional] 
**Subject** | Pointer to [**PkixNameModelModel**](PkixNameModel.md) |  | [optional] 
**NotBefore** | Pointer to **string** |  | [optional] 
**NotAfter** | Pointer to **string** |  | [optional] 
**DnsNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCertificateMetaModel

`func NewCertificateMetaModel() *CertificateMetaModel`

NewCertificateMetaModel instantiates a new CertificateMetaModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateMetaModelWithDefaults

`func NewCertificateMetaModelWithDefaults() *CertificateMetaModel`

NewCertificateMetaModelWithDefaults instantiates a new CertificateMetaModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateMetaModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateMetaModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateMetaModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CertificateMetaModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFingerprint

`func (o *CertificateMetaModel) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *CertificateMetaModel) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *CertificateMetaModel) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *CertificateMetaModel) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetHasPrivate

`func (o *CertificateMetaModel) GetHasPrivate() string`

GetHasPrivate returns the HasPrivate field if non-nil, zero value otherwise.

### GetHasPrivateOk

`func (o *CertificateMetaModel) GetHasPrivateOk() (*string, bool)`

GetHasPrivateOk returns a tuple with the HasPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrivate

`func (o *CertificateMetaModel) SetHasPrivate(v string)`

SetHasPrivate sets HasPrivate field to given value.

### HasHasPrivate

`func (o *CertificateMetaModel) HasHasPrivate() bool`

HasHasPrivate returns a boolean if a field has been set.

### GetIssuer

`func (o *CertificateMetaModel) GetIssuer() PkixNameModelModel`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CertificateMetaModel) GetIssuerOk() (*PkixNameModelModel, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CertificateMetaModel) SetIssuer(v PkixNameModelModel)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *CertificateMetaModel) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSubject

`func (o *CertificateMetaModel) GetSubject() PkixNameModelModel`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CertificateMetaModel) GetSubjectOk() (*PkixNameModelModel, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CertificateMetaModel) SetSubject(v PkixNameModelModel)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CertificateMetaModel) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetNotBefore

`func (o *CertificateMetaModel) GetNotBefore() string`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CertificateMetaModel) GetNotBeforeOk() (*string, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CertificateMetaModel) SetNotBefore(v string)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *CertificateMetaModel) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetNotAfter

`func (o *CertificateMetaModel) GetNotAfter() string`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CertificateMetaModel) GetNotAfterOk() (*string, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CertificateMetaModel) SetNotAfter(v string)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *CertificateMetaModel) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetDnsNames

`func (o *CertificateMetaModel) GetDnsNames() []string`

GetDnsNames returns the DnsNames field if non-nil, zero value otherwise.

### GetDnsNamesOk

`func (o *CertificateMetaModel) GetDnsNamesOk() (*[]string, bool)`

GetDnsNamesOk returns a tuple with the DnsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNames

`func (o *CertificateMetaModel) SetDnsNames(v []string)`

SetDnsNames sets DnsNames field to given value.

### HasDnsNames

`func (o *CertificateMetaModel) HasDnsNames() bool`

HasDnsNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


