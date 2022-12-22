# ServiceDiscoveryConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheTimeout** | Pointer to **int64** |  | [optional] 
**DataPath** | Pointer to **string** |  | [optional] 
**EndpointReturnsList** | Pointer to **bool** |  | [optional] 
**ParentDataPath** | Pointer to **string** |  | [optional] 
**PortDataPath** | Pointer to **string** |  | [optional] 
**QueryEndpoint** | Pointer to **string** |  | [optional] 
**TargetPath** | Pointer to **string** |  | [optional] 
**UseDiscoveryService** | Pointer to **bool** |  | [optional] 
**UseNestedQuery** | Pointer to **bool** |  | [optional] 
**UseTargetList** | Pointer to **bool** |  | [optional] 

## Methods

### NewServiceDiscoveryConfiguration

`func NewServiceDiscoveryConfiguration() *ServiceDiscoveryConfiguration`

NewServiceDiscoveryConfiguration instantiates a new ServiceDiscoveryConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDiscoveryConfigurationWithDefaults

`func NewServiceDiscoveryConfigurationWithDefaults() *ServiceDiscoveryConfiguration`

NewServiceDiscoveryConfigurationWithDefaults instantiates a new ServiceDiscoveryConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheTimeout

`func (o *ServiceDiscoveryConfiguration) GetCacheTimeout() int64`

GetCacheTimeout returns the CacheTimeout field if non-nil, zero value otherwise.

### GetCacheTimeoutOk

`func (o *ServiceDiscoveryConfiguration) GetCacheTimeoutOk() (*int64, bool)`

GetCacheTimeoutOk returns a tuple with the CacheTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTimeout

`func (o *ServiceDiscoveryConfiguration) SetCacheTimeout(v int64)`

SetCacheTimeout sets CacheTimeout field to given value.

### HasCacheTimeout

`func (o *ServiceDiscoveryConfiguration) HasCacheTimeout() bool`

HasCacheTimeout returns a boolean if a field has been set.

### GetDataPath

`func (o *ServiceDiscoveryConfiguration) GetDataPath() string`

GetDataPath returns the DataPath field if non-nil, zero value otherwise.

### GetDataPathOk

`func (o *ServiceDiscoveryConfiguration) GetDataPathOk() (*string, bool)`

GetDataPathOk returns a tuple with the DataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPath

`func (o *ServiceDiscoveryConfiguration) SetDataPath(v string)`

SetDataPath sets DataPath field to given value.

### HasDataPath

`func (o *ServiceDiscoveryConfiguration) HasDataPath() bool`

HasDataPath returns a boolean if a field has been set.

### GetEndpointReturnsList

`func (o *ServiceDiscoveryConfiguration) GetEndpointReturnsList() bool`

GetEndpointReturnsList returns the EndpointReturnsList field if non-nil, zero value otherwise.

### GetEndpointReturnsListOk

`func (o *ServiceDiscoveryConfiguration) GetEndpointReturnsListOk() (*bool, bool)`

GetEndpointReturnsListOk returns a tuple with the EndpointReturnsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointReturnsList

`func (o *ServiceDiscoveryConfiguration) SetEndpointReturnsList(v bool)`

SetEndpointReturnsList sets EndpointReturnsList field to given value.

### HasEndpointReturnsList

`func (o *ServiceDiscoveryConfiguration) HasEndpointReturnsList() bool`

HasEndpointReturnsList returns a boolean if a field has been set.

### GetParentDataPath

`func (o *ServiceDiscoveryConfiguration) GetParentDataPath() string`

GetParentDataPath returns the ParentDataPath field if non-nil, zero value otherwise.

### GetParentDataPathOk

`func (o *ServiceDiscoveryConfiguration) GetParentDataPathOk() (*string, bool)`

GetParentDataPathOk returns a tuple with the ParentDataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDataPath

`func (o *ServiceDiscoveryConfiguration) SetParentDataPath(v string)`

SetParentDataPath sets ParentDataPath field to given value.

### HasParentDataPath

`func (o *ServiceDiscoveryConfiguration) HasParentDataPath() bool`

HasParentDataPath returns a boolean if a field has been set.

### GetPortDataPath

`func (o *ServiceDiscoveryConfiguration) GetPortDataPath() string`

GetPortDataPath returns the PortDataPath field if non-nil, zero value otherwise.

