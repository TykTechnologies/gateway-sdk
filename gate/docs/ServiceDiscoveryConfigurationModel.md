# ServiceDiscoveryConfigurationModel

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

### NewServiceDiscoveryConfigurationModel

`func NewServiceDiscoveryConfigurationModel() *ServiceDiscoveryConfigurationModel`

NewServiceDiscoveryConfigurationModel instantiates a new ServiceDiscoveryConfigurationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDiscoveryConfigurationModelWithDefaults

`func NewServiceDiscoveryConfigurationModelWithDefaults() *ServiceDiscoveryConfigurationModel`

NewServiceDiscoveryConfigurationModelWithDefaults instantiates a new ServiceDiscoveryConfigurationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheTimeout

`func (o *ServiceDiscoveryConfigurationModel) GetCacheTimeout() int64`

GetCacheTimeout returns the CacheTimeout field if non-nil, zero value otherwise.

### GetCacheTimeoutOk

`func (o *ServiceDiscoveryConfigurationModel) GetCacheTimeoutOk() (*int64, bool)`

GetCacheTimeoutOk returns a tuple with the CacheTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTimeout

`func (o *ServiceDiscoveryConfigurationModel) SetCacheTimeout(v int64)`

SetCacheTimeout sets CacheTimeout field to given value.

### HasCacheTimeout

`func (o *ServiceDiscoveryConfigurationModel) HasCacheTimeout() bool`

HasCacheTimeout returns a boolean if a field has been set.

### GetDataPath

`func (o *ServiceDiscoveryConfigurationModel) GetDataPath() string`

GetDataPath returns the DataPath field if non-nil, zero value otherwise.

### GetDataPathOk

`func (o *ServiceDiscoveryConfigurationModel) GetDataPathOk() (*string, bool)`

GetDataPathOk returns a tuple with the DataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPath

`func (o *ServiceDiscoveryConfigurationModel) SetDataPath(v string)`

SetDataPath sets DataPath field to given value.

### HasDataPath

`func (o *ServiceDiscoveryConfigurationModel) HasDataPath() bool`

HasDataPath returns a boolean if a field has been set.

### GetEndpointReturnsList

`func (o *ServiceDiscoveryConfigurationModel) GetEndpointReturnsList() bool`

GetEndpointReturnsList returns the EndpointReturnsList field if non-nil, zero value otherwise.

### GetEndpointReturnsListOk

`func (o *ServiceDiscoveryConfigurationModel) GetEndpointReturnsListOk() (*bool, bool)`

GetEndpointReturnsListOk returns a tuple with the EndpointReturnsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointReturnsList

`func (o *ServiceDiscoveryConfigurationModel) SetEndpointReturnsList(v bool)`

SetEndpointReturnsList sets EndpointReturnsList field to given value.

### HasEndpointReturnsList

`func (o *ServiceDiscoveryConfigurationModel) HasEndpointReturnsList() bool`

HasEndpointReturnsList returns a boolean if a field has been set.

### GetParentDataPath

`func (o *ServiceDiscoveryConfigurationModel) GetParentDataPath() string`

GetParentDataPath returns the ParentDataPath field if non-nil, zero value otherwise.

### GetParentDataPathOk

`func (o *ServiceDiscoveryConfigurationModel) GetParentDataPathOk() (*string, bool)`

GetParentDataPathOk returns a tuple with the ParentDataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDataPath

`func (o *ServiceDiscoveryConfigurationModel) SetParentDataPath(v string)`

SetParentDataPath sets ParentDataPath field to given value.

### HasParentDataPath

`func (o *ServiceDiscoveryConfigurationModel) HasParentDataPath() bool`

HasParentDataPath returns a boolean if a field has been set.

### GetPortDataPath

`func (o *ServiceDiscoveryConfigurationModel) GetPortDataPath() string`

GetPortDataPath returns the PortDataPath field if non-nil, zero value otherwise.

### GetPortDataPathOk

`func (o *ServiceDiscoveryConfigurationModel) GetPortDataPathOk() (*string, bool)`

GetPortDataPathOk returns a tuple with the PortDataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDataPath

