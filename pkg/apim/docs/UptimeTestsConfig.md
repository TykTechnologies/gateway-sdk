# UptimeTestsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpireUtimeAfter** | Pointer to **int32** |  | [optional] 
**RecheckWait** | Pointer to **int32** |  | [optional] 
**ServiceDiscovery** | Pointer to [**ServiceDiscoveryConfiguration**](ServiceDiscoveryConfiguration.md) |  | [optional] 

## Methods

### NewUptimeTestsConfig

`func NewUptimeTestsConfig() *UptimeTestsConfig`

NewUptimeTestsConfig instantiates a new UptimeTestsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUptimeTestsConfigWithDefaults

`func NewUptimeTestsConfigWithDefaults() *UptimeTestsConfig`

NewUptimeTestsConfigWithDefaults instantiates a new UptimeTestsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpireUtimeAfter

`func (o *UptimeTestsConfig) GetExpireUtimeAfter() int32`

GetExpireUtimeAfter returns the ExpireUtimeAfter field if non-nil, zero value otherwise.

### GetExpireUtimeAfterOk

`func (o *UptimeTestsConfig) GetExpireUtimeAfterOk() (*int32, bool)`

GetExpireUtimeAfterOk returns a tuple with the ExpireUtimeAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireUtimeAfter

`func (o *UptimeTestsConfig) SetExpireUtimeAfter(v int32)`

SetExpireUtimeAfter sets ExpireUtimeAfter field to given value.

### HasExpireUtimeAfter

`func (o *UptimeTestsConfig) HasExpireUtimeAfter() bool`

HasExpireUtimeAfter returns a boolean if a field has been set.

### GetRecheckWait

`func (o *UptimeTestsConfig) GetRecheckWait() int32`

GetRecheckWait returns the RecheckWait field if non-nil, zero value otherwise.

### GetRecheckWaitOk

`func (o *UptimeTestsConfig) GetRecheckWaitOk() (*int32, bool)`

GetRecheckWaitOk returns a tuple with the RecheckWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecheckWait

`func (o *UptimeTestsConfig) SetRecheckWait(v int32)`

SetRecheckWait sets RecheckWait field to given value.

### HasRecheckWait

`func (o *UptimeTestsConfig) HasRecheckWait() bool`

HasRecheckWait returns a boolean if a field has been set.

### GetServiceDiscovery

`func (o *UptimeTestsConfig) GetServiceDiscovery() ServiceDiscoveryConfiguration`

GetServiceDiscovery returns the ServiceDiscovery field if non-nil, zero value otherwise.

### GetServiceDiscoveryOk

`func (o *UptimeTestsConfig) GetServiceDiscoveryOk() (*ServiceDiscoveryConfiguration, bool)`

GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDiscovery

`func (o *UptimeTestsConfig) SetServiceDiscovery(v ServiceDiscoveryConfiguration)`

SetServiceDiscovery sets ServiceDiscovery field to given value.

### HasServiceDiscovery

`func (o *UptimeTestsConfig) HasServiceDiscovery() bool`

HasServiceDiscovery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


