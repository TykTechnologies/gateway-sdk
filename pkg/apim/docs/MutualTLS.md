# MutualTLS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainToCertificateMapping** | Pointer to [**[]DomainToCertificate**](DomainToCertificate.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewMutualTLS

`func NewMutualTLS() *MutualTLS`

NewMutualTLS instantiates a new MutualTLS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualTLSWithDefaults

`func NewMutualTLSWithDefaults() *MutualTLS`

NewMutualTLSWithDefaults instantiates a new MutualTLS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainToCertificateMapping

`func (o *MutualTLS) GetDomainToCertificateMapping() []DomainToCertificate`

GetDomainToCertificateMapping returns the DomainToCertificateMapping field if non-nil, zero value otherwise.

### GetDomainToCertificateMappingOk

`func (o *MutualTLS) GetDomainToCertificateMappingOk() (*[]DomainToCertificate, bool)`

GetDomainToCertificateMappingOk returns a tuple with the DomainToCertificateMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainToCertificateMapping

`func (o *MutualTLS) SetDomainToCertificateMapping(v []DomainToCertificate)`

SetDomainToCertificateMapping sets DomainToCertificateMapping field to given value.

### HasDomainToCertificateMapping

`func (o *MutualTLS) HasDomainToCertificateMapping() bool`

HasDomainToCertificateMapping returns a boolean if a field has been set.

### SetDomainToCertificateMappingNil

`func (o *MutualTLS) SetDomainToCertificateMappingNil(b bool)`

 SetDomainToCertificateMappingNil sets the value for DomainToCertificateMapping to be an explicit nil

### UnsetDomainToCertificateMapping
`func (o *MutualTLS) UnsetDomainToCertificateMapping()`

UnsetDomainToCertificateMapping ensures that no value is present for DomainToCertificateMapping, not even an explicit nil
### GetEnabled

`func (o *MutualTLS) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MutualTLS) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MutualTLS) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MutualTLS) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


