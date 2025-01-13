# ProxyConfig

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
**Transport** | Pointer to [**ProxyConfigTransport**](ProxyConfigTransport.md) |  | [optional] 

## Methods

### NewProxyConfig

`func NewProxyConfig() *ProxyConfig`

NewProxyConfig instantiates a new ProxyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyConfigWithDefaults

`func NewProxyConfigWithDefaults() *ProxyConfig`

NewProxyConfigWithDefaults instantiates a new ProxyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckHostAgainstUptimeTests

`func (o *ProxyConfig) GetCheckHostAgainstUptimeTests() bool`

GetCheckHostAgainstUptimeTests returns the CheckHostAgainstUptimeTests field if non-nil, zero value otherwise.

### GetCheckHostAgainstUptimeTestsOk

`func (o *ProxyConfig) GetCheckHostAgainstUptimeTestsOk() (*bool, bool)`

GetCheckHostAgainstUptimeTestsOk returns a tuple with the CheckHostAgainstUptimeTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckHostAgainstUptimeTests

`func (o *ProxyConfig) SetCheckHostAgainstUptimeTests(v bool)`

SetCheckHostAgainstUptimeTests sets CheckHostAgainstUptimeTests field to given value.

### HasCheckHostAgainstUptimeTests

`func (o *ProxyConfig) HasCheckHostAgainstUptimeTests() bool`

HasCheckHostAgainstUptimeTests returns a boolean if a field has been set.

### GetDisableStripSlash

`func (o *ProxyConfig) GetDisableStripSlash() bool`

GetDisableStripSlash returns the DisableStripSlash field if non-nil, zero value otherwise.

### GetDisableStripSlashOk

`func (o *ProxyConfig) GetDisableStripSlashOk() (*bool, bool)`

GetDisableStripSlashOk returns a tuple with the DisableStripSlash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableStripSlash

`func (o *ProxyConfig) SetDisableStripSlash(v bool)`

SetDisableStripSlash sets DisableStripSlash field to given value.

### HasDisableStripSlash

`func (o *ProxyConfig) HasDisableStripSlash() bool`

HasDisableStripSlash returns a boolean if a field has been set.

### GetEnableLoadBalancing

`func (o *ProxyConfig) GetEnableLoadBalancing() bool`

GetEnableLoadBalancing returns the EnableLoadBalancing field if non-nil, zero value otherwise.

### GetEnableLoadBalancingOk

`func (o *ProxyConfig) GetEnableLoadBalancingOk() (*bool, bool)`

GetEnableLoadBalancingOk returns a tuple with the EnableLoadBalancing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLoadBalancing

`func (o *ProxyConfig) SetEnableLoadBalancing(v bool)`

SetEnableLoadBalancing sets EnableLoadBalancing field to given value.

### HasEnableLoadBalancing

`func (o *ProxyConfig) HasEnableLoadBalancing() bool`

HasEnableLoadBalancing returns a boolean if a field has been set.

### GetListenPath

`func (o *ProxyConfig) GetListenPath() string`

GetListenPath returns the ListenPath field if non-nil, zero value otherwise.

### GetListenPathOk

`func (o *ProxyConfig) GetListenPathOk() (*string, bool)`

GetListenPathOk returns a tuple with the ListenPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPath

`func (o *ProxyConfig) SetListenPath(v string)`

SetListenPath sets ListenPath field to given value.

### HasListenPath

`func (o *ProxyConfig) HasListenPath() bool`

HasListenPath returns a boolean if a field has been set.

### GetPreserveHostHeader

`func (o *ProxyConfig) GetPreserveHostHeader() bool`

GetPreserveHostHeader returns the PreserveHostHeader field if non-nil, zero value otherwise.

### GetPreserveHostHeaderOk

`func (o *ProxyConfig) GetPreserveHostHeaderOk() (*bool, bool)`

GetPreserveHostHeaderOk returns a tuple with the PreserveHostHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveHostHeader

`func (o *ProxyConfig) SetPreserveHostHeader(v bool)`

SetPreserveHostHeader sets PreserveHostHeader field to given value.

### HasPreserveHostHeader

`func (o *ProxyConfig) HasPreserveHostHeader() bool`

HasPreserveHostHeader returns a boolean if a field has been set.

### GetServiceDiscovery

`func (o *ProxyConfig) GetServiceDiscovery() ServiceDiscoveryConfiguration`

GetServiceDiscovery returns the ServiceDiscovery field if non-nil, zero value otherwise.

### GetServiceDiscoveryOk

`func (o *ProxyConfig) GetServiceDiscoveryOk() (*ServiceDiscoveryConfiguration, bool)`

GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDiscovery

`func (o *ProxyConfig) SetServiceDiscovery(v ServiceDiscoveryConfiguration)`

SetServiceDiscovery sets ServiceDiscovery field to given value.

### HasServiceDiscovery

`func (o *ProxyConfig) HasServiceDiscovery() bool`

HasServiceDiscovery returns a boolean if a field has been set.

### GetStripListenPath

`func (o *ProxyConfig) GetStripListenPath() bool`

GetStripListenPath returns the StripListenPath field if non-nil, zero value otherwise.

### GetStripListenPathOk

`func (o *ProxyConfig) GetStripListenPathOk() (*bool, bool)`

GetStripListenPathOk returns a tuple with the StripListenPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripListenPath

`func (o *ProxyConfig) SetStripListenPath(v bool)`

SetStripListenPath sets StripListenPath field to given value.

### HasStripListenPath

`func (o *ProxyConfig) HasStripListenPath() bool`

HasStripListenPath returns a boolean if a field has been set.

### GetTargetList

`func (o *ProxyConfig) GetTargetList() []string`

GetTargetList returns the TargetList field if non-nil, zero value otherwise.

### GetTargetListOk

`func (o *ProxyConfig) GetTargetListOk() (*[]string, bool)`

GetTargetListOk returns a tuple with the TargetList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetList

`func (o *ProxyConfig) SetTargetList(v []string)`

SetTargetList sets TargetList field to given value.

### HasTargetList

`func (o *ProxyConfig) HasTargetList() bool`

HasTargetList returns a boolean if a field has been set.

### SetTargetListNil

`func (o *ProxyConfig) SetTargetListNil(b bool)`

 SetTargetListNil sets the value for TargetList to be an explicit nil

### UnsetTargetList
`func (o *ProxyConfig) UnsetTargetList()`

UnsetTargetList ensures that no value is present for TargetList, not even an explicit nil
### GetTargetUrl

`func (o *ProxyConfig) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *ProxyConfig) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *ProxyConfig) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *ProxyConfig) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### GetTransport

`func (o *ProxyConfig) GetTransport() ProxyConfigTransport`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *ProxyConfig) GetTransportOk() (*ProxyConfigTransport, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *ProxyConfig) SetTransport(v ProxyConfigTransport)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *ProxyConfig) HasTransport() bool`

HasTransport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


