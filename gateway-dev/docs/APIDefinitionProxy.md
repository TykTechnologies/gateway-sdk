# APIDefinitionProxy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckHostAgainstUptimeTests** | Pointer to **bool** |  | [optional] 
**DisableStripSlash** | Pointer to **bool** |  | [optional] 
**EnableLoadBalancing** | Pointer to **bool** |  | [optional] 
**ListenPath** | Pointer to **string** |  | [optional] 
**PreserveHostHeader** | Pointer to **bool** |  | [optional] 
**ServiceDiscovery** | Pointer to [**ServiceDiscoveryConfiguration**](ServiceDiscoveryConfiguration.md) |  | [optional] 
**StripListenPath** | Pointer to **bool** |  | [optional] 
**TargetList** | Pointer to **[]string** |  | [optional] 
**TargetUrl** | Pointer to **string** |  | [optional] 
**Transport** | Pointer to [**APIDefinitionProxyTransport**](APIDefinitionProxyTransport.md) |  | [optional] 

## Methods

### NewAPIDefinitionProxy

`func NewAPIDefinitionProxy() *APIDefinitionProxy`

NewAPIDefinitionProxy instantiates a new APIDefinitionProxy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIDefinitionProxyWithDefaults

`func NewAPIDefinitionProxyWithDefaults() *APIDefinitionProxy`

NewAPIDefinitionProxyWithDefaults instantiates a new APIDefinitionProxy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckHostAgainstUptimeTests

`func (o *APIDefinitionProxy) GetCheckHostAgainstUptimeTests() bool`

GetCheckHostAgainstUptimeTests returns the CheckHostAgainstUptimeTests field if non-nil, zero value otherwise.

### GetCheckHostAgainstUptimeTestsOk

`func (o *APIDefinitionProxy) GetCheckHostAgainstUptimeTestsOk() (*bool, bool)`

GetCheckHostAgainstUptimeTestsOk returns a tuple with the CheckHostAgainstUptimeTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckHostAgainstUptimeTests

`func (o *APIDefinitionProxy) SetCheckHostAgainstUptimeTests(v bool)`

SetCheckHostAgainstUptimeTests sets CheckHostAgainstUptimeTests field to given value.

### HasCheckHostAgainstUptimeTests

`func (o *APIDefinitionProxy) HasCheckHostAgainstUptimeTests() bool`

HasCheckHostAgainstUptimeTests returns a boolean if a field has been set.

### GetDisableStripSlash

`func (o *APIDefinitionProxy) GetDisableStripSlash() bool`

GetDisableStripSlash returns the DisableStripSlash field if non-nil, zero value otherwise.

### GetDisableStripSlashOk

`func (o *APIDefinitionProxy) GetDisableStripSlashOk() (*bool, bool)`

GetDisableStripSlashOk returns a tuple with the DisableStripSlash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableStripSlash

`func (o *APIDefinitionProxy) SetDisableStripSlash(v bool)`

SetDisableStripSlash sets DisableStripSlash field to given value.

### HasDisableStripSlash

`func (o *APIDefinitionProxy) HasDisableStripSlash() bool`

HasDisableStripSlash returns a boolean if a field has been set.

### GetEnableLoadBalancing

`func (o *APIDefinitionProxy) GetEnableLoadBalancing() bool`

GetEnableLoadBalancing returns the EnableLoadBalancing field if non-nil, zero value otherwise.

### GetEnableLoadBalancingOk

`func (o *APIDefinitionProxy) GetEnableLoadBalancingOk() (*bool, bool)`

GetEnableLoadBalancingOk returns a tuple with the EnableLoadBalancing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLoadBalancing

`func (o *APIDefinitionProxy) SetEnableLoadBalancing(v bool)`

SetEnableLoadBalancing sets EnableLoadBalancing field to given value.

### HasEnableLoadBalancing

`func (o *APIDefinitionProxy) HasEnableLoadBalancing() bool`

HasEnableLoadBalancing returns a boolean if a field has been set.

