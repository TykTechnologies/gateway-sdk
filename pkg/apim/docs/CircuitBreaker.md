# CircuitBreaker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoolDownPeriod** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**HalfOpenStateEnabled** | Pointer to **bool** |  | [optional] 
**SampleSize** | Pointer to **int32** |  | [optional] 
**Threshold** | Pointer to **float32** |  | [optional] 

## Methods

### NewCircuitBreaker

`func NewCircuitBreaker() *CircuitBreaker`

NewCircuitBreaker instantiates a new CircuitBreaker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitBreakerWithDefaults

`func NewCircuitBreakerWithDefaults() *CircuitBreaker`

NewCircuitBreakerWithDefaults instantiates a new CircuitBreaker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoolDownPeriod

`func (o *CircuitBreaker) GetCoolDownPeriod() int32`

GetCoolDownPeriod returns the CoolDownPeriod field if non-nil, zero value otherwise.

### GetCoolDownPeriodOk

`func (o *CircuitBreaker) GetCoolDownPeriodOk() (*int32, bool)`

GetCoolDownPeriodOk returns a tuple with the CoolDownPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolDownPeriod

`func (o *CircuitBreaker) SetCoolDownPeriod(v int32)`

SetCoolDownPeriod sets CoolDownPeriod field to given value.

### HasCoolDownPeriod

`func (o *CircuitBreaker) HasCoolDownPeriod() bool`

HasCoolDownPeriod returns a boolean if a field has been set.

### GetEnabled

`func (o *CircuitBreaker) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CircuitBreaker) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CircuitBreaker) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CircuitBreaker) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHalfOpenStateEnabled

`func (o *CircuitBreaker) GetHalfOpenStateEnabled() bool`

GetHalfOpenStateEnabled returns the HalfOpenStateEnabled field if non-nil, zero value otherwise.

### GetHalfOpenStateEnabledOk

`func (o *CircuitBreaker) GetHalfOpenStateEnabledOk() (*bool, bool)`

GetHalfOpenStateEnabledOk returns a tuple with the HalfOpenStateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHalfOpenStateEnabled

`func (o *CircuitBreaker) SetHalfOpenStateEnabled(v bool)`

SetHalfOpenStateEnabled sets HalfOpenStateEnabled field to given value.

### HasHalfOpenStateEnabled

`func (o *CircuitBreaker) HasHalfOpenStateEnabled() bool`

HasHalfOpenStateEnabled returns a boolean if a field has been set.

### GetSampleSize

`func (o *CircuitBreaker) GetSampleSize() int32`

GetSampleSize returns the SampleSize field if non-nil, zero value otherwise.

### GetSampleSizeOk

`func (o *CircuitBreaker) GetSampleSizeOk() (*int32, bool)`

GetSampleSizeOk returns a tuple with the SampleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleSize

`func (o *CircuitBreaker) SetSampleSize(v int32)`

SetSampleSize sets SampleSize field to given value.

### HasSampleSize

`func (o *CircuitBreaker) HasSampleSize() bool`

HasSampleSize returns a boolean if a field has been set.

### GetThreshold

`func (o *CircuitBreaker) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *CircuitBreaker) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *CircuitBreaker) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *CircuitBreaker) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


