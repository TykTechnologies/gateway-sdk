# RateLimitType2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Per** | Pointer to **float32** |  | [optional] 
**Rate** | Pointer to **float32** |  | [optional] 
**Smoothing** | Pointer to [**RateLimitSmoothing**](RateLimitSmoothing.md) |  | [optional] 

## Methods

### NewRateLimitType2

`func NewRateLimitType2() *RateLimitType2`

NewRateLimitType2 instantiates a new RateLimitType2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitType2WithDefaults

`func NewRateLimitType2WithDefaults() *RateLimitType2`

NewRateLimitType2WithDefaults instantiates a new RateLimitType2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPer

`func (o *RateLimitType2) GetPer() float32`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *RateLimitType2) GetPerOk() (*float32, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *RateLimitType2) SetPer(v float32)`

SetPer sets Per field to given value.

### HasPer

`func (o *RateLimitType2) HasPer() bool`

HasPer returns a boolean if a field has been set.

### GetRate

`func (o *RateLimitType2) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *RateLimitType2) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *RateLimitType2) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *RateLimitType2) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetSmoothing

`func (o *RateLimitType2) GetSmoothing() RateLimitSmoothing`

GetSmoothing returns the Smoothing field if non-nil, zero value otherwise.

### GetSmoothingOk

`func (o *RateLimitType2) GetSmoothingOk() (*RateLimitSmoothing, bool)`

GetSmoothingOk returns a tuple with the Smoothing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmoothing

`func (o *RateLimitType2) SetSmoothing(v RateLimitSmoothing)`

SetSmoothing sets Smoothing field to given value.

### HasSmoothing

`func (o *RateLimitType2) HasSmoothing() bool`

HasSmoothing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


