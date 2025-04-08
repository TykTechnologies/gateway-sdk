# ServiceDiscovery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to [**ServiceDiscoveryCache**](ServiceDiscoveryCache.md) |  | [optional] 
**CacheTimeout** | Pointer to **int32** |  | [optional] 
**DataPath** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**EndpointReturnsList** | Pointer to **bool** |  | [optional] 
**ParentDataPath** | Pointer to **string** |  | [optional] 
**PortDataPath** | Pointer to **string** |  | [optional] 
**QueryEndpoint** | Pointer to **string** |  | [optional] 
**TargetPath** | Pointer to **string** |  | [optional] 
**UseNestedQuery** | Pointer to **bool** |  | [optional] 
**UseTargetList** | Pointer to **bool** |  | [optional] 

## Methods

### NewServiceDiscovery

`func NewServiceDiscovery() *ServiceDiscovery`

NewServiceDiscovery instantiates a new ServiceDiscovery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDiscoveryWithDefaults

`func NewServiceDiscoveryWithDefaults() *ServiceDiscovery`

NewServiceDiscoveryWithDefaults instantiates a new ServiceDiscovery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *ServiceDiscovery) GetCache() ServiceDiscoveryCache`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *ServiceDiscovery) GetCacheOk() (*ServiceDiscoveryCache, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *ServiceDiscovery) SetCache(v ServiceDiscoveryCache)`

SetCache sets Cache field to given value.

### HasCache

`func (o *ServiceDiscovery) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetCacheTimeout

`func (o *ServiceDiscovery) GetCacheTimeout() int32`

GetCacheTimeout returns the CacheTimeout field if non-nil, zero value otherwise.

### GetCacheTimeoutOk

`func (o *ServiceDiscovery) GetCacheTimeoutOk() (*int32, bool)`

GetCacheTimeoutOk returns a tuple with the CacheTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTimeout

`func (o *ServiceDiscovery) SetCacheTimeout(v int32)`

SetCacheTimeout sets CacheTimeout field to given value.

### HasCacheTimeout

`func (o *ServiceDiscovery) HasCacheTimeout() bool`

HasCacheTimeout returns a boolean if a field has been set.

### GetDataPath

`func (o *ServiceDiscovery) GetDataPath() string`

GetDataPath returns the DataPath field if non-nil, zero value otherwise.

### GetDataPathOk

`func (o *ServiceDiscovery) GetDataPathOk() (*string, bool)`

GetDataPathOk returns a tuple with the DataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPath

`func (o *ServiceDiscovery) SetDataPath(v string)`

SetDataPath sets DataPath field to given value.

### HasDataPath

`func (o *ServiceDiscovery) HasDataPath() bool`

HasDataPath returns a boolean if a field has been set.

### GetEnabled

`func (o *ServiceDiscovery) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServiceDiscovery) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServiceDiscovery) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ServiceDiscovery) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEndpointReturnsList

`func (o *ServiceDiscovery) GetEndpointReturnsList() bool`

GetEndpointReturnsList returns the EndpointReturnsList field if non-nil, zero value otherwise.

### GetEndpointReturnsListOk

`func (o *ServiceDiscovery) GetEndpointReturnsListOk() (*bool, bool)`

GetEndpointReturnsListOk returns a tuple with the EndpointReturnsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointReturnsList

`func (o *ServiceDiscovery) SetEndpointReturnsList(v bool)`

SetEndpointReturnsList sets EndpointReturnsList field to given value.

### HasEndpointReturnsList

`func (o *ServiceDiscovery) HasEndpointReturnsList() bool`

HasEndpointReturnsList returns a boolean if a field has been set.

### GetParentDataPath

`func (o *ServiceDiscovery) GetParentDataPath() string`

GetParentDataPath returns the ParentDataPath field if non-nil, zero value otherwise.

### GetParentDataPathOk

`func (o *ServiceDiscovery) GetParentDataPathOk() (*string, bool)`

GetParentDataPathOk returns a tuple with the ParentDataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDataPath

`func (o *ServiceDiscovery) SetParentDataPath(v string)`

SetParentDataPath sets ParentDataPath field to given value.

### HasParentDataPath

`func (o *ServiceDiscovery) HasParentDataPath() bool`

HasParentDataPath returns a boolean if a field has been set.

### GetPortDataPath

`func (o *ServiceDiscovery) GetPortDataPath() string`

GetPortDataPath returns the PortDataPath field if non-nil, zero value otherwise.

### GetPortDataPathOk

`func (o *ServiceDiscovery) GetPortDataPathOk() (*string, bool)`

GetPortDataPathOk returns a tuple with the PortDataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDataPath

`func (o *ServiceDiscovery) SetPortDataPath(v string)`

SetPortDataPath sets PortDataPath field to given value.

### HasPortDataPath

`func (o *ServiceDiscovery) HasPortDataPath() bool`

HasPortDataPath returns a boolean if a field has been set.

### GetQueryEndpoint

`func (o *ServiceDiscovery) GetQueryEndpoint() string`

GetQueryEndpoint returns the QueryEndpoint field if non-nil, zero value otherwise.

### GetQueryEndpointOk

`func (o *ServiceDiscovery) GetQueryEndpointOk() (*string, bool)`

GetQueryEndpointOk returns a tuple with the QueryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryEndpoint

`func (o *ServiceDiscovery) SetQueryEndpoint(v string)`

SetQueryEndpoint sets QueryEndpoint field to given value.

### HasQueryEndpoint

`func (o *ServiceDiscovery) HasQueryEndpoint() bool`

HasQueryEndpoint returns a boolean if a field has been set.

### GetTargetPath

`func (o *ServiceDiscovery) GetTargetPath() string`

GetTargetPath returns the TargetPath field if non-nil, zero value otherwise.

### GetTargetPathOk

`func (o *ServiceDiscovery) GetTargetPathOk() (*string, bool)`

GetTargetPathOk returns a tuple with the TargetPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPath

`func (o *ServiceDiscovery) SetTargetPath(v string)`

SetTargetPath sets TargetPath field to given value.

### HasTargetPath

`func (o *ServiceDiscovery) HasTargetPath() bool`

HasTargetPath returns a boolean if a field has been set.

### GetUseNestedQuery

`func (o *ServiceDiscovery) GetUseNestedQuery() bool`

GetUseNestedQuery returns the UseNestedQuery field if non-nil, zero value otherwise.

### GetUseNestedQueryOk

`func (o *ServiceDiscovery) GetUseNestedQueryOk() (*bool, bool)`

GetUseNestedQueryOk returns a tuple with the UseNestedQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNestedQuery

`func (o *ServiceDiscovery) SetUseNestedQuery(v bool)`

SetUseNestedQuery sets UseNestedQuery field to given value.

### HasUseNestedQuery

`func (o *ServiceDiscovery) HasUseNestedQuery() bool`

HasUseNestedQuery returns a boolean if a field has been set.

### GetUseTargetList

`func (o *ServiceDiscovery) GetUseTargetList() bool`

GetUseTargetList returns the UseTargetList field if non-nil, zero value otherwise.

### GetUseTargetListOk

`func (o *ServiceDiscovery) GetUseTargetListOk() (*bool, bool)`

GetUseTargetListOk returns a tuple with the UseTargetList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTargetList

`func (o *ServiceDiscovery) SetUseTargetList(v bool)`

SetUseTargetList sets UseTargetList field to given value.

### HasUseTargetList

`func (o *ServiceDiscovery) HasUseTargetList() bool`

HasUseTargetList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