`func (o *ServiceDiscoveryConfigurationModel) SetPortDataPath(v string)`

SetPortDataPath sets PortDataPath field to given value.

### HasPortDataPath

`func (o *ServiceDiscoveryConfigurationModel) HasPortDataPath() bool`

HasPortDataPath returns a boolean if a field has been set.

### GetQueryEndpoint

`func (o *ServiceDiscoveryConfigurationModel) GetQueryEndpoint() string`

GetQueryEndpoint returns the QueryEndpoint field if non-nil, zero value otherwise.

### GetQueryEndpointOk

`func (o *ServiceDiscoveryConfigurationModel) GetQueryEndpointOk() (*string, bool)`

GetQueryEndpointOk returns a tuple with the QueryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryEndpoint

`func (o *ServiceDiscoveryConfigurationModel) SetQueryEndpoint(v string)`

SetQueryEndpoint sets QueryEndpoint field to given value.

### HasQueryEndpoint

`func (o *ServiceDiscoveryConfigurationModel) HasQueryEndpoint() bool`

HasQueryEndpoint returns a boolean if a field has been set.

### GetTargetPath

`func (o *ServiceDiscoveryConfigurationModel) GetTargetPath() string`

GetTargetPath returns the TargetPath field if non-nil, zero value otherwise.

### GetTargetPathOk

`func (o *ServiceDiscoveryConfigurationModel) GetTargetPathOk() (*string, bool)`

GetTargetPathOk returns a tuple with the TargetPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPath

`func (o *ServiceDiscoveryConfigurationModel) SetTargetPath(v string)`

SetTargetPath sets TargetPath field to given value.

### HasTargetPath

`func (o *ServiceDiscoveryConfigurationModel) HasTargetPath() bool`

HasTargetPath returns a boolean if a field has been set.

### GetUseDiscoveryService

`func (o *ServiceDiscoveryConfigurationModel) GetUseDiscoveryService() bool`

GetUseDiscoveryService returns the UseDiscoveryService field if non-nil, zero value otherwise.

### GetUseDiscoveryServiceOk

`func (o *ServiceDiscoveryConfigurationModel) GetUseDiscoveryServiceOk() (*bool, bool)`

GetUseDiscoveryServiceOk returns a tuple with the UseDiscoveryService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDiscoveryService

`func (o *ServiceDiscoveryConfigurationModel) SetUseDiscoveryService(v bool)`

SetUseDiscoveryService sets UseDiscoveryService field to given value.

### HasUseDiscoveryService

`func (o *ServiceDiscoveryConfigurationModel) HasUseDiscoveryService() bool`

HasUseDiscoveryService returns a boolean if a field has been set.

### GetUseNestedQuery

`func (o *ServiceDiscoveryConfigurationModel) GetUseNestedQuery() bool`

GetUseNestedQuery returns the UseNestedQuery field if non-nil, zero value otherwise.

### GetUseNestedQueryOk

`func (o *ServiceDiscoveryConfigurationModel) GetUseNestedQueryOk() (*bool, bool)`

GetUseNestedQueryOk returns a tuple with the UseNestedQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNestedQuery

`func (o *ServiceDiscoveryConfigurationModel) SetUseNestedQuery(v bool)`

SetUseNestedQuery sets UseNestedQuery field to given value.

### HasUseNestedQuery

`func (o *ServiceDiscoveryConfigurationModel) HasUseNestedQuery() bool`

HasUseNestedQuery returns a boolean if a field has been set.

### GetUseTargetList

`func (o *ServiceDiscoveryConfigurationModel) GetUseTargetList() bool`

GetUseTargetList returns the UseTargetList field if non-nil, zero value otherwise.

### GetUseTargetListOk

`func (o *ServiceDiscoveryConfigurationModel) GetUseTargetListOk() (*bool, bool)`

GetUseTargetListOk returns a tuple with the UseTargetList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTargetList

`func (o *ServiceDiscoveryConfigurationModel) SetUseTargetList(v bool)`

SetUseTargetList sets UseTargetList field to given value.

### HasUseTargetList

`func (o *ServiceDiscoveryConfigurationModel) HasUseTargetList() bool`

HasUseTargetList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


