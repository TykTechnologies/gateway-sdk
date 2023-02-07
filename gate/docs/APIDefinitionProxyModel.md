# APIDefinitionProxyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckHostAgainstUptimeTests** | Pointer to **bool** |  | [optional] 
**DisableStripSlash** | Pointer to **bool** |  | [optional] 
**EnableLoadBalancing** | Pointer to **bool** |  | [optional] 
**ListenPath** | Pointer to **string** |  | [optional] 
**PreserveHostHeader** | Pointer to **bool** |  | [optional] 
**ServiceDiscovery** | Pointer to [**ServiceDiscoveryConfigurationModelModel**](ServiceDiscoveryConfigurationModel.md) |  | [optional] 
**StripListenPath** | Pointer to **bool** |  | [optional] 
**TargetList** | Pointer to **[]string** |  | [optional] 
**TargetUrl** | Pointer to **string** |  | [optional] 
**Transport** | Pointer to [**APIDefinitionProxyTransportModelModel**](APIDefinitionProxyTransportModel.md) |  | [optional] 

## Methods

### NewAPIDefinitionProxyModel

`func NewAPIDefinitionProxyModel() *APIDefinitionProxyModel`

NewAPIDefinitionProxyModel instantiates a new APIDefinitionProxyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIDefinitionProxyModelWithDefaults

`func NewAPIDefinitionProxyModelWithDefaults() *APIDefinitionProxyModel`

NewAPIDefinitionProxyModelWithDefaults instantiates a new APIDefinitionProxyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckHostAgainstUptimeTests

`func (o *APIDefinitionProxyModel) GetCheckHostAgainstUptimeTests() bool`

GetCheckHostAgainstUptimeTests returns the CheckHostAgainstUptimeTests field if non-nil, zero value otherwise.

### GetCheckHostAgainstUptimeTestsOk

`func (o *APIDefinitionProxyModel) GetCheckHostAgainstUptimeTestsOk() (*bool, bool)`

GetCheckHostAgainstUptimeTestsOk returns a tuple with the CheckHostAgainstUptimeTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckHostAgainstUptimeTests

`func (o *APIDefinitionProxyModel) SetCheckHostAgainstUptimeTests(v bool)`

SetCheckHostAgainstUptimeTests sets CheckHostAgainstUptimeTests field to given value.

### HasCheckHostAgainstUptimeTests

`func (o *APIDefinitionProxyModel) HasCheckHostAgainstUptimeTests() bool`

HasCheckHostAgainstUptimeTests returns a boolean if a field has been set.

### GetDisableStripSlash

`func (o *APIDefinitionProxyModel) GetDisableStripSlash() bool`

GetDisableStripSlash returns the DisableStripSlash field if non-nil, zero value otherwise.

### GetDisableStripSlashOk

`func (o *APIDefinitionProxyModel) GetDisableStripSlashOk() (*bool, bool)`

GetDisableStripSlashOk returns a tuple with the DisableStripSlash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableStripSlash

`func (o *APIDefinitionProxyModel) SetDisableStripSlash(v bool)`

SetDisableStripSlash sets DisableStripSlash field to given value.

### HasDisableStripSlash

`func (o *APIDefinitionProxyModel) HasDisableStripSlash() bool`

HasDisableStripSlash returns a boolean if a field has been set.

### GetEnableLoadBalancing

`func (o *APIDefinitionProxyModel) GetEnableLoadBalancing() bool`

GetEnableLoadBalancing returns the EnableLoadBalancing field if non-nil, zero value otherwise.

### GetEnableLoadBalancingOk

`func (o *APIDefinitionProxyModel) GetEnableLoadBalancingOk() (*bool, bool)`

GetEnableLoadBalancingOk returns a tuple with the EnableLoadBalancing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLoadBalancing

`func (o *APIDefinitionProxyModel) SetEnableLoadBalancing(v bool)`

SetEnableLoadBalancing sets EnableLoadBalancing field to given value.

### HasEnableLoadBalancing

`func (o *APIDefinitionProxyModel) HasEnableLoadBalancing() bool`

HasEnableLoadBalancing returns a boolean if a field has been set.

