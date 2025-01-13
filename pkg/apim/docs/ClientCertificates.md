# ClientCertificates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowlist** | Pointer to **[]string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewClientCertificates

`func NewClientCertificates() *ClientCertificates`

NewClientCertificates instantiates a new ClientCertificates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCertificatesWithDefaults

`func NewClientCertificatesWithDefaults() *ClientCertificates`

NewClientCertificatesWithDefaults instantiates a new ClientCertificates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowlist

`func (o *ClientCertificates) GetAllowlist() []string`

GetAllowlist returns the Allowlist field if non-nil, zero value otherwise.

### GetAllowlistOk

`func (o *ClientCertificates) GetAllowlistOk() (*[]string, bool)`

GetAllowlistOk returns a tuple with the Allowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlist

`func (o *ClientCertificates) SetAllowlist(v []string)`

SetAllowlist sets Allowlist field to given value.

### HasAllowlist

`func (o *ClientCertificates) HasAllowlist() bool`

HasAllowlist returns a boolean if a field has been set.

### SetAllowlistNil

`func (o *ClientCertificates) SetAllowlistNil(b bool)`

 SetAllowlistNil sets the value for Allowlist to be an explicit nil

### UnsetAllowlist
`func (o *ClientCertificates) UnsetAllowlist()`

UnsetAllowlist ensures that no value is present for Allowlist, not even an explicit nil
### GetEnabled

`func (o *ClientCertificates) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClientCertificates) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClientCertificates) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ClientCertificates) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


