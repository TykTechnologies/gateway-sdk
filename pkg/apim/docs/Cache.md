# Cache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheAllSafeRequests** | Pointer to **bool** |  | [optional] 
**CacheByHeaders** | Pointer to **[]string** |  | [optional] 
**CacheResponseCodes** | Pointer to **[]int32** |  | [optional] 
**ControlTTLHeaderName** | Pointer to **string** |  | [optional] 
**EnableUpstreamCacheControl** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 

## Methods

### NewCache

`func NewCache() *Cache`

NewCache instantiates a new Cache object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheWithDefaults

`func NewCacheWithDefaults() *Cache`

NewCacheWithDefaults instantiates a new Cache object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheAllSafeRequests

`func (o *Cache) GetCacheAllSafeRequests() bool`

GetCacheAllSafeRequests returns the CacheAllSafeRequests field if non-nil, zero value otherwise.

### GetCacheAllSafeRequestsOk

`func (o *Cache) GetCacheAllSafeRequestsOk() (*bool, bool)`

GetCacheAllSafeRequestsOk returns a tuple with the CacheAllSafeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheAllSafeRequests

`func (o *Cache) SetCacheAllSafeRequests(v bool)`

SetCacheAllSafeRequests sets CacheAllSafeRequests field to given value.

### HasCacheAllSafeRequests

`func (o *Cache) HasCacheAllSafeRequests() bool`

HasCacheAllSafeRequests returns a boolean if a field has been set.

### GetCacheByHeaders

`func (o *Cache) GetCacheByHeaders() []string`

GetCacheByHeaders returns the CacheByHeaders field if non-nil, zero value otherwise.

### GetCacheByHeadersOk

`func (o *Cache) GetCacheByHeadersOk() (*[]string, bool)`

GetCacheByHeadersOk returns a tuple with the CacheByHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheByHeaders

`func (o *Cache) SetCacheByHeaders(v []string)`

SetCacheByHeaders sets CacheByHeaders field to given value.

### HasCacheByHeaders

`func (o *Cache) HasCacheByHeaders() bool`

HasCacheByHeaders returns a boolean if a field has been set.

### GetCacheResponseCodes

`func (o *Cache) GetCacheResponseCodes() []int32`

GetCacheResponseCodes returns the CacheResponseCodes field if non-nil, zero value otherwise.

### GetCacheResponseCodesOk

`func (o *Cache) GetCacheResponseCodesOk() (*[]int32, bool)`

GetCacheResponseCodesOk returns a tuple with the CacheResponseCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheResponseCodes

`func (o *Cache) SetCacheResponseCodes(v []int32)`

SetCacheResponseCodes sets CacheResponseCodes field to given value.

### HasCacheResponseCodes

`func (o *Cache) HasCacheResponseCodes() bool`

HasCacheResponseCodes returns a boolean if a field has been set.

### GetControlTTLHeaderName

`func (o *Cache) GetControlTTLHeaderName() string`

GetControlTTLHeaderName returns the ControlTTLHeaderName field if non-nil, zero value otherwise.

### GetControlTTLHeaderNameOk

`func (o *Cache) GetControlTTLHeaderNameOk() (*string, bool)`

GetControlTTLHeaderNameOk returns a tuple with the ControlTTLHeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlTTLHeaderName

`func (o *Cache) SetControlTTLHeaderName(v string)`

SetControlTTLHeaderName sets ControlTTLHeaderName field to given value.

### HasControlTTLHeaderName

`func (o *Cache) HasControlTTLHeaderName() bool`

HasControlTTLHeaderName returns a boolean if a field has been set.

### GetEnableUpstreamCacheControl

`func (o *Cache) GetEnableUpstreamCacheControl() bool`

GetEnableUpstreamCacheControl returns the EnableUpstreamCacheControl field if non-nil, zero value otherwise.

### GetEnableUpstreamCacheControlOk

`func (o *Cache) GetEnableUpstreamCacheControlOk() (*bool, bool)`

GetEnableUpstreamCacheControlOk returns a tuple with the EnableUpstreamCacheControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUpstreamCacheControl

`func (o *Cache) SetEnableUpstreamCacheControl(v bool)`

SetEnableUpstreamCacheControl sets EnableUpstreamCacheControl field to given value.

### HasEnableUpstreamCacheControl

`func (o *Cache) HasEnableUpstreamCacheControl() bool`

HasEnableUpstreamCacheControl returns a boolean if a field has been set.

### GetEnabled

`func (o *Cache) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Cache) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Cache) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Cache) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTimeout

`func (o *Cache) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *Cache) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *Cache) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *Cache) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


