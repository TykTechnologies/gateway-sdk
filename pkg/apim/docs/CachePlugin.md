# CachePlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheByRegex** | Pointer to **string** |  | [optional] 
**CacheResponseCodes** | Pointer to **[]int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 

## Methods

### NewCachePlugin

`func NewCachePlugin() *CachePlugin`

NewCachePlugin instantiates a new CachePlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCachePluginWithDefaults

`func NewCachePluginWithDefaults() *CachePlugin`

NewCachePluginWithDefaults instantiates a new CachePlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheByRegex

`func (o *CachePlugin) GetCacheByRegex() string`

GetCacheByRegex returns the CacheByRegex field if non-nil, zero value otherwise.

### GetCacheByRegexOk

`func (o *CachePlugin) GetCacheByRegexOk() (*string, bool)`

GetCacheByRegexOk returns a tuple with the CacheByRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheByRegex

`func (o *CachePlugin) SetCacheByRegex(v string)`

SetCacheByRegex sets CacheByRegex field to given value.

### HasCacheByRegex

`func (o *CachePlugin) HasCacheByRegex() bool`

HasCacheByRegex returns a boolean if a field has been set.

### GetCacheResponseCodes

`func (o *CachePlugin) GetCacheResponseCodes() []int32`

GetCacheResponseCodes returns the CacheResponseCodes field if non-nil, zero value otherwise.

### GetCacheResponseCodesOk

`func (o *CachePlugin) GetCacheResponseCodesOk() (*[]int32, bool)`

GetCacheResponseCodesOk returns a tuple with the CacheResponseCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheResponseCodes

`func (o *CachePlugin) SetCacheResponseCodes(v []int32)`

SetCacheResponseCodes sets CacheResponseCodes field to given value.

### HasCacheResponseCodes

`func (o *CachePlugin) HasCacheResponseCodes() bool`

HasCacheResponseCodes returns a boolean if a field has been set.

### GetEnabled

`func (o *CachePlugin) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CachePlugin) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CachePlugin) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CachePlugin) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTimeout

`func (o *CachePlugin) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CachePlugin) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CachePlugin) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CachePlugin) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


