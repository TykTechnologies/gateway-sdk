# CertsCertificateBasics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsNames** | Pointer to **[]string** |  | [optional] 
**HasPrivate** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsCa** | Pointer to **bool** |  | [optional] 
**IssuerCn** | Pointer to **string** |  | [optional] 
**NotAfter** | Pointer to **time.Time** |  | [optional] 
**NotBefore** | Pointer to **time.Time** |  | [optional] 
**SubjectCn** | Pointer to **string** |  | [optional] 

## Methods

### NewCertsCertificateBasics

`func NewCertsCertificateBasics() *CertsCertificateBasics`

NewCertsCertificateBasics instantiates a new CertsCertificateBasics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertsCertificateBasicsWithDefaults

`func NewCertsCertificateBasicsWithDefaults() *CertsCertificateBasics`

NewCertsCertificateBasicsWithDefaults instantiates a new CertsCertificateBasics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsNames

`func (o *CertsCertificateBasics) GetDnsNames() []string`

GetDnsNames returns the DnsNames field if non-nil, zero value otherwise.

### GetDnsNamesOk

`func (o *CertsCertificateBasics) GetDnsNamesOk() (*[]string, bool)`

GetDnsNamesOk returns a tuple with the DnsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNames

`func (o *CertsCertificateBasics) SetDnsNames(v []string)`

SetDnsNames sets DnsNames field to given value.

### HasDnsNames

`func (o *CertsCertificateBasics) HasDnsNames() bool`

HasDnsNames returns a boolean if a field has been set.

### SetDnsNamesNil

`func (o *CertsCertificateBasics) SetDnsNamesNil(b bool)`

 SetDnsNamesNil sets the value for DnsNames to be an explicit nil

### UnsetDnsNames
`func (o *CertsCertificateBasics) UnsetDnsNames()`

UnsetDnsNames ensures that no value is present for DnsNames, not even an explicit nil
### GetHasPrivate

`func (o *CertsCertificateBasics) GetHasPrivate() bool`

GetHasPrivate returns the HasPrivate field if non-nil, zero value otherwise.

### GetHasPrivateOk

`func (o *CertsCertificateBasics) GetHasPrivateOk() (*bool, bool)`

GetHasPrivateOk returns a tuple with the HasPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrivate

`func (o *CertsCertificateBasics) SetHasPrivate(v bool)`

SetHasPrivate sets HasPrivate field to given value.

### HasHasPrivate

`func (o *CertsCertificateBasics) HasHasPrivate() bool`

HasHasPrivate returns a boolean if a field has been set.

### GetId

`func (o *CertsCertificateBasics) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertsCertificateBasics) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertsCertificateBasics) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CertsCertificateBasics) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsCa

`func (o *CertsCertificateBasics) GetIsCa() bool`

GetIsCa returns the IsCa field if non-nil, zero value otherwise.

### GetIsCaOk

`func (o *CertsCertificateBasics) GetIsCaOk() (*bool, bool)`

GetIsCaOk returns a tuple with the IsCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCa

`func (o *CertsCertificateBasics) SetIsCa(v bool)`

SetIsCa sets IsCa field to given value.

### HasIsCa

`func (o *CertsCertificateBasics) HasIsCa() bool`

HasIsCa returns a boolean if a field has been set.

### GetIssuerCn

`func (o *CertsCertificateBasics) GetIssuerCn() string`

GetIssuerCn returns the IssuerCn field if non-nil, zero value otherwise.

### GetIssuerCnOk

`func (o *CertsCertificateBasics) GetIssuerCnOk() (*string, bool)`

GetIssuerCnOk returns a tuple with the IssuerCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCn

`func (o *CertsCertificateBasics) SetIssuerCn(v string)`

SetIssuerCn sets IssuerCn field to given value.

### HasIssuerCn

`func (o *CertsCertificateBasics) HasIssuerCn() bool`

HasIssuerCn returns a boolean if a field has been set.

### GetNotAfter

`func (o *CertsCertificateBasics) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CertsCertificateBasics) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CertsCertificateBasics) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *CertsCertificateBasics) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetNotBefore

`func (o *CertsCertificateBasics) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CertsCertificateBasics) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CertsCertificateBasics) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *CertsCertificateBasics) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetSubjectCn

`func (o *CertsCertificateBasics) GetSubjectCn() string`

GetSubjectCn returns the SubjectCn field if non-nil, zero value otherwise.

### GetSubjectCnOk

`func (o *CertsCertificateBasics) GetSubjectCnOk() (*string, bool)`

GetSubjectCnOk returns a tuple with the SubjectCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCn

`func (o *CertsCertificateBasics) SetSubjectCn(v string)`

SetSubjectCn sets SubjectCn field to given value.

### HasSubjectCn

`func (o *CertsCertificateBasics) HasSubjectCn() bool`

HasSubjectCn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


