# RateLimitMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Per** | Pointer to **float32** |  | [optional] 
**Rate** | Pointer to **float32** |  | [optional] 

## Methods

### NewRateLimitMeta

`func NewRateLimitMeta() *RateLimitMeta`

NewRateLimitMeta instantiates a new RateLimitMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitMetaWithDefaults

`func NewRateLimitMetaWithDefaults() *RateLimitMeta`

NewRateLimitMetaWithDefaults instantiates a new RateLimitMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *RateLimitMeta) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *RateLimitMeta) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *RateLimitMeta) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *RateLimitMeta) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetMethod

`func (o *RateLimitMeta) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RateLimitMeta) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RateLimitMeta) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RateLimitMeta) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPath

`func (o *RateLimitMeta) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RateLimitMeta) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RateLimitMeta) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RateLimitMeta) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPer

`func (o *RateLimitMeta) GetPer() float32`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *RateLimitMeta) GetPerOk() (*float32, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *RateLimitMeta) SetPer(v float32)`

SetPer sets Per field to given value.

### HasPer

`func (o *RateLimitMeta) HasPer() bool`

HasPer returns a boolean if a field has been set.

### GetRate

`func (o *RateLimitMeta) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *RateLimitMeta) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *RateLimitMeta) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *RateLimitMeta) HasRate() bool`

HasRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


