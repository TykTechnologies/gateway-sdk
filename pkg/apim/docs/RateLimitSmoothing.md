# RateLimitSmoothing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delay** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Step** | Pointer to **int32** |  | [optional] 
**Threshold** | Pointer to **int32** |  | [optional] 
**Trigger** | Pointer to **float32** |  | [optional] 

## Methods

### NewRateLimitSmoothing

`func NewRateLimitSmoothing() *RateLimitSmoothing`

NewRateLimitSmoothing instantiates a new RateLimitSmoothing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitSmoothingWithDefaults

`func NewRateLimitSmoothingWithDefaults() *RateLimitSmoothing`

NewRateLimitSmoothingWithDefaults instantiates a new RateLimitSmoothing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelay

`func (o *RateLimitSmoothing) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *RateLimitSmoothing) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *RateLimitSmoothing) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *RateLimitSmoothing) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetEnabled

`func (o *RateLimitSmoothing) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RateLimitSmoothing) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RateLimitSmoothing) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RateLimitSmoothing) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStep

`func (o *RateLimitSmoothing) GetStep() int32`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *RateLimitSmoothing) GetStepOk() (*int32, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *RateLimitSmoothing) SetStep(v int32)`

SetStep sets Step field to given value.

### HasStep

`func (o *RateLimitSmoothing) HasStep() bool`

HasStep returns a boolean if a field has been set.

### GetThreshold

`func (o *RateLimitSmoothing) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *RateLimitSmoothing) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *RateLimitSmoothing) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *RateLimitSmoothing) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetTrigger

`func (o *RateLimitSmoothing) GetTrigger() float32`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *RateLimitSmoothing) GetTriggerOk() (*float32, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *RateLimitSmoothing) SetTrigger(v float32)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *RateLimitSmoothing) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


