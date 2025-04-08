# CertificatePinning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainToPublicKeysMapping** | Pointer to [**[]PinnedPublicKey**](PinnedPublicKey.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCertificatePinning

`func NewCertificatePinning() *CertificatePinning`

NewCertificatePinning instantiates a new CertificatePinning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatePinningWithDefaults

`func NewCertificatePinningWithDefaults() *CertificatePinning`

NewCertificatePinningWithDefaults instantiates a new CertificatePinning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainToPublicKeysMapping

`func (o *CertificatePinning) GetDomainToPublicKeysMapping() []PinnedPublicKey`

GetDomainToPublicKeysMapping returns the DomainToPublicKeysMapping field if non-nil, zero value otherwise.

### GetDomainToPublicKeysMappingOk

`func (o *CertificatePinning) GetDomainToPublicKeysMappingOk() (*[]PinnedPublicKey, bool)`

GetDomainToPublicKeysMappingOk returns a tuple with the DomainToPublicKeysMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainToPublicKeysMapping

`func (o *CertificatePinning) SetDomainToPublicKeysMapping(v []PinnedPublicKey)`

SetDomainToPublicKeysMapping sets DomainToPublicKeysMapping field to given value.

### HasDomainToPublicKeysMapping

`func (o *CertificatePinning) HasDomainToPublicKeysMapping() bool`

HasDomainToPublicKeysMapping returns a boolean if a field has been set.

### SetDomainToPublicKeysMappingNil

`func (o *CertificatePinning) SetDomainToPublicKeysMappingNil(b bool)`

 SetDomainToPublicKeysMappingNil sets the value for DomainToPublicKeysMapping to be an explicit nil

### UnsetDomainToPublicKeysMapping
`func (o *CertificatePinning) UnsetDomainToPublicKeysMapping()`

UnsetDomainToPublicKeysMapping ensures that no value is present for DomainToPublicKeysMapping, not even an explicit nil
### GetEnabled

`func (o *CertificatePinning) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CertificatePinning) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CertificatePinning) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CertificatePinning) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


