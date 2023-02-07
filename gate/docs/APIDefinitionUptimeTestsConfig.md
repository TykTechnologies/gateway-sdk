# APIDefinitionUptimeTestsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpireUtimeAfter** | Pointer to **int64** |  | [optional] 
**RecheckWait** | Pointer to **int64** |  | [optional] 
**ServiceDiscovery** | Pointer to [**ServiceDiscoveryConfiguration**](ServiceDiscoveryConfiguration.md) |  | [optional] 

## Methods

### NewAPIDefinitionUptimeTestsConfig

`func NewAPIDefinitionUptimeTestsConfig() *APIDefinitionUptimeTestsConfig`

NewAPIDefinitionUptimeTestsConfig instantiates a new APIDefinitionUptimeTestsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIDefinitionUptimeTestsConfigWithDefaults

`func NewAPIDefinitionUptimeTestsConfigWithDefaults() *APIDefinitionUptimeTestsConfig`

NewAPIDefinitionUptimeTestsConfigWithDefaults instantiates a new APIDefinitionUptimeTestsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpireUtimeAfter

`func (o *APIDefinitionUptimeTestsConfig) GetExpireUtimeAfter() int64`

GetExpireUtimeAfter returns the ExpireUtimeAfter field if non-nil, zero value otherwise.

### GetExpireUtimeAfterOk

`func (o *APIDefinitionUptimeTestsConfig) GetExpireUtimeAfterOk() (*int64, bool)`

GetExpireUtimeAfterOk returns a tuple with the ExpireUtimeAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireUtimeAfter

`func (o *APIDefinitionUptimeTestsConfig) SetExpireUtimeAfter(v int64)`

SetExpireUtimeAfter sets ExpireUtimeAfter field to given value.

### HasExpireUtimeAfter

`func (o *APIDefinitionUptimeTestsConfig) HasExpireUtimeAfter() bool`

HasExpireUtimeAfter returns a boolean if a field has been set.

### GetRecheckWait

`func (o *APIDefinitionUptimeTestsConfig) GetRecheckWait() int64`

GetRecheckWait returns the RecheckWait field if non-nil, zero value otherwise.

### GetRecheckWaitOk

`func (o *APIDefinitionUptimeTestsConfig) GetRecheckWaitOk() (*int64, bool)`

GetRecheckWaitOk returns a tuple with the RecheckWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecheckWait

`func (o *APIDefinitionUptimeTestsConfig) SetRecheckWait(v int64)`

SetRecheckWait sets RecheckWait field to given value.

### HasRecheckWait

`func (o *APIDefinitionUptimeTestsConfig) HasRecheckWait() bool`

HasRecheckWait returns a boolean if a field has been set.

### GetServiceDiscovery

`func (o *APIDefinitionUptimeTestsConfig) GetServiceDiscovery() ServiceDiscoveryConfiguration`

GetServiceDiscovery returns the ServiceDiscovery field if non-nil, zero value otherwise.

### GetServiceDiscoveryOk

`func (o *APIDefinitionUptimeTestsConfig) GetServiceDiscoveryOk() (*ServiceDiscoveryConfiguration, bool)`

GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDiscovery

`func (o *APIDefinitionUptimeTestsConfig) SetServiceDiscovery(v ServiceDiscoveryConfiguration)`

SetServiceDiscovery sets ServiceDiscovery field to given value.

### HasServiceDiscovery

`func (o *APIDefinitionUptimeTestsConfig) HasServiceDiscovery() bool`

HasServiceDiscovery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


