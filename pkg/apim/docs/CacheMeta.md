# CacheMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheKeyRegex** | Pointer to **string** |  | [optional] 
**CacheResponseCodes** | Pointer to **[]int32** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int64** |  | [optional] 

## Methods

### NewCacheMeta

`func NewCacheMeta() *CacheMeta`

NewCacheMeta instantiates a new CacheMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheMetaWithDefaults

`func NewCacheMetaWithDefaults() *CacheMeta`

NewCacheMetaWithDefaults instantiates a new CacheMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheKeyRegex

`func (o *CacheMeta) GetCacheKeyRegex() string`

GetCacheKeyRegex returns the CacheKeyRegex field if non-nil, zero value otherwise.

### GetCacheKeyRegexOk

`func (o *CacheMeta) GetCacheKeyRegexOk() (*string, bool)`

GetCacheKeyRegexOk returns a tuple with the CacheKeyRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheKeyRegex

`func (o *CacheMeta) SetCacheKeyRegex(v string)`

SetCacheKeyRegex sets CacheKeyRegex field to given value.

### HasCacheKeyRegex

`func (o *CacheMeta) HasCacheKeyRegex() bool`

HasCacheKeyRegex returns a boolean if a field has been set.

### GetCacheResponseCodes

`func (o *CacheMeta) GetCacheResponseCodes() []int32`

GetCacheResponseCodes returns the CacheResponseCodes field if non-nil, zero value otherwise.

### GetCacheResponseCodesOk

`func (o *CacheMeta) GetCacheResponseCodesOk() (*[]int32, bool)`

GetCacheResponseCodesOk returns a tuple with the CacheResponseCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheResponseCodes

`func (o *CacheMeta) SetCacheResponseCodes(v []int32)`

SetCacheResponseCodes sets CacheResponseCodes field to given value.

### HasCacheResponseCodes

`func (o *CacheMeta) HasCacheResponseCodes() bool`

HasCacheResponseCodes returns a boolean if a field has been set.

### SetCacheResponseCodesNil

`func (o *CacheMeta) SetCacheResponseCodesNil(b bool)`

 SetCacheResponseCodesNil sets the value for CacheResponseCodes to be an explicit nil

### UnsetCacheResponseCodes
`func (o *CacheMeta) UnsetCacheResponseCodes()`

UnsetCacheResponseCodes ensures that no value is present for CacheResponseCodes, not even an explicit nil
### GetDisabled

`func (o *CacheMeta) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *CacheMeta) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *CacheMeta) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *CacheMeta) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetMethod

`func (o *CacheMeta) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CacheMeta) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CacheMeta) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *CacheMeta) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPath

`func (o *CacheMeta) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CacheMeta) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CacheMeta) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CacheMeta) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetTimeout

`func (o *CacheMeta) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CacheMeta) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CacheMeta) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CacheMeta) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


