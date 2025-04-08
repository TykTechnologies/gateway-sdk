# RateLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Per** | Pointer to **int32** |  | [optional] 
**Rate** | Pointer to **int32** |  | [optional] 

## Methods

### NewRateLimit

`func NewRateLimit() *RateLimit`

NewRateLimit instantiates a new RateLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitWithDefaults

`func NewRateLimitWithDefaults() *RateLimit`

NewRateLimitWithDefaults instantiates a new RateLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *RateLimit) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RateLimit) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RateLimit) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RateLimit) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPer

`func (o *RateLimit) GetPer() int32`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *RateLimit) GetPerOk() (*int32, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *RateLimit) SetPer(v int32)`

SetPer sets Per field to given value.

### HasPer

`func (o *RateLimit) HasPer() bool`

HasPer returns a boolean if a field has been set.

### GetRate

`func (o *RateLimit) GetRate() int32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *RateLimit) GetRateOk() (*int32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *RateLimit) SetRate(v int32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *RateLimit) HasRate() bool`

HasRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


