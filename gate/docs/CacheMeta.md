# CacheMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheResponseCodes** | Pointer to **[]int64** |  | [optional] 
**CacheKeyRegex** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 

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

### GetCacheResponseCodes

`func (o *CacheMeta) GetCacheResponseCodes() []int64`

GetCacheResponseCodes returns the CacheResponseCodes field if non-nil, zero value otherwise.

### GetCacheResponseCodesOk

`func (o *CacheMeta) GetCacheResponseCodesOk() (*[]int64, bool)`

GetCacheResponseCodesOk returns a tuple with the CacheResponseCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheResponseCodes

`func (o *CacheMeta) SetCacheResponseCodes(v []int64)`

SetCacheResponseCodes sets CacheResponseCodes field to given value.

### HasCacheResponseCodes

`func (o *CacheMeta) HasCacheResponseCodes() bool`

HasCacheResponseCodes returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