### GetListenPath

`func (o *APIDefinitionProxyModel) GetListenPath() string`

GetListenPath returns the ListenPath field if non-nil, zero value otherwise.

### GetListenPathOk

`func (o *APIDefinitionProxyModel) GetListenPathOk() (*string, bool)`

GetListenPathOk returns a tuple with the ListenPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPath

`func (o *APIDefinitionProxyModel) SetListenPath(v string)`

SetListenPath sets ListenPath field to given value.

### HasListenPath

`func (o *APIDefinitionProxyModel) HasListenPath() bool`

HasListenPath returns a boolean if a field has been set.

### GetPreserveHostHeader

`func (o *APIDefinitionProxyModel) GetPreserveHostHeader() bool`

GetPreserveHostHeader returns the PreserveHostHeader field if non-nil, zero value otherwise.

### GetPreserveHostHeaderOk

`func (o *APIDefinitionProxyModel) GetPreserveHostHeaderOk() (*bool, bool)`

GetPreserveHostHeaderOk returns a tuple with the PreserveHostHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveHostHeader

`func (o *APIDefinitionProxyModel) SetPreserveHostHeader(v bool)`

SetPreserveHostHeader sets PreserveHostHeader field to given value.

### HasPreserveHostHeader

`func (o *APIDefinitionProxyModel) HasPreserveHostHeader() bool`

HasPreserveHostHeader returns a boolean if a field has been set.

### GetServiceDiscovery

`func (o *APIDefinitionProxyModel) GetServiceDiscovery() ServiceDiscoveryConfigurationModelModel`

GetServiceDiscovery returns the ServiceDiscovery field if non-nil, zero value otherwise.

### GetServiceDiscoveryOk

`func (o *APIDefinitionProxyModel) GetServiceDiscoveryOk() (*ServiceDiscoveryConfigurationModelModel, bool)`

GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDiscovery

`func (o *APIDefinitionProxyModel) SetServiceDiscovery(v ServiceDiscoveryConfigurationModelModel)`

SetServiceDiscovery sets ServiceDiscovery field to given value.

### HasServiceDiscovery

`func (o *APIDefinitionProxyModel) HasServiceDiscovery() bool`

HasServiceDiscovery returns a boolean if a field has been set.

### GetStripListenPath

`func (o *APIDefinitionProxyModel) GetStripListenPath() bool`

GetStripListenPath returns the StripListenPath field if non-nil, zero value otherwise.

### GetStripListenPathOk

`func (o *APIDefinitionProxyModel) GetStripListenPathOk() (*bool, bool)`

GetStripListenPathOk returns a tuple with the StripListenPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripListenPath

`func (o *APIDefinitionProxyModel) SetStripListenPath(v bool)`

SetStripListenPath sets StripListenPath field to given value.

### HasStripListenPath

`func (o *APIDefinitionProxyModel) HasStripListenPath() bool`

HasStripListenPath returns a boolean if a field has been set.

### GetTargetList

`func (o *APIDefinitionProxyModel) GetTargetList() []string`

GetTargetList returns the TargetList field if non-nil, zero value otherwise.

### GetTargetListOk

`func (o *APIDefinitionProxyModel) GetTargetListOk() (*[]string, bool)`

GetTargetListOk returns a tuple with the TargetList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetList

`func (o *APIDefinitionProxyModel) SetTargetList(v []string)`

SetTargetList sets TargetList field to given value.

### HasTargetList

`func (o *APIDefinitionProxyModel) HasTargetList() bool`

HasTargetList returns a boolean if a field has been set.

### GetTargetUrl

`func (o *APIDefinitionProxyModel) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *APIDefinitionProxyModel) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *APIDefinitionProxyModel) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *APIDefinitionProxyModel) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### GetTransport

`func (o *APIDefinitionProxyModel) GetTransport() APIDefinitionProxyTransportModelModel`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *APIDefinitionProxyModel) GetTransportOk() (*APIDefinitionProxyTransportModelModel, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *APIDefinitionProxyModel) SetTransport(v APIDefinitionProxyTransportModelModel)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *APIDefinitionProxyModel) HasTransport() bool`

HasTransport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


