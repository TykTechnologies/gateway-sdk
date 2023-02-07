# CacheMetaModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheResponseCodes** | Pointer to **[]int64** |  | [optional] 
**CacheKeyRegex** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewCacheMetaModel

`func NewCacheMetaModel() *CacheMetaModel`

NewCacheMetaModel instantiates a new CacheMetaModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheMetaModelWithDefaults

`func NewCacheMetaModelWithDefaults() *CacheMetaModel`

NewCacheMetaModelWithDefaults instantiates a new CacheMetaModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheResponseCodes

`func (o *CacheMetaModel) GetCacheResponseCodes() []int64`

GetCacheResponseCodes returns the CacheResponseCodes field if non-nil, zero value otherwise.

### GetCacheResponseCodesOk

`func (o *CacheMetaModel) GetCacheResponseCodesOk() (*[]int64, bool)`

GetCacheResponseCodesOk returns a tuple with the CacheResponseCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheResponseCodes

`func (o *CacheMetaModel) SetCacheResponseCodes(v []int64)`

SetCacheResponseCodes sets CacheResponseCodes field to given value.

### HasCacheResponseCodes

`func (o *CacheMetaModel) HasCacheResponseCodes() bool`

HasCacheResponseCodes returns a boolean if a field has been set.

### GetCacheKeyRegex

`func (o *CacheMetaModel) GetCacheKeyRegex() string`

GetCacheKeyRegex returns the CacheKeyRegex field if non-nil, zero value otherwise.

### GetCacheKeyRegexOk

`func (o *CacheMetaModel) GetCacheKeyRegexOk() (*string, bool)`

GetCacheKeyRegexOk returns a tuple with the CacheKeyRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheKeyRegex

`func (o *CacheMetaModel) SetCacheKeyRegex(v string)`

SetCacheKeyRegex sets CacheKeyRegex field to given value.

### HasCacheKeyRegex

`func (o *CacheMetaModel) HasCacheKeyRegex() bool`

HasCacheKeyRegex returns a boolean if a field has been set.

### GetMethod

`func (o *CacheMetaModel) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CacheMetaModel) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CacheMetaModel) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *CacheMetaModel) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPath

`func (o *CacheMetaModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CacheMetaModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CacheMetaModel) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CacheMetaModel) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


