# CacheOptionsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheAllSafeRequests** | Pointer to **bool** |  | [optional] 
**CacheControlTtlHeader** | Pointer to **string** |  | [optional] 
**CacheResponseCodes** | Pointer to **[]int64** |  | [optional] 
**CacheTimeout** | Pointer to **int64** |  | [optional] 
**EnableCache** | Pointer to **bool** |  | [optional] 
**EnableUpstreamCacheControl** | Pointer to **bool** |  | [optional] 

## Methods

### NewCacheOptionsModel

`func NewCacheOptionsModel() *CacheOptionsModel`

NewCacheOptionsModel instantiates a new CacheOptionsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheOptionsModelWithDefaults

`func NewCacheOptionsModelWithDefaults() *CacheOptionsModel`

NewCacheOptionsModelWithDefaults instantiates a new CacheOptionsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheAllSafeRequests

`func (o *CacheOptionsModel) GetCacheAllSafeRequests() bool`

GetCacheAllSafeRequests returns the CacheAllSafeRequests field if non-nil, zero value otherwise.

### GetCacheAllSafeRequestsOk

`func (o *CacheOptionsModel) GetCacheAllSafeRequestsOk() (*bool, bool)`

GetCacheAllSafeRequestsOk returns a tuple with the CacheAllSafeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheAllSafeRequests

`func (o *CacheOptionsModel) SetCacheAllSafeRequests(v bool)`

SetCacheAllSafeRequests sets CacheAllSafeRequests field to given value.

### HasCacheAllSafeRequests

`func (o *CacheOptionsModel) HasCacheAllSafeRequests() bool`

HasCacheAllSafeRequests returns a boolean if a field has been set.

### GetCacheControlTtlHeader

`func (o *CacheOptionsModel) GetCacheControlTtlHeader() string`

GetCacheControlTtlHeader returns the CacheControlTtlHeader field if non-nil, zero value otherwise.

### GetCacheControlTtlHeaderOk

`func (o *CacheOptionsModel) GetCacheControlTtlHeaderOk() (*string, bool)`

GetCacheControlTtlHeaderOk returns a tuple with the CacheControlTtlHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheControlTtlHeader

`func (o *CacheOptionsModel) SetCacheControlTtlHeader(v string)`

SetCacheControlTtlHeader sets CacheControlTtlHeader field to given value.

### HasCacheControlTtlHeader

`func (o *CacheOptionsModel) HasCacheControlTtlHeader() bool`

HasCacheControlTtlHeader returns a boolean if a field has been set.

### GetCacheResponseCodes

`func (o *CacheOptionsModel) GetCacheResponseCodes() []int64`

GetCacheResponseCodes returns the CacheResponseCodes field if non-nil, zero value otherwise.

### GetCacheResponseCodesOk

`func (o *CacheOptionsModel) GetCacheResponseCodesOk() (*[]int64, bool)`

GetCacheResponseCodesOk returns a tuple with the CacheResponseCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheResponseCodes

`func (o *CacheOptionsModel) SetCacheResponseCodes(v []int64)`

SetCacheResponseCodes sets CacheResponseCodes field to given value.

### HasCacheResponseCodes

`func (o *CacheOptionsModel) HasCacheResponseCodes() bool`

HasCacheResponseCodes returns a boolean if a field has been set.

### GetCacheTimeout

`func (o *CacheOptionsModel) GetCacheTimeout() int64`

GetCacheTimeout returns the CacheTimeout field if non-nil, zero value otherwise.

### GetCacheTimeoutOk

`func (o *CacheOptionsModel) GetCacheTimeoutOk() (*int64, bool)`

GetCacheTimeoutOk returns a tuple with the CacheTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTimeout

`func (o *CacheOptionsModel) SetCacheTimeout(v int64)`

SetCacheTimeout sets CacheTimeout field to given value.

### HasCacheTimeout

`func (o *CacheOptionsModel) HasCacheTimeout() bool`

HasCacheTimeout returns a boolean if a field has been set.

### GetEnableCache

`func (o *CacheOptionsModel) GetEnableCache() bool`

GetEnableCache returns the EnableCache field if non-nil, zero value otherwise.

### GetEnableCacheOk

`func (o *CacheOptionsModel) GetEnableCacheOk() (*bool, bool)`

GetEnableCacheOk returns a tuple with the EnableCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCache

`func (o *CacheOptionsModel) SetEnableCache(v bool)`

SetEnableCache sets EnableCache field to given value.

### HasEnableCache

`func (o *CacheOptionsModel) HasEnableCache() bool`

HasEnableCache returns a boolean if a field has been set.

### GetEnableUpstreamCacheControl

`func (o *CacheOptionsModel) GetEnableUpstreamCacheControl() bool`

GetEnableUpstreamCacheControl returns the EnableUpstreamCacheControl field if non-nil, zero value otherwise.

### GetEnableUpstreamCacheControlOk

`func (o *CacheOptionsModel) GetEnableUpstreamCacheControlOk() (*bool, bool)`

GetEnableUpstreamCacheControlOk returns a tuple with the EnableUpstreamCacheControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUpstreamCacheControl

`func (o *CacheOptionsModel) SetEnableUpstreamCacheControl(v bool)`

SetEnableUpstreamCacheControl sets EnableUpstreamCacheControl field to given value.

### HasEnableUpstreamCacheControl

`func (o *CacheOptionsModel) HasEnableUpstreamCacheControl() bool`

HasEnableUpstreamCacheControl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


