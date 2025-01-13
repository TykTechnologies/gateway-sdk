# APILimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxQueryDepth** | Pointer to **int32** |  | [optional] 
**Per** | Pointer to **float32** |  | [optional] 
**QuotaMax** | Pointer to **int32** |  | [optional] 
**QuotaRemaining** | Pointer to **int32** |  | [optional] 
**QuotaRenewalRate** | Pointer to **int32** |  | [optional] 
**QuotaRenews** | Pointer to **int32** |  | [optional] 
**Rate** | Pointer to **float32** |  | [optional] 
**Smoothing** | Pointer to [**RateLimitSmoothing**](RateLimitSmoothing.md) |  | [optional] 
**ThrottleInterval** | Pointer to **float32** |  | [optional] 
**ThrottleRetryLimit** | Pointer to **int32** |  | [optional] 

## Methods

### NewAPILimit

`func NewAPILimit() *APILimit`

NewAPILimit instantiates a new APILimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPILimitWithDefaults

`func NewAPILimitWithDefaults() *APILimit`

NewAPILimitWithDefaults instantiates a new APILimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxQueryDepth

`func (o *APILimit) GetMaxQueryDepth() int32`

GetMaxQueryDepth returns the MaxQueryDepth field if non-nil, zero value otherwise.

### GetMaxQueryDepthOk

`func (o *APILimit) GetMaxQueryDepthOk() (*int32, bool)`

GetMaxQueryDepthOk returns a tuple with the MaxQueryDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueryDepth

`func (o *APILimit) SetMaxQueryDepth(v int32)`

SetMaxQueryDepth sets MaxQueryDepth field to given value.

### HasMaxQueryDepth

`func (o *APILimit) HasMaxQueryDepth() bool`

HasMaxQueryDepth returns a boolean if a field has been set.

### GetPer

`func (o *APILimit) GetPer() float32`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *APILimit) GetPerOk() (*float32, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *APILimit) SetPer(v float32)`

SetPer sets Per field to given value.

### HasPer

`func (o *APILimit) HasPer() bool`

HasPer returns a boolean if a field has been set.

### GetQuotaMax

`func (o *APILimit) GetQuotaMax() int32`

GetQuotaMax returns the QuotaMax field if non-nil, zero value otherwise.

### GetQuotaMaxOk

`func (o *APILimit) GetQuotaMaxOk() (*int32, bool)`

GetQuotaMaxOk returns a tuple with the QuotaMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaMax

`func (o *APILimit) SetQuotaMax(v int32)`

SetQuotaMax sets QuotaMax field to given value.

### HasQuotaMax

`func (o *APILimit) HasQuotaMax() bool`

HasQuotaMax returns a boolean if a field has been set.

### GetQuotaRemaining

`func (o *APILimit) GetQuotaRemaining() int32`

GetQuotaRemaining returns the QuotaRemaining field if non-nil, zero value otherwise.

### GetQuotaRemainingOk

`func (o *APILimit) GetQuotaRemainingOk() (*int32, bool)`

GetQuotaRemainingOk returns a tuple with the QuotaRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaRemaining

`func (o *APILimit) SetQuotaRemaining(v int32)`

SetQuotaRemaining sets QuotaRemaining field to given value.

### HasQuotaRemaining

`func (o *APILimit) HasQuotaRemaining() bool`

HasQuotaRemaining returns a boolean if a field has been set.

### GetQuotaRenewalRate

`func (o *APILimit) GetQuotaRenewalRate() int32`

GetQuotaRenewalRate returns the QuotaRenewalRate field if non-nil, zero value otherwise.

### GetQuotaRenewalRateOk

`func (o *APILimit) GetQuotaRenewalRateOk() (*int32, bool)`

GetQuotaRenewalRateOk returns a tuple with the QuotaRenewalRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaRenewalRate

`func (o *APILimit) SetQuotaRenewalRate(v int32)`

SetQuotaRenewalRate sets QuotaRenewalRate field to given value.

### HasQuotaRenewalRate

`func (o *APILimit) HasQuotaRenewalRate() bool`

HasQuotaRenewalRate returns a boolean if a field has been set.

### GetQuotaRenews

`func (o *APILimit) GetQuotaRenews() int32`

GetQuotaRenews returns the QuotaRenews field if non-nil, zero value otherwise.

### GetQuotaRenewsOk

`func (o *APILimit) GetQuotaRenewsOk() (*int32, bool)`

GetQuotaRenewsOk returns a tuple with the QuotaRenews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaRenews

`func (o *APILimit) SetQuotaRenews(v int32)`

SetQuotaRenews sets QuotaRenews field to given value.

### HasQuotaRenews

`func (o *APILimit) HasQuotaRenews() bool`

HasQuotaRenews returns a boolean if a field has been set.

### GetRate

`func (o *APILimit) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *APILimit) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *APILimit) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *APILimit) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetSmoothing

`func (o *APILimit) GetSmoothing() RateLimitSmoothing`

GetSmoothing returns the Smoothing field if non-nil, zero value otherwise.

### GetSmoothingOk

`func (o *APILimit) GetSmoothingOk() (*RateLimitSmoothing, bool)`

GetSmoothingOk returns a tuple with the Smoothing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmoothing

`func (o *APILimit) SetSmoothing(v RateLimitSmoothing)`

SetSmoothing sets Smoothing field to given value.

### HasSmoothing

`func (o *APILimit) HasSmoothing() bool`

HasSmoothing returns a boolean if a field has been set.

### GetThrottleInterval

`func (o *APILimit) GetThrottleInterval() float32`

GetThrottleInterval returns the ThrottleInterval field if non-nil, zero value otherwise.

### GetThrottleIntervalOk

`func (o *APILimit) GetThrottleIntervalOk() (*float32, bool)`

GetThrottleIntervalOk returns a tuple with the ThrottleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleInterval

`func (o *APILimit) SetThrottleInterval(v float32)`

SetThrottleInterval sets ThrottleInterval field to given value.

### HasThrottleInterval

`func (o *APILimit) HasThrottleInterval() bool`

HasThrottleInterval returns a boolean if a field has been set.

### GetThrottleRetryLimit

`func (o *APILimit) GetThrottleRetryLimit() int32`

GetThrottleRetryLimit returns the ThrottleRetryLimit field if non-nil, zero value otherwise.

### GetThrottleRetryLimitOk

`func (o *APILimit) GetThrottleRetryLimitOk() (*int32, bool)`

GetThrottleRetryLimitOk returns a tuple with the ThrottleRetryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleRetryLimit

`func (o *APILimit) SetThrottleRetryLimit(v int32)`

SetThrottleRetryLimit sets ThrottleRetryLimit field to given value.

### HasThrottleRetryLimit

`func (o *APILimit) HasThrottleRetryLimit() bool`

HasThrottleRetryLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


