# CacheOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheAllSafeRequests** | Pointer to **bool** |  | [optional] 
**CacheByHeaders** | Pointer to **[]string** |  | [optional] 
**CacheControlTtlHeader** | Pointer to **string** |  | [optional] 
**CacheResponseCodes** | Pointer to **[]int32** |  | [optional] 
**CacheTimeout** | Pointer to **int64** |  | [optional] 
**EnableCache** | Pointer to **bool** |  | [optional] 
**EnableUpstreamCacheControl** | Pointer to **bool** |  | [optional] 

## Methods

### NewCacheOptions

`func NewCacheOptions() *CacheOptions`

NewCacheOptions instantiates a new CacheOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheOptionsWithDefaults

`func NewCacheOptionsWithDefaults() *CacheOptions`

NewCacheOptionsWithDefaults instantiates a new CacheOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheAllSafeRequests

`func (o *CacheOptions) GetCacheAllSafeRequests() bool`

GetCacheAllSafeRequests returns the CacheAllSafeRequests field if non-nil, zero value otherwise.

### GetCacheAllSafeRequestsOk

`func (o *CacheOptions) GetCacheAllSafeRequestsOk() (*bool, bool)`

GetCacheAllSafeRequestsOk returns a tuple with the CacheAllSafeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheAllSafeRequests

`func (o *CacheOptions) SetCacheAllSafeRequests(v bool)`

SetCacheAllSafeRequests sets CacheAllSafeRequests field to given value.

### HasCacheAllSafeRequests

`func (o *CacheOptions) HasCacheAllSafeRequests() bool`

HasCacheAllSafeRequests returns a boolean if a field has been set.

### GetCacheByHeaders

`func (o *CacheOptions) GetCacheByHeaders() []string`

GetCacheByHeaders returns the CacheByHeaders field if non-nil, zero value otherwise.

### GetCacheByHeadersOk

`func (o *CacheOptions) GetCacheByHeadersOk() (*[]string, bool)`

GetCacheByHeadersOk returns a tuple with the CacheByHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheByHeaders

`func (o *CacheOptions) SetCacheByHeaders(v []string)`

SetCacheByHeaders sets CacheByHeaders field to given value.

### HasCacheByHeaders

`func (o *CacheOptions) HasCacheByHeaders() bool`

HasCacheByHeaders returns a boolean if a field has been set.

### SetCacheByHeadersNil

`func (o *CacheOptions) SetCacheByHeadersNil(b bool)`

 SetCacheByHeadersNil sets the value for CacheByHeaders to be an explicit nil

### UnsetCacheByHeaders
`func (o *CacheOptions) UnsetCacheByHeaders()`

UnsetCacheByHeaders ensures that no value is present for CacheByHeaders, not even an explicit nil
### GetCacheControlTtlHeader

`func (o *CacheOptions) GetCacheControlTtlHeader() string`

GetCacheControlTtlHeader returns the CacheControlTtlHeader field if non-nil, zero value otherwise.

### GetCacheControlTtlHeaderOk

`func (o *CacheOptions) GetCacheControlTtlHeaderOk() (*string, bool)`

GetCacheControlTtlHeaderOk returns a tuple with the CacheControlTtlHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheControlTtlHeader

`func (o *CacheOptions) SetCacheControlTtlHeader(v string)`

SetCacheControlTtlHeader sets CacheControlTtlHeader field to given value.

### HasCacheControlTtlHeader

`func (o *CacheOptions) HasCacheControlTtlHeader() bool`

HasCacheControlTtlHeader returns a boolean if a field has been set.

### GetCacheResponseCodes

`func (o *CacheOptions) GetCacheResponseCodes() []int32`

GetCacheResponseCodes returns the CacheResponseCodes field if non-nil, zero value otherwise.

### GetCacheResponseCodesOk

`func (o *CacheOptions) GetCacheResponseCodesOk() (*[]int32, bool)`

GetCacheResponseCodesOk returns a tuple with the CacheResponseCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheResponseCodes

`func (o *CacheOptions) SetCacheResponseCodes(v []int32)`

SetCacheResponseCodes sets CacheResponseCodes field to given value.

### HasCacheResponseCodes

`func (o *CacheOptions) HasCacheResponseCodes() bool`

HasCacheResponseCodes returns a boolean if a field has been set.

### SetCacheResponseCodesNil

`func (o *CacheOptions) SetCacheResponseCodesNil(b bool)`

 SetCacheResponseCodesNil sets the value for CacheResponseCodes to be an explicit nil

### UnsetCacheResponseCodes
`func (o *CacheOptions) UnsetCacheResponseCodes()`

UnsetCacheResponseCodes ensures that no value is present for CacheResponseCodes, not even an explicit nil
### GetCacheTimeout

`func (o *CacheOptions) GetCacheTimeout() int64`

GetCacheTimeout returns the CacheTimeout field if non-nil, zero value otherwise.

### GetCacheTimeoutOk

`func (o *CacheOptions) GetCacheTimeoutOk() (*int64, bool)`

GetCacheTimeoutOk returns a tuple with the CacheTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTimeout

`func (o *CacheOptions) SetCacheTimeout(v int64)`

SetCacheTimeout sets CacheTimeout field to given value.

### HasCacheTimeout

`func (o *CacheOptions) HasCacheTimeout() bool`

HasCacheTimeout returns a boolean if a field has been set.

### GetEnableCache

`func (o *CacheOptions) GetEnableCache() bool`

GetEnableCache returns the EnableCache field if non-nil, zero value otherwise.

### GetEnableCacheOk

`func (o *CacheOptions) GetEnableCacheOk() (*bool, bool)`

GetEnableCacheOk returns a tuple with the EnableCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCache

`func (o *CacheOptions) SetEnableCache(v bool)`

SetEnableCache sets EnableCache field to given value.

### HasEnableCache

`func (o *CacheOptions) HasEnableCache() bool`

HasEnableCache returns a boolean if a field has been set.

### GetEnableUpstreamCacheControl

`func (o *CacheOptions) GetEnableUpstreamCacheControl() bool`

GetEnableUpstreamCacheControl returns the EnableUpstreamCacheControl field if non-nil, zero value otherwise.

### GetEnableUpstreamCacheControlOk

`func (o *CacheOptions) GetEnableUpstreamCacheControlOk() (*bool, bool)`

GetEnableUpstreamCacheControlOk returns a tuple with the EnableUpstreamCacheControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUpstreamCacheControl

`func (o *CacheOptions) SetEnableUpstreamCacheControl(v bool)`

SetEnableUpstreamCacheControl sets EnableUpstreamCacheControl field to given value.

### HasEnableUpstreamCacheControl

`func (o *CacheOptions) HasEnableUpstreamCacheControl() bool`

HasEnableUpstreamCacheControl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


