# CircuitBreakerMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**ReturnToServiceAfter** | Pointer to **int64** |  | [optional] 
**Samples** | Pointer to **int64** |  | [optional] 
**ThresholdPercent** | Pointer to **float64** |  | [optional] 

## Methods

### NewCircuitBreakerMeta

`func NewCircuitBreakerMeta() *CircuitBreakerMeta`

NewCircuitBreakerMeta instantiates a new CircuitBreakerMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitBreakerMetaWithDefaults

`func NewCircuitBreakerMetaWithDefaults() *CircuitBreakerMeta`

NewCircuitBreakerMetaWithDefaults instantiates a new CircuitBreakerMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *CircuitBreakerMeta) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CircuitBreakerMeta) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CircuitBreakerMeta) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *CircuitBreakerMeta) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPath

`func (o *CircuitBreakerMeta) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CircuitBreakerMeta) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CircuitBreakerMeta) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CircuitBreakerMeta) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetReturnToServiceAfter

`func (o *CircuitBreakerMeta) GetReturnToServiceAfter() int64`

GetReturnToServiceAfter returns the ReturnToServiceAfter field if non-nil, zero value otherwise.

### GetReturnToServiceAfterOk

`func (o *CircuitBreakerMeta) GetReturnToServiceAfterOk() (*int64, bool)`

GetReturnToServiceAfterOk returns a tuple with the ReturnToServiceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnToServiceAfter

`func (o *CircuitBreakerMeta) SetReturnToServiceAfter(v int64)`

SetReturnToServiceAfter sets ReturnToServiceAfter field to given value.

### HasReturnToServiceAfter

`func (o *CircuitBreakerMeta) HasReturnToServiceAfter() bool`

HasReturnToServiceAfter returns a boolean if a field has been set.

### GetSamples

`func (o *CircuitBreakerMeta) GetSamples() int64`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *CircuitBreakerMeta) GetSamplesOk() (*int64, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *CircuitBreakerMeta) SetSamples(v int64)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *CircuitBreakerMeta) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetThresholdPercent

`func (o *CircuitBreakerMeta) GetThresholdPercent() float64`

GetThresholdPercent returns the ThresholdPercent field if non-nil, zero value otherwise.

### GetThresholdPercentOk

`func (o *CircuitBreakerMeta) GetThresholdPercentOk() (*float64, bool)`

GetThresholdPercentOk returns a tuple with the ThresholdPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdPercent

`func (o *CircuitBreakerMeta) SetThresholdPercent(v float64)`

SetThresholdPercent sets ThresholdPercent field to given value.

### HasThresholdPercent

`func (o *CircuitBreakerMeta) HasThresholdPercent() bool`

HasThresholdPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