### GetListenPath

`func (o *APIDefinitionProxy) GetListenPath() string`

GetListenPath returns the ListenPath field if non-nil, zero value otherwise.

### GetListenPathOk

`func (o *APIDefinitionProxy) GetListenPathOk() (*string, bool)`

GetListenPathOk returns a tuple with the ListenPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPath

`func (o *APIDefinitionProxy) SetListenPath(v string)`

SetListenPath sets ListenPath field to given value.

### HasListenPath

`func (o *APIDefinitionProxy) HasListenPath() bool`

HasListenPath returns a boolean if a field has been set.

### GetPreserveHostHeader

`func (o *APIDefinitionProxy) GetPreserveHostHeader() bool`

GetPreserveHostHeader returns the PreserveHostHeader field if non-nil, zero value otherwise.

### GetPreserveHostHeaderOk

`func (o *APIDefinitionProxy) GetPreserveHostHeaderOk() (*bool, bool)`

GetPreserveHostHeaderOk returns a tuple with the PreserveHostHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveHostHeader

`func (o *APIDefinitionProxy) SetPreserveHostHeader(v bool)`

SetPreserveHostHeader sets PreserveHostHeader field to given value.

### HasPreserveHostHeader

`func (o *APIDefinitionProxy) HasPreserveHostHeader() bool`

HasPreserveHostHeader returns a boolean if a field has been set.

### GetServiceDiscovery

`func (o *APIDefinitionProxy) GetServiceDiscovery() ServiceDiscoveryConfiguration`

GetServiceDiscovery returns the ServiceDiscovery field if non-nil, zero value otherwise.

### GetServiceDiscoveryOk

`func (o *APIDefinitionProxy) GetServiceDiscoveryOk() (*ServiceDiscoveryConfiguration, bool)`

GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDiscovery

`func (o *APIDefinitionProxy) SetServiceDiscovery(v ServiceDiscoveryConfiguration)`

SetServiceDiscovery sets ServiceDiscovery field to given value.

### HasServiceDiscovery

`func (o *APIDefinitionProxy) HasServiceDiscovery() bool`

HasServiceDiscovery returns a boolean if a field has been set.

### GetStripListenPath

`func (o *APIDefinitionProxy) GetStripListenPath() bool`

GetStripListenPath returns the StripListenPath field if non-nil, zero value otherwise.

### GetStripListenPathOk

`func (o *APIDefinitionProxy) GetStripListenPathOk() (*bool, bool)`

GetStripListenPathOk returns a tuple with the StripListenPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripListenPath

`func (o *APIDefinitionProxy) SetStripListenPath(v bool)`

SetStripListenPath sets StripListenPath field to given value.

### HasStripListenPath

`func (o *APIDefinitionProxy) HasStripListenPath() bool`

HasStripListenPath returns a boolean if a field has been set.

### GetTargetList

`func (o *APIDefinitionProxy) GetTargetList() []string`

GetTargetList returns the TargetList field if non-nil, zero value otherwise.

### GetTargetListOk

`func (o *APIDefinitionProxy) GetTargetListOk() (*[]string, bool)`

GetTargetListOk returns a tuple with the TargetList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetList

`func (o *APIDefinitionProxy) SetTargetList(v []string)`

SetTargetList sets TargetList field to given value.

### HasTargetList

`func (o *APIDefinitionProxy) HasTargetList() bool`

HasTargetList returns a boolean if a field has been set.

### GetTargetUrl

`func (o *APIDefinitionProxy) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *APIDefinitionProxy) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *APIDefinitionProxy) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *APIDefinitionProxy) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### GetTransport

`func (o *APIDefinitionProxy) GetTransport() APIDefinitionProxyTransport`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *APIDefinitionProxy) GetTransportOk() (*APIDefinitionProxyTransport, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *APIDefinitionProxy) SetTransport(v APIDefinitionProxyTransport)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *APIDefinitionProxy) HasTransport() bool`

HasTransport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