### GetPortDataPathOk

`func (o *ServiceDiscoveryConfiguration) GetPortDataPathOk() (*string, bool)`

GetPortDataPathOk returns a tuple with the PortDataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDataPath

`func (o *ServiceDiscoveryConfiguration) SetPortDataPath(v string)`

SetPortDataPath sets PortDataPath field to given value.

### HasPortDataPath

`func (o *ServiceDiscoveryConfiguration) HasPortDataPath() bool`

HasPortDataPath returns a boolean if a field has been set.

### GetQueryEndpoint

`func (o *ServiceDiscoveryConfiguration) GetQueryEndpoint() string`

GetQueryEndpoint returns the QueryEndpoint field if non-nil, zero value otherwise.

### GetQueryEndpointOk

`func (o *ServiceDiscoveryConfiguration) GetQueryEndpointOk() (*string, bool)`

GetQueryEndpointOk returns a tuple with the QueryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryEndpoint

`func (o *ServiceDiscoveryConfiguration) SetQueryEndpoint(v string)`

SetQueryEndpoint sets QueryEndpoint field to given value.

### HasQueryEndpoint

`func (o *ServiceDiscoveryConfiguration) HasQueryEndpoint() bool`

HasQueryEndpoint returns a boolean if a field has been set.

### GetTargetPath

`func (o *ServiceDiscoveryConfiguration) GetTargetPath() string`

GetTargetPath returns the TargetPath field if non-nil, zero value otherwise.

### GetTargetPathOk

`func (o *ServiceDiscoveryConfiguration) GetTargetPathOk() (*string, bool)`

GetTargetPathOk returns a tuple with the TargetPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPath

`func (o *ServiceDiscoveryConfiguration) SetTargetPath(v string)`

SetTargetPath sets TargetPath field to given value.

### HasTargetPath

`func (o *ServiceDiscoveryConfiguration) HasTargetPath() bool`

HasTargetPath returns a boolean if a field has been set.

### GetUseDiscoveryService

`func (o *ServiceDiscoveryConfiguration) GetUseDiscoveryService() bool`

GetUseDiscoveryService returns the UseDiscoveryService field if non-nil, zero value otherwise.

### GetUseDiscoveryServiceOk

`func (o *ServiceDiscoveryConfiguration) GetUseDiscoveryServiceOk() (*bool, bool)`

GetUseDiscoveryServiceOk returns a tuple with the UseDiscoveryService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDiscoveryService

`func (o *ServiceDiscoveryConfiguration) SetUseDiscoveryService(v bool)`

SetUseDiscoveryService sets UseDiscoveryService field to given value.

### HasUseDiscoveryService

`func (o *ServiceDiscoveryConfiguration) HasUseDiscoveryService() bool`

HasUseDiscoveryService returns a boolean if a field has been set.

### GetUseNestedQuery

`func (o *ServiceDiscoveryConfiguration) GetUseNestedQuery() bool`

GetUseNestedQuery returns the UseNestedQuery field if non-nil, zero value otherwise.

### GetUseNestedQueryOk

`func (o *ServiceDiscoveryConfiguration) GetUseNestedQueryOk() (*bool, bool)`

GetUseNestedQueryOk returns a tuple with the UseNestedQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNestedQuery

`func (o *ServiceDiscoveryConfiguration) SetUseNestedQuery(v bool)`

SetUseNestedQuery sets UseNestedQuery field to given value.

### HasUseNestedQuery

`func (o *ServiceDiscoveryConfiguration) HasUseNestedQuery() bool`

HasUseNestedQuery returns a boolean if a field has been set.

### GetUseTargetList

`func (o *ServiceDiscoveryConfiguration) GetUseTargetList() bool`

GetUseTargetList returns the UseTargetList field if non-nil, zero value otherwise.

### GetUseTargetListOk

`func (o *ServiceDiscoveryConfiguration) GetUseTargetListOk() (*bool, bool)`

GetUseTargetListOk returns a tuple with the UseTargetList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTargetList

`func (o *ServiceDiscoveryConfiguration) SetUseTargetList(v bool)`

SetUseTargetList sets UseTargetList field to given value.

### HasUseTargetList

`func (o *ServiceDiscoveryConfiguration) HasUseTargetList() bool`

HasUseTargetList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


